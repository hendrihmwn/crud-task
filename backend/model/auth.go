package model

type LoginParam struct {
	Username string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
}

type AuthResponse struct {
	Token          string `json:"token"`
	ExpirationTime int64  `json:"expiration_time"`
}
