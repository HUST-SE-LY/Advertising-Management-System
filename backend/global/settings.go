package global

type AppSetting struct {
	Address   string `json:"address"`
	JwtSecret string `json:"jwt_secret"`
}

type DatabaseSetting struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}
