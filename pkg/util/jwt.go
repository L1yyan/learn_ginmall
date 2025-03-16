package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("liyanbebetter")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

type EmailClaims struct {
	UserID uint `json:"user_id"`
	Email string `json:"email"`
	Password string `json:"password"`
	OperationType uint `json:"operation_type"`
	jwt.StandardClaims
}
//签发Emailtoken
func GenerateEmailToken(userId, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24*time.Hour)
	emailClaims := EmailClaims {
		UserID: userId,
		Email: email,
		Password: password,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expireTime.Unix(),
			Issuer: "liyan",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,emailClaims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//验证用户Emailtoken

func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if emailClaims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return emailClaims, nil
		}
	}
	return nil, err
}

//签发token
func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24*time.Hour)
	claims := Claims {
		ID: id,
		UserName: userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expireTime.Unix(),
			Issuer: "liyan",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}