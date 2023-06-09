package company

import (
	"backend/global"
	"backend/models/company_model"
	"backend/models/company_model/entity"
	"backend/models/company_model/request"
	recordEntity "backend/models/record_model/entity"
	"backend/models/record_model/enum"
	"backend/utils/jwt"
	"backend/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CompanyAccountService struct {
}

func (c *CompanyAccountService) CompanyRegister(req *request.CompanyRegisterReq) error {
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

func (c *CompanyAccountService) CompanyUpdateInfo(req request.CompanyUpdateInfoReq, token string) (err error) {
	var company entity.Company
	claims, _ := jwt.ParseToken(token)
	accountToBeRevised := claims.Username
	err = global.GVA_DB.Where("account = ?", accountToBeRevised).Take(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = status.AccountNotFound
		}
		return
	}
	companyInfoPendingReview := entity.NewCompanyInfoPendingReview(
		company.Id, *company_model.NewCompanyInfo(
			accountToBeRevised,
			req.Name,
			req.Address,
			req.ManagerName,
			req.ManagerTel,
			req.BusinessLicenseNumber,
		),
	)
	err = global.GVA_DB.Save(&companyInfoPendingReview).Error

	date, _ := strconv.Atoi(time.Now().Format("20060102"))
	// add to application record
	applicationRecord := recordEntity.ApplicationRecord{
		Account: accountToBeRevised,
		Status:  enum.UnderReview,
		Type:    enum.Info,
		Date:    date,
	}
	err = global.GVA_DB.Save(&applicationRecord).Error
	return
}

func (c *CompanyAccountService) CompanyRecharge(req request.CompanyRechargeReq, account string) (err error) {
	var company entity.Company
	err = global.GVA_DB.Where("account = ?", account).Take(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = status.AccountNotFound
		}
		return
	}
	cash := company.Balance + req.Money
	err = global.GVA_DB.Model(&company).Update("balance", cash).Error
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

func (c *CompanyAccountService) CompanyGetInfo(token string) (company entity.Company, err error) {
	claims, err := jwt.ParseToken(token)
	account := claims.Username
	if err = global.GVA_DB.Where("account = ?", account).First(&company).Error; err != nil {
		return entity.Company{}, err
	}
	return company, nil
}

func (c *CompanyAccountService) CompanyGetAllApplication(token string) (applicationRecord []recordEntity.ApplicationRecord, err error) {
	claims, err := jwt.ParseToken(token)
	account := claims.Username
	if err = global.GVA_DB.Where("account = ?", account).Find(&applicationRecord).Error; err != nil {
		return
	}
	return applicationRecord, nil
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
