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

func (m *ManageCompanyService) GetCompaniesCount() (int64, error) {
	var count int64 = 0
	var companies []entity.Company
	if err := global.GVA_DB.Find(&companies).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (m *ManageCompanyService) GetInfoPendingReviewCompaniesCount() (int64, error) {
	var count int64 = 0
	var companies []entity.CompanyInfoPendingReview
	if err := global.GVA_DB.Find(&companies).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (m *ManageCompanyService) GetPendingReviewCompaniesCount() (int64, error) {
	var count int64 = 0
	var companies []entity.CompanyPendingReview
	if err := global.GVA_DB.Find(&companies).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (m *ManageCompanyService) GetAllCompanies() (companies []entity.Company, err error) {
	if err = global.GVA_DB.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (m *ManageCompanyService) GetAllCompaniesToBeReviewed() (companies []entity.CompanyPendingReview, err error) {
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

func (m *ManageCompanyService) AllowRegistrationForCompanies(companyAccounts []string) (registeredCompaniesAccount []string, err error) {
	var companiesToBeReviewed []entity.CompanyPendingReview
	if err = global.GVA_DB.Where("account IN ?", companyAccounts).Find(&companiesToBeReviewed).Error; err != nil {
		return nil, err
	}
	companies := functional.Map(companiesToBeReviewed, entity.CompanyPendingReview.ToCompany)
	companiesNumber := len(companies)

	if err = global.GVA_DB.CreateInBatches(&companies, companiesNumber).Error; err != nil {
		return nil, err
	}

	return functional.Map(companies, func(com entity.Company) string {
		return com.Account
	}), nil
}

func (m *ManageCompanyService) AllowCompaniesInfoUpdate(companyAccounts []string) (allowAccounts []string, err error) {
	var companiesToBeUpdated []entity.CompanyInfoPendingReview
	if err = global.GVA_DB.Where("account IN ?", companyAccounts).Find(&companiesToBeUpdated).Error; err != nil {
		return
	}
	companies := functional.Map(companiesToBeUpdated, func(com entity.CompanyInfoPendingReview) entity.Company {
		return entity.Company{
			Id:          com.Id,
			CompanyInfo: com.CompanyInfo,
		}
	})

	for _, company := range companies {
		global.GVA_DB.Model(&company).Updates(company)
	}

	return functional.Map(companies, func(com entity.Company) string {
		return com.Account
	}), nil
}
