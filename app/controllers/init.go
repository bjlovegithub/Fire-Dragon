package controllers

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"net/http"
	"os"
	"strings"
	"time"
)

func Authenticate(c *revel.Controller) revel.Result {
	jwtToken, err := getJWTToken(c)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Authentication Failed(no jwt token)"})
	}
	println("jwt token: " + jwtToken)

	var claims jwt.MapClaims
	claims, err = decodeToken(jwtToken)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(map[string]interface{}{"message": "Authentication Failed(invalid jwt token)"})
	}

	_, found := claims["email"]
	if !found {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Authentication Failed(email not found in db)"})
	}

	return nil
}

func getJWTToken(c *revel.Controller) (token string, err error) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("No auth in header")
	}

	arr := strings.Split(header, " ")
	if len(arr) != 2 {
		return "", errors.New("Invalid auth header")
	}

	return arr[1], nil
}

var hmacSecret = []byte(os.Getenv("JWT_SECRECT_KEY"))

func createJWT(info JWTInfo) string {
	// create a new jwt token based on the token from google.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   info.email,
		"exp":     time.Now().Unix() + 7*24*3600,
		"user_id": info.userId,
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		panic(err)
	}

	println(string(hmacSecret))

	return tokenString
}

func decodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println("email and nbf:", claims["email"], claims["nbf"])
	} else {
		return nil, err
	}
	return claims, nil
}

func init() {
	revel.InterceptFunc(Authenticate, revel.BEFORE, &WishApp{})
}
