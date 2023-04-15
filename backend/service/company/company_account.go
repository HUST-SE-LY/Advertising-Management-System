package company

import (
	"backend/global"
	"backend/models/company_model"
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
	// Check in CompanyToBeReviewed
	if !errors.Is(global.GVA_DB.Where("account = ?", req.Account).First(&company_model.CompanyToBeReviewed{}).Error,
		gorm.ErrRecordNotFound) {
		return status.SameAccountExists
	}
	// Check in Company
	if !errors.Is(global.GVA_DB.Where("account = ?", req.Account).First(&company_model.Company{}).Error,
		gorm.ErrRecordNotFound) {
		return status.SameAccountExists
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	companyToBeReviewed := company_model.CompanyToBeReviewed{
		Account:               req.Account,
		Password:              string(encryptedPassword),
		Name:                  req.Name,
		Address:               req.Address,
		ManagerName:           req.ManagerName,
		ManagerTel:            req.ManagerTel,
		BusinessLicenseNumber: req.BusinessLicenseNumber,
	}
	err = global.GVA_DB.Create(&companyToBeReviewed).Error
	return err
}

func (c *CompanyAccountService) CompanyLogin(req request.CompanyLoginReq) (company company_model.Company, companyToken company_model.CompanyToken, err error) {
	err = global.GVA_DB.Where("account = ?", req.Account).Take(&company).Error
	if company == (company_model.Company{}) {
		// Then search in CompanyToBeReviewed
		var companyToBeReviewed *company_model.CompanyToBeReviewed
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
		global.GVA_DB.Take(companyToken, company.Id)

		if companyToken == (company_model.CompanyToken{}) {
			companyToken = company_model.CompanyToken{
				CompanyId: company.Id,
				Token:     token,
			}
			if err = global.GVA_DB.Create(&companyToken).Error; err != nil {
				return
			}
		} else {
			companyToken.Token = token
			if err = global.GVA_DB.Save(&companyToken).Error; err != nil {
				return
			}
		}
	}
	return company, companyToken, err
}

func (c *CompanyAccountService) FindCompanyToken(token string) (companyToken company_model.CompanyToken, err error) {
	err = global.GVA_DB.Take(&companyToken, "token = ?", token).Error
	return
}

func (c *CompanyAccountService) DeleteCompanyToken(token string) error {
	err := global.GVA_DB.Delete(&[]company_model.CompanyToken{}, "token = ?", token).Error
	return err
}

func getNewToken(account, password string) (token string) {
	token, _ = jwt.GenerateToken(account, password, company_model.COMPANY_ROLE)
	return
}
