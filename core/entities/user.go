package entities

import "github.com/golang-jwt/jwt/v4"

type User struct {
	UserName    string `form:"user_name"`
	DisplayName string `form:"display_name"`
	Password    string `form:"password"`
	Avatar      string `form:"avatar"`
	GoogleID    string `form:"google_id"`
	CreateToken int64  `form:"create_token"`
}
type LoginRequest struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}
type ReferenceTokens struct {
	ExpireAccess int64 `json:"expire_access"`
}
type Session struct {
	Token  string `json:"token"`
	Device string `json:"device"`
}
type LogoutRequest struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
type UserJwtClaim struct {
	*jwt.StandardClaims
	UserName string `json:"user_name"`
}
