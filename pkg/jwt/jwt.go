package custome_jwt

type UserClaims struct {
}

type JWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateToken(claims *UserClaims) (*JWT, error) {
	return nil, nil
}

func ValidateToken(token string) (*UserClaims, error) {
	return nil, nil
}
