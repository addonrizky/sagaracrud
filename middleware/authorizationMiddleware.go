package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entityjwt"
	"github.com/addonrizky/sagaracrud/utility"
	"net/http"
	"strings"
)

func AuthorizationJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		var tokenString string
		var tokenInfo entityjwt.MyClaims
		switch r.Method {
		case "POST":
			if r.URL.Path == "/login" {
				next.ServeHTTP(w, r)
				return
			}

			authorizationHeader := r.Header.Get("Authorization")
			if !strings.Contains(authorizationHeader, "Bearer") {
				http.Error(w, "Invalid token", http.StatusBadRequest)
				return
			}
			tokenString = strings.Replace(authorizationHeader, "Bearer ", "", -1)
		case "GET":
			tokenString = r.URL.Query().Get("tok")
			if tokenString == "" {
				fmt.Println("go to hell")
				return
			}
		case "DELETE":
			tokenString = r.URL.Query().Get("tok")
			if tokenString == "" {
				fmt.Println("go to hell")
				return
			}
		}

		tokenClaim, err := utility.Decodejwt(tokenString)

		if err != nil {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenInfoJson, err := json.Marshal(tokenClaim)

		if err != nil {
			http.Error(w, "Invalid token, fail to encode to json", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(tokenInfoJson, &tokenInfo)

		if err != nil {
			http.Error(w, "Invalid token, fail to parse to object", http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(r.Context(), constant.CtxUsername, tokenInfo.Username)
		ctx = context.WithValue(ctx, constant.CtxEmail, tokenInfo.Email)
		ctx = context.WithValue(ctx, constant.CtxTypeUser, tokenInfo.TypeUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
