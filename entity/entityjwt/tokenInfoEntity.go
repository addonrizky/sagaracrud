package entityjwt

import jwt "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
	TypeUser string `json:"type_user"`
}
