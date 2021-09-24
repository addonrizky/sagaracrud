package utility

import (
	"fmt"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entityjwt"
	"time"
	//"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
)

type M map[string]interface{}

func EncodeTokenJwt(username string, email string, typeUser string) string {
	claims := entityjwt.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    constant.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(constant.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: username,
		Email:    email,
		TypeUser: typeUser,
	}

	token := jwt.NewWithClaims(
		constant.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(constant.JWT_SIGNATURE_KEY)
	if err != nil {
		fmt.Println(err)
	}

	return string(signedToken)
}

func Decodejwt(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != constant.JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return constant.JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token claim invalid")
	}

	return claims, nil
}
