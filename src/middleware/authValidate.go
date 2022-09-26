package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			libs.Respone("invalid header type", 401, true).Send(w)
			return
		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkTokens, err := libs.CheckToken(token)
		if err != nil {
			libs.Respone(err.Error(), 401, true).Send(w)
			return
		}
		ctx := context.WithValue(r.Context(), "username", checkTokens.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
func AdminRole(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			libs.Respone("invalid header type", 401, true).Send(w)
			return
		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkTokens, err := libs.CheckToken(token)
		if err != nil {
			libs.Respone(err.Error(), 401, true).Send(w)
			return
		}

		if checkTokens.Role != "admin" {
			libs.Respone("anda bukan admin", 401, true).Send(w)
			return

		}
		ctx := context.WithValue(r.Context(), "role", checkTokens.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
