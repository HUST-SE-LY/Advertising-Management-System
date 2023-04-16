package manage

import (
	"backend/global"
	"backend/models/company_model/entity"
	"backend/models/manage_model/enum"
	"backend/utils/functional"
	"errors"
	"fmt"
	"gorm.io/gorm"
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

func (m *ManageCompanyService) GetCompaniesByTerm(term string, termType enum.CompanySearchType) (companies []entity.Company, err error) {
	whereCondition := fmt.Sprintf("%s LIKE ?", termType.ToString())
	term = fmt.Sprintf("%%%s%%", term)
	if err = global.GVA_DB.Where(whereCondition, term).Find(&companies).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Company{}, nil
		}
		return nil, err
	}
	return companies, nil
}

func (m *ManageCompanyService) AllowRegistrationForCompanies(companyAccounts []string) (registeredCompanies *[]entity.Company, err error) {
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
