package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	// 初始化错误日志引起和配置文件
	_ "github.com/qnfnypen/changuan/internal/pkg"
)

// 注意：这里要用byte切片
var jwtSecret = []byte("#$&test!@*")

// Claims jwt的payload部分字段
type Claims struct {
	UID      string `json:"uid"`      // 用户唯一标识
	Password string `json:"password"` // 登录校验码: 密码 验证码(留空)，用于用户修改密码后注销之前登录
	MID      string `json:"mid"`      // 机器码，用于限制同时登录人数
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(uid, password, mid string) (string, error) {
	claims := Claims{uid, password, mid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(viper.GetDuration("HTTP.TokenExpDur") * time.Hour * 24).Unix(),
			Issuer:    "changuan",
		},
	}

	// 注意，这里加密方式要选HS256
	tc := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tc.SignedString(jwtSecret)
}

// ParseToken 解析token
func ParseToken(token string) *Claims {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims
	}

	return nil
}
