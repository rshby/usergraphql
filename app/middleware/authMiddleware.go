package middleware

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"log"
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

			log.Println("middleware : token string :", tokenString)

			claims := jwt.MapClaims{}
			secretKey := os.Getenv("SECRET_KEY")
			log.Println(secretKey)
			token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if _, ok := token.Claims.(jwt.MapClaims); !ok {
				w.WriteHeader(http.StatusOK)

				response := map[string]any{
					"errors": &dto.ApiResponse[string]{
						StatusCode: http.StatusOK,
						Status:     "unauthorized",
						Message:    "token not valid",
					},
				}
				responseJwt, _ := json.Marshal(&response)
				w.Write(responseJwt)
				return
			}

			log.Println("middleware : masuk ini")

			// success
			next.ServeHTTP(w, r)
		})
	}
}
