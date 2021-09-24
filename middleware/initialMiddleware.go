package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entityself"
	"github.com/google/uuid"
)

func InitializeContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		idTransaction := generateUUID(r)
		responseDefault := entityself.NewResponse(idTransaction)

		ctx := context.WithValue(r.Context(), constant.CtxKeyStartTransaction, start)
		ctx = context.WithValue(ctx, constant.CtxKeyIdTransaction, idTransaction)
		ctx = context.WithValue(ctx, constant.CtxKeyResponseDefault, responseDefault)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateUUID(req *http.Request) string {
	requestIDObject := uuid.Must(uuid.NewRandom())
	requestID := strings.Replace(fmt.Sprintf("%v", requestIDObject), "-", "", -1)
	return requestID
}
