package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Interface untuk JWT Service
type JWTService interface {
	GenerateToken(userID uint, role string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "default_secret"
	}
	return secret
}

// GenerateToken membuat token baru
func (s *jwtService) GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken
func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
}
