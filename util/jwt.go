package util

import (
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> fd910d7 (golang)
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("shudi")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}
type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	UserName      string `json:"user_name"`
	Email         string `json:"email" form:"email"`
	Passwrod      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	Authority     int    `json:"authority"`
	jwt.StandardClaims
}

//签发token

func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    "MALL-TEST",
			Subject:   "user token",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
<<<<<<< HEAD
	fmt.Println("=================1212121")
=======
>>>>>>> fd910d7 (golang)
	return nil, err
}
func GenerateEmailToken(userId, OperationType uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
<<<<<<< HEAD
	fmt.Println("11111")
=======
>>>>>>> fd910d7 (golang)
	claims := EmailClaims{
		UserID:        userId,
		Email:         email,
		Passwrod:      password,
		OperationType: OperationType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    "Email post",
			Subject:   "email token",
		},
	}
<<<<<<< HEAD
	fmt.Println("22222")
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("33333")
	token, err := tokenClaims.SignedString(jwtSecret)
	fmt.Println(err)
=======
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
>>>>>>> fd910d7 (golang)
	return token, err
}
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
