package admin

import (
	"backend/global"
	"backend/models/admin_model/entity"
	"backend/models/admin_model/request"
	"backend/utils/jwt"
	"backend/utils/status"
	"golang.org/x/crypto/bcrypt"
)

type AdminAccountService struct {
}

func (a *AdminAccountService) AdminLogin(req *request.AdminLoginReq) (err error, admin *entity.Admin, adminToken *entity.AdminToken) {
	err = global.GVA_DB.Where("account = ?", req.Account).Take(&admin).Error
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
		global.GVA_DB.Take(&adminToken, admin.Id)
		//global.GVA_DB.Where("admin_id = ?", admin.Id).Take(&adminToken)

		// There was no token stored in the database before.
		if adminToken == nil {
			adminToken = &entity.AdminToken{
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

func (a *AdminAccountService) FindAdminToken(token string) (adminToken *entity.AdminToken, err error) {
	err = global.GVA_DB.Take(&adminToken, "token = ?", token).Error
	return
}

func (a *AdminAccountService) DeleteAdminToken(token string) error {
	err := global.GVA_DB.Delete(&[]entity.AdminToken{}, "token = ?", token).Error
	return err
}

func getNewToken(account, password string) (token string) {
	token, _ = jwt.GenerateToken(account, password, entity.ADMIN_ROLE)
	return
}
