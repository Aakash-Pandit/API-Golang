package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		authHeader := strings.Split(request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(map[string]string{"detail": "Unauthorized"})
			return

		} else {

			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(GoDotEnvVariable("ACCESS_SECRET")), nil
			})

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				handler.ServeHTTP(response, request)
			} else {
				fmt.Println(err)
				response.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(response).Encode(map[string]string{"detail": "Unauthorized"})
				return
			}
		}
	}
}
