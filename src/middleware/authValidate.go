package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

func CheckAuth(roles ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			headerToken := r.Header.Get("Authorization")

			if !strings.Contains(headerToken, "Bearer") {
				libs.Respone("invalid header", 401, true).Send(w)
				return
			}

			token := strings.ReplaceAll(headerToken, "Bearer ", "")

			checkToken, err := libs.CheckToken(token)

			if err != nil {
				libs.Respone("invalid token", 401, true).Send(w)
				return
			}

			// roles := fmt.Sprint(res)
			// range strings.Split(roles, " ")
			var checkRole bool
			for _, v := range roles {
				if strings.ToLower(v) == strings.ToLower(checkToken.Role) {
					checkRole = true
					break
				}
			}

			if !checkRole {
				libs.Respone("unauthorized", 401, true).Send(w)
				return
			}

			ctx := context.WithValue(r.Context(), "username", checkToken.Username)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
