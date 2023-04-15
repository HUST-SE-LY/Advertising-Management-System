package manage

import (
	"backend/global"
	"backend/models/admin_model"
	"backend/models/admin_model/request"
	"backend/utils/jwt"
	"backend/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CRUD for admin_model

type ManageAdminService struct {
}

func (m *ManageAdminService) CreateAdmin(req request.AdminRegisterReq) (admin *admin_model.Admin, err error) {
	if !errors.Is(global.GVA_DB.Where("account = ?", req.Account).First(&admin_model.Admin{}).Error, gorm.ErrRecordNotFound) {
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

func (m *ManageAdminService) AdminLogin(req *request.AdminLoginReq) (err error, admin *admin_model.Admin, adminToken *admin_model.AdminToken) {
	err = global.GVA_DB.Where("account = ?", req.Account).Take(admin).Error
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
		global.GVA_DB.Take(adminToken, admin.Id)
		//global.GVA_DB.Where("admin_id = ?", admin.Id).Take(&adminToken)

		// There was no token stored in the database before.
		if adminToken == nil {
			adminToken = &admin_model.AdminToken{
				AdminId: admin.Id,
				Token:   token,
			}
			if err = global.GVA_DB.Create(adminToken).Error; err != nil {
				return
			}
		} else {
			adminToken.Token = token
			if err = global.GVA_DB.Save(adminToken).Error; err != nil {
				return
			}
		}
	}
	return err, admin, adminToken
}

func (m *ManageAdminService) FindAdminToken(token string) (adminToken *admin_model.AdminToken, err error) {
	err = global.GVA_DB.Take(&adminToken, "token = ?", token).Error
	return
}

func (m *ManageAdminService) DeleteAdminToken(token string) error {
	err := global.GVA_DB.Delete(&[]admin_model.AdminToken{}, "token = ?", token).Error
	return err
}

func getNewToken(account, password string) (token string) {
	token, _ = jwt.GenerateToken(account, password, admin_model.ADMIN_ROLE)
	return
}
