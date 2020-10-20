package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/soganhei.com.br/atlas"
)

type (
	Services struct {
		secret []byte
	}

	claims struct {
		IDUser   int64  `json:"id_user,omitempty"`
		NameUser string `json:"name_user"`
		Role     int    `json:"role"`
		jwt.StandardClaims
	}
)

func NewServices() (*Services, error) {

	ts := os.Getenv("TOKEN_SECRET")

	secret := []byte(ts)

	service := &Services{
		secret,
	}

	return service, nil

}

// GenerateToken generates a new JWT token.
func (service *Services) GenerateToken(payload atlas.TokenData) (string, error) {

	expireToken := time.Now().Add(time.Hour * 8).Unix()

	cl := claims{
		NameUser: payload.NameUser,
		IDUser:   payload.IDUser,
		Role:     payload.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)

	signedToken, err := token.SignedString(service.secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ParseAndVerifyToken parses a JWT token and verify its validity. It returns an error if token is invalid.
func (service *Services) ParseAndVerifyToken(token string) (*atlas.TokenData, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, msg
		}
		return service.secret, nil
	})
	if err == nil && parsedToken != nil {

		if cl, ok := parsedToken.Claims.(*claims); ok && parsedToken.Valid {

			tokenData := &atlas.TokenData{
				NameUser: cl.NameUser,
				IDUser:   cl.IDUser,
				Role:     cl.Role,
			}
			return tokenData, nil
		}
	}

	return nil, errors.New("invalid token jwt")
}
