package manage

import (
	"backend/global"
	adminModel "backend/models/admin_model"
	"backend/models/admin_model/request"
	"backend/utils/jwt"
	"backend/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
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

func (m *ManageAdminService) AdminLogin(req *request.AdminLoginReq) (err error, admin *adminModel.Admin, adminToken *adminModel.AdminToken) {
	if err != nil {
		return
	}

	err = global.GVA_DB.Where("account = ?", req.Account).First(&admin).Error
	if admin == nil {
		err = status.AccountNotFound
		return
	} else {
		// Check password
		if err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
			err = status.WrongPassword
			return
		}

		token := getNewToken(admin.Account, admin.Password)
		global.GVA_DB.Where("admin_id = ?", admin.Id).First(&adminToken)

		// There was no token stored in the database before.
		if adminToken == nil {
			adminToken = &adminModel.AdminToken{
				AdminId: admin.Id,
				Token:   token,
			}
			if err = global.GVA_DB.Create(&adminToken).Error; err != nil {
				return
			}
		} else {
			adminToken.Token = token
			log.Println(adminToken.AdminId)
			if err = global.GVA_DB.Save(&adminToken).Error; err != nil {
				return
			}
		}
	}
	return err, admin, adminToken
}

func getNewToken(account string, password string) (token string) {
	token, _ = jwt.GenerateToken(account, password, adminModel.ADMIN_ROLE)
	return
}
