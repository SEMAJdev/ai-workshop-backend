package jwt

import (
	"errors"
	"fmt"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type HS256Signer struct {
	secret []byte
}

func NewHS256Signer(secret string) *HS256Signer {
	return &HS256Signer{secret: []byte(secret)}
}

func (s *HS256Signer) Sign(userID int64, expiresAt time.Time) (string, error) {
	claims := jwtv5.MapClaims{
		"sub": userID,
		"exp": expiresAt.Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *HS256Signer) Verify(token string) (int64, error) {
	parsed, err := jwtv5.Parse(token, func(t *jwtv5.Token) (interface{}, error) {
		if t.Method != jwtv5.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return s.secret, nil
	})
	if err != nil || !parsed.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsed.Claims.(jwtv5.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	// Accept numeric or string subject
	var id int64
	if v, ok := claims["sub"].(float64); ok {
		id = int64(v)
	} else if vStr, ok := claims["sub"].(string); ok {
		var parsedID int64
		if _, convErr := fmt.Sscan(vStr, &parsedID); convErr == nil {
			id = parsedID
		}
	}
	if id == 0 {
		return 0, errors.New("invalid subject")
	}
	return id, nil
}
