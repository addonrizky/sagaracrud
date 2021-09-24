package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/log"
	"io/ioutil"
	"net/http"
)

func ReadRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var ctx context.Context

		switch r.Method {
		case "POST":
			var i interface{}
			var body []byte
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&i)

			if err != nil {
				fmt.Println(err)
				fmt.Println("masuk sini???")
				log.LogDebug("-- fail read request body")
				return
			} else {
				body, _ = json.Marshal(i)
				// And now set a new body, which will simulate the same data we read:
				r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				requestBody := string(body)
				ctx = context.WithValue(r.Context(), constant.CtxKeyRequestBody, requestBody)
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
		case "GET":
			ctx = context.WithValue(r.Context(), constant.CtxKeyRequestBody, "")
		case "DELETE":
			ctx = context.WithValue(r.Context(), constant.CtxKeyRequestBody, "")
		}

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
