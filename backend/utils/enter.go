package utils

import (
	"backend/global"
	"backend/utils/jwt"
)

func Setup() {
	jwt.JwtSecret = []byte(global.GVA_APP_SETTING.JwtSecret)
}
