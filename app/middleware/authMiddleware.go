package middleware

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"usergraphql/app/model/dto"
)

func AuthMiddleware() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// allow request without header Authorization
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenHeader := strings.Split(header, " ")
			tokenString := tokenHeader[len(tokenHeader)-1]

			_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET_KEY")), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusOK)

				response := map[string]any{
					"errors": &dto.ApiResponse[string]{
						StatusCode: http.StatusOK,
						Status:     "unauthorized",
						Message:    err.Error(),
					},
				}
				responseJwt, _ := json.Marshal(&response)
				w.Write(responseJwt)
				return
			}

			// success
			next.ServeHTTP(w, r)
		})
	}
}
