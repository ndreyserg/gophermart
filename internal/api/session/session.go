package session

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Session interface {
	Open(userID int, w http.ResponseWriter, r *http.Request) error
	GetUserID(r *http.Request) (string, error)
}

func NewSession(secret string) Session {
	return &session{
		secret: secret,
	}
}

type claims struct {
	jwt.RegisteredClaims
	UserID int
}

const tokenKey = "token"

type session struct {
	secret string
}

func (s *session) Open(userID int, w http.ResponseWriter, r *http.Request) error {
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

func (s *session) newToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		},
		UserID: userID,
	})
	return token.SignedString([]byte(s.secret))
}

func (s *session) GetUserID(r *http.Request) (string, error) {
	return "", nil
}
