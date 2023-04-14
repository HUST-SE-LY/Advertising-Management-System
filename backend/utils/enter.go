package utils

import (
	"backend/global/settings"
	"backend/utils/jwt"
)

func Setup() {
	jwt.jwtSecret = []byte(settings.GetAppSetting().JwtSecret)
}
