package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint //1表示普通用户, 2表示管理员
	jwt.StandardClaims
}
