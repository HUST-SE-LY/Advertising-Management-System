package manage

import (
	"backend/global"
	"backend/models/admin_model"
	"backend/models/manage_model/request"
	"backend/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CRUD for admin_model

type ManageAdminService struct {
}

func (m *ManageAdminService) CreateAdmin(req *request.AdminCreateReq) (admin *admin_model.Admin, err error) {
	if !errors.Is(global.GVA_DB.Where("account = ?", req.Account).Take(&admin_model.Admin{}).Error, gorm.ErrRecordNotFound) {
		return nil, status.SameAccountExists
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	admin = &admin_model.Admin{
		Account:  req.Account,
		Password: string(encryptedPassword),
	}
	err = global.GVA_DB.Create(admin).Error
	return
}
