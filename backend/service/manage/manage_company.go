package manage

import (
	"backend/global"
	"backend/models/company_model"
	"backend/utils/functional"
)

type ManageCompanyService struct {
}

func (m *ManageCompanyService) ReadAllCompaniesToBeReviewed() (companies *[]company_model.CompanyToBeReviewed, err error) {
	if err = global.GVA_DB.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (m *ManageCompanyService) AllowRegistrationForCompanies(companyAccounts *[]string) (registeredCompanies *[]company_model.Company, err error) {
	var companiesToBeReviewed *[]company_model.CompanyToBeReviewed
	if err = global.GVA_DB.Where("account IN ?", companyAccounts).Find(&companiesToBeReviewed).Error; err != nil {
		return nil, err
	}
	companies := functional.Map(*companiesToBeReviewed, company_model.CompanyToBeReviewed.ToCompany)
	companiesNumber := len(companies)

	if err = global.GVA_DB.CreateInBatches(&companies, companiesNumber).Error; err != nil {
		return nil, err
	}
	return &companies, nil
}
