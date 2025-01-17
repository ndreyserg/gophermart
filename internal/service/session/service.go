package session

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.SessionService = (*service)(nil)

func NewService(secret string) *service {
	return &service{
		secret: secret,
	}
}

type claims struct {
	jwt.RegisteredClaims
	UserID int
}

const tokenKey = "token"

type service struct {
	secret string
}

func (s *service) Open(userID int, w http.ResponseWriter, r *http.Request) error {
	strToken, err := s.newToken(userID)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     tokenKey,
		Value:    strToken,
		HttpOnly: true,
	})
	return nil
}

func (s *service) newToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		},
		UserID: userID,
	})

	res, err := token.SignedString([]byte(s.secret))

	if err != nil {
		return "", fmt.Errorf("new token error: %w", err)
	}

	return res, nil
}

func (s *service) GetUserID(r *http.Request) (int, error) {
	var strToken string

	cookies := r.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == tokenKey {
			strToken = cookie.Value
			break
		}
	}
	if strToken == "" {
		return 0, errors.New("empty token")
	}
	claims := &claims{}
	token, err := jwt.ParseWithClaims(strToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return 0, fmt.Errorf("token parse error: %w", err)
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.UserID, nil
}
