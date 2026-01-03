package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(username string, isAdmin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
	expiry    time.Duration
}

type jwtCustomClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func NewJWTService() JWTService {

	secretKey := os.Getenv("JWT_SECRET_KEY")
	issuer := os.Getenv("JWT_ISSUER")
	expiry_str := os.Getenv("JWT_EXPIRY_MINS") // in mins default to 15mins

	// parse expiry_str to time.Duration
	var expiry time.Duration
	if expiry_str != "" {
		expiryMinutes, err := time.ParseDuration(expiry_str + "m")
		if err != nil {
			expiry = 15 * time.Minute
		} else {
			expiry = expiryMinutes
		}
	} else {
		expiry = 15 * time.Minute
	}

	if secretKey != "" && issuer != "" {
		return &jwtService{
			secretKey: secretKey,
			issuer:    issuer,
			expiry:    expiry,
		}
	}

	return &jwtService{
		secretKey: "your-secret-key",
		issuer:    "your-app-name",
		expiry:    15 * time.Minute,
	}
}

func (s *jwtService) GenerateToken(username string, isAdmin bool) string {
	claims := &jwtCustomClaims{
		username,
		isAdmin,
		jwt.RegisteredClaims{
			Issuer:    s.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return ""
	}
	return t
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
