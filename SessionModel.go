package session

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type Session struct {
}

//Claims is the token that we will send user.
//Before sending the token, we will hash it by HS256 method.
//Remember, default session expires time is 24 hours unless you change it.

func (session Session) SetSession(c *fiber.Ctx, issuer string, secretKey string, expTime int) error {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewTime(1900000000),
	})
	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:     "authtoken",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * time.Duration(expTime)),
		HTTPOnly: true,
		Secure:   false,
	}
	c.Cookie(&cookie)
	return nil
}

//Gets the auth token from client
//if there is error, the user is not validate
//otherwise, the function returns Issuer (user id)
//you can use error to check if the token validate, the Issuer to have authorization system.

func (session Session) GetSession(c *fiber.Ctx, secretKey string) (string, error) {
	cookie := c.Cookies("authtoken")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	return token.Claims.(*jwt.StandardClaims).Issuer, nil
}
