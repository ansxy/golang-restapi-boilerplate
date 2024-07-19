package custome_jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type JWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateAccessToken(claims *UserClaims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "boiyler",
	}

	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenString, err := accesToken.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func GenerateRefreshToken(claims *UserClaims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		Issuer:    "boiyler",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshTokenString, err := refreshToken.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func GenerateToken(claims *UserClaims) (*JWT, error) {
	accessTokenString, err := GenerateAccessToken(claims)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := GenerateRefreshToken(claims)
	if err != nil {
		return nil, err
	}

	res := &JWT{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return res, nil
}

func GenerateTokenWithRefreshToken(token string) (*JWT, error) {
	valid, err := ValidateToken(token)
	if err != nil {
		return nil, err
	}

	accessTokenString, err := GenerateAccessToken(valid)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := GenerateRefreshToken(valid)
	if err != nil {
		return nil, err
	}

	res := &JWT{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return res, nil
}

func ValidateToken(token string) (*UserClaims, error) {
	claims := &UserClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
