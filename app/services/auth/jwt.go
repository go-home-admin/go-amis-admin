package auth

import (
	"github.com/go-home-admin/home/app"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

// Jwt @Bean
type Jwt struct {
	secretKey []byte
}

func (j *Jwt) Init() {
	j.secretKey = []byte(app.Config("auth.secret", "test"))
}

func (j *Jwt) GenerateToken(id int) (string, error) {
	// 设置过期时间为 24 小时
	expireTime := time.Now().Add(24 * time.Hour)
	// 创建一个自定义的 Claims 结构体
	claims := MyClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "my_app",
		},
	}
	// 使用 HS256 算法创建一个新的 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用 secretKey 对 Token 进行签名并返回字符串格式
	return token.SignedString(j.secretKey)
}

func (j *Jwt) GetUid(token string) (int, error) {
	var claims MyClaims
	NToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := NToken.Claims.(*MyClaims); ok && NToken.Valid {
		return claims.ID, err
	} else {
		return 0, err
	}
}
