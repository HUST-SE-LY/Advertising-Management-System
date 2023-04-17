package company

import (
	"backend/global"
	"backend/models/company_model"
	"backend/models/company_model/entity"
	"backend/models/company_model/request"
	"backend/utils/jwt"
	"backend/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CompanyAccountService struct {
}

func (c *CompanyAccountService) RegisterCompany(req *request.CompanyRegisterReq) error {
	// Check in CompanyPendingReview
	if !errors.Is(global.GVA_DB.Where("account = ?", req.Account).First(&entity.CompanyPendingReview{}).Error,
		gorm.ErrRecordNotFound) {
		return status.SameAccountExists
	}
	// Check in Company
	if !errors.Is(global.GVA_DB.Where("account = ?", req.Account).First(&entity.Company{}).Error,
		gorm.ErrRecordNotFound) {
		return status.SameAccountExists
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	companyToBeReviewed := entity.NewCompanyPendingReview(
		*entity.NewCompanyWithoutId(string(encryptedPassword), *company_model.NewCompanyInfo(
			req.Account,
			req.Name,
			req.Address,
			req.ManagerName,
			req.ManagerTel,
			req.BusinessLicenseNumber,
		)),
	)
	err = global.GVA_DB.Create(&companyToBeReviewed).Error
	return err
}

func (c *CompanyAccountService) CompanyLogin(req request.CompanyLoginReq) (company entity.Company, companyToken entity.CompanyToken, err error) {
	err = global.GVA_DB.Where("account = ?", req.Account).Take(&company).Error
	if company == (entity.Company{}) {
		// Then search in CompanyPendingReview
		var companyToBeReviewed *entity.CompanyPendingReview
		err = global.GVA_DB.Where("account = ?", req.Account).Take(&companyToBeReviewed).Error
		if companyToBeReviewed != nil {
			err = status.UserIsPendingReview
			return
		} else {
			err = status.AccountNotFound
			return
		}
	} else {
		// Check password
		if err = bcrypt.CompareHashAndPassword([]byte(company.Password), []byte(req.Password)); err != nil {
			err = status.WrongPassword
			return
		}
		token := getNewToken(company.Account, company.Password)

		companyToken = *entity.NewCompanyToken(
			company.Id,
			token,
		)
		if err = global.GVA_DB.Save(&companyToken).Error; err != nil {
			return
		}
	}
	return company, companyToken, err
}

func (c *CompanyAccountService) CompanyLogout(token string) (err error) {
	return c.DeleteCompanyToken(token)
}

func (c *CompanyAccountService) CompanyUpdateInfo(req request.CompanyUpdateInfoReq) (err error) {
	var company entity.Company
	err = global.GVA_DB.Where("account = ?", req.Account).Take(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = status.AccountNotFound
		}
		return
	}
	companyInfoPendingReview := entity.NewCompanyInfoPendingReview(
		company.Id, *company_model.NewCompanyInfo(
			req.Account,
			req.Name,
			req.Address,
			req.ManagerName,
			req.ManagerTel,
			req.BusinessLicenseNumber,
		),
	)
	err = global.GVA_DB.Save(&companyInfoPendingReview).Error
	return
}

func (c *CompanyAccountService) CompanyUpdatePwd(req request.CompanyUpdatePwdReq, account string) (err error) {
	var company entity.Company
	err = global.GVA_DB.Where("account = ?", account).Take(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = status.AccountNotFound
		}
		return
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&company).Update("password", newPassword).Error
	return
}

func (c *CompanyAccountService) CompanyCancel(token string) (err error) {
	// First, cancel the account, and then delete the token.
	var company entity.Company

	claims, _ := jwt.ParseToken(token)
	account := claims.Username
	err = global.GVA_DB.Where("account = ?", account).Take(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = status.AccountNotFound
		}
		return
	}

	err = global.GVA_DB.Delete(&entity.Company{}, company.Id).Error
	if err != nil {
		err = status.CancelError
		return
	}
	err = c.DeleteCompanyToken(token)
	if err != nil {
		err = status.LogoutError
		return
	}
	return nil
}

func (c *CompanyAccountService) FindCompanyToken(token string) (companyToken entity.CompanyToken, err error) {
	err = global.GVA_DB.Take(&companyToken, "token = ?", token).Error
	return
}

func (c *CompanyAccountService) DeleteCompanyToken(token string) error {
	err := global.GVA_DB.Delete(&[]entity.CompanyToken{}, "token = ?", token).Error
	return err
}

func getNewToken(account, password string) (token string) {
	token, _ = jwt.GenerateToken(account, password, entity.COMPANY_ROLE)
	return
}
