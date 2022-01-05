package token

import (
	"fmt"
	"go-api-jwt/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var API_SECRET = utils.GetEnv("API_SECRET", "itsasecret")

func GenerateToken(user_id uint) (string, error) {
	token_lifespan, err := strconv.Atoi(utils.GetEnv("TOKEN_HOUR_LIFESPAN", "1"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return string of token that already signed with API_SECRET
	return token.SignedString([]byte(API_SECRET))
}

/* Extract token */
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}
	// berarerToken = "Bearer tokenstring"
	bearerToken := c.Request.Header.Get("Authorization")
	// check authorization header consist of 2 word (Bearer and tokenstring)
	// split it into array
	// get tokenstring from second element
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

/* Cek if token valid or not */
func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	// hanya untuk mencek apakah token valid
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// cek algoritma pengenkripisan token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}

	return nil
}

/* Extract User ID from token */
func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)

	// extract token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// cek algoritma pengenkripisan token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)

		if err != nil {
			return 0, err
		}

		return uint(uid), nil
	}

	// default value if all conditions fails
	return 0, nil
}
