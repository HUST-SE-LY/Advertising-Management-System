package manage

import (
	"backend/global"
	adminModel "backend/models/admin_model"
	"backend/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CRUD for admin_model

type ManageAdminService struct {
}

func (m *ManageAdminService) CreateAdmin(admin *adminModel.Admin) error {
	if !errors.Is(global.GVA_DB.Where("account = ?", admin.Account).First(&adminModel.Admin{}).Error, gorm.ErrRecordNotFound) {
		return status.SameAccountExists
	}
	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(password)
	err = global.GVA_DB.Create(&admin).Error
	return err
}
