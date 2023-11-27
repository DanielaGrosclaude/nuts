package models

type SecretModel struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	JWTSing  string `json:"jwtsing"`
	Database string `json:"database"`
}
