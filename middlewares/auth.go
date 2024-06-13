package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateAuthorization(authorized_users []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		//Get authorization token from cookies
		token_string := strings.Split(context.Request.Header.Get("Authorization"), " ")[1]

		// Parse the token
		token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, errors.New("error parsing token"))
		}

		// Read parsed token
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			//Check for expired token
			if float64(time.Now().Unix()) > claims["expiry"].(float64) {
				context.AbortWithError(http.StatusUnauthorized, errors.New("token expired"))
			} else {
				//Check for expected user type
				for _, authorized_user := range authorized_users {
					if claims["user_type"] == authorized_user {
						context.Next()
						return
					}
				}
				context.AbortWithError(http.StatusUnauthorized, errors.New("un-authorized user"))
			}
		} else {
			context.AbortWithError(http.StatusUnauthorized, errors.New("error parsing token"))
		}
	}
}
