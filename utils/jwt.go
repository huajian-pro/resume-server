package utils

import (
	"resume-server/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte(conf.Cfg.JWT.Secret)

type Claims struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Status   int    `json:"status"`
	Phone    string `json:"phone"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userid, username, email, phone string, role, status int) (string, error) {
	expireTime := time.Now().Add(time.Duration(conf.Cfg.JWT.ExpireTime) * time.Hour)
	claims := Claims{
		Userid:   userid,
		Username: username,
		Email:    email,
		Role:     role,
		Phone:    phone,
		Status:   status,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    conf.Cfg.JWT.Issuer,
			Subject:   conf.Cfg.JWT.Subject,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
