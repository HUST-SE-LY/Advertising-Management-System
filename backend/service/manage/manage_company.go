package manage

import (
	"backend/global"
	"backend/models/company_model/entity"
	"backend/utils/functional"
)

type ManageCompanyService struct {
}

func (m *ManageCompanyService) GetAllCompanies() (companies []entity.Company, err error) {
	if err = global.GVA_DB.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (m *ManageCompanyService) GetAllCompaniesToBeReviewed() (companies []entity.CompanyToBeReviewed, err error) {
	if err = global.GVA_DB.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (m *ManageCompanyService) AllowRegistrationForCompanies(companyAccounts *[]string) (registeredCompanies *[]entity.Company, err error) {
	var companiesToBeReviewed *[]entity.CompanyToBeReviewed
	if err = global.GVA_DB.Where("account IN ?", companyAccounts).Find(&companiesToBeReviewed).Error; err != nil {
		return nil, err
	}
	companies := functional.Map(*companiesToBeReviewed, entity.CompanyToBeReviewed.ToCompany)
	companiesNumber := len(companies)

	if err = global.GVA_DB.CreateInBatches(&companies, companiesNumber).Error; err != nil {
		return nil, err
	}
	return &companies, nil
}
