package manage

import (
	"backend/global"
	"backend/models/company_model/entity"
	"backend/models/manage_model/enum"
	recordEntity "backend/models/record_model/entity"
	enum2 "backend/models/record_model/enum"
	"backend/utils/functional"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ManageCompanyService struct {
}

func (m *ManageCompanyService) GetInfoPendingReviewCompaniesCount() (int, error) {
	var companies []entity.CompanyInfoPendingReview
	if err := global.GVA_DB.Find(&companies).Error; err != nil {
		return 0, err
	}
	return len(companies), nil
}

func (m *ManageCompanyService) GetPendingReviewCompaniesCount() (int, error) {
	var companies []entity.CompanyPendingReview
	if err := global.GVA_DB.Find(&companies).Error; err != nil {
		return 0, err
	}
	return len(companies), nil
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

	if err = global.GVA_DB.Delete(&companiesToBeReviewed).Error; err != nil {
		return nil, err
	}

	return functional.Map(companies, func(com entity.Company) string {
		return com.Account
	}), nil
}

func (m *ManageCompanyService) RejectRegistrationForCompanies(companyAccounts []string) (err error) {
	var companiesToBeReviewed []entity.CompanyPendingReview
	if err = global.GVA_DB.Where("account IN ?", companyAccounts).Find(&companiesToBeReviewed).Error; err != nil {
		return err
	}

	if err = global.GVA_DB.Delete(&companiesToBeReviewed).Error; err != nil {
		return err
	}

	return nil
}

func (m *ManageCompanyService) AllowCompaniesInfoUpdate(companyAccounts []string) (allowAccounts []string, err error) {
	var companiesToBeUpdated []entity.CompanyInfoPendingReview
	var companiesApplictionRecord []recordEntity.ApplicationRecord
	if err = global.GVA_DB.Where("account IN ?", companyAccounts).Find(&companiesToBeUpdated).Error; err != nil {
		return
	}
	if err = global.GVA_DB.Where("account IN ? AND type = ? AND status = ?", companyAccounts, enum2.Info, enum2.UnderReview).Find(&companiesApplictionRecord).Error; err != nil {
		return nil, err
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

	recordToBeUpdate := functional.Map(companiesApplictionRecord, func(apr recordEntity.ApplicationRecord) recordEntity.ApplicationRecord {
		return recordEntity.ApplicationRecord{
			Id:      apr.Id,
			Account: apr.Account,
			Status:  enum2.PASSED,
			Type:    apr.Type,
			Date:    apr.Date,
		}
	})

	for _, record := range recordToBeUpdate {
		global.GVA_DB.Model(&record).Updates(record)
	}

	if err := global.GVA_DB.Delete(&companiesToBeUpdated).Error; err != nil {
		return nil, err
	}

	return functional.Map(companies, func(com entity.Company) string {
		return com.Account
	}), nil
}

func (m *ManageCompanyService) RejectCompaniesInfoUpdate(companyAccounts []string) (allowAccounts []string, err error) {
	var companiesToBeUpdated []entity.CompanyInfoPendingReview
	var companiesApplicationRecord []recordEntity.ApplicationRecord
	if err = global.GVA_DB.Where("account IN ?", companyAccounts).Find(&companiesToBeUpdated).Error; err != nil {
		return
	}
	if err = global.GVA_DB.Where("account IN ? AND type = ? AND status = ?", companyAccounts, enum2.Info, enum2.UnderReview).Find(&companiesApplicationRecord).Error; err != nil {
		return nil, err
	}

	companies := functional.Map(companiesToBeUpdated, func(com entity.CompanyInfoPendingReview) entity.Company {
		return entity.Company{
			Id:          com.Id,
			CompanyInfo: com.CompanyInfo,
		}
	})

	recordToBeUpdate := functional.Map(companiesApplicationRecord, func(apr recordEntity.ApplicationRecord) recordEntity.ApplicationRecord {
		return recordEntity.ApplicationRecord{
			Id:      apr.Id,
			Account: apr.Account,
			Status:  enum2.REJECTED,
			Type:    apr.Type,
			Date:    apr.Date,
		}
	})

	for _, record := range recordToBeUpdate {
		global.GVA_DB.Model(&record).Updates(record)
	}

	if err := global.GVA_DB.Delete(&companiesToBeUpdated).Error; err != nil {
		return nil, err
	}

	return functional.Map(companies, func(com entity.Company) string {
		return com.Account
	}), nil
}
