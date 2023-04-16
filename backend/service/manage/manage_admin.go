package manage

import (
	"backend/global"
	"backend/models/admin_model/entity"
	"backend/models/manage_model/request"
	"backend/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CRUD for admin_model

type ManageAdminService struct {
}

func (m *ManageAdminService) CreateAdmin(req *request.AdminCreateReq) (admin *entity.Admin, err error) {
	if !errors.Is(global.GVA_DB.Where("account = ?", req.Account).Take(&entity.Admin{}).Error, gorm.ErrRecordNotFound) {
		return nil, status.SameAccountExists
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	admin = &entity.Admin{
		Account:  req.Account,
		Password: string(encryptedPassword),
	}
	err = global.GVA_DB.Create(admin).Error
	return
}
