package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"

)

var JWT_SECRET string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		token := strings.Split(header, " ")

		if len(token) != 2 || token[1] == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenRetrieved := token[1]
		_, err := jwt.Parse(tokenRetrieved, func (token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWT_SECRET), nil
		})
		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			log.Println(err)
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func DeliverToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	return tokenString, err
}