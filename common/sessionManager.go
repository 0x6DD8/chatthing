package common

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

const SESSION_COOKIE_NAME = "session"

func ManageSession(w http.ResponseWriter, req *http.Request) {
	var cookie, err = req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name:    SESSION_COOKIE_NAME,
			Value:   generateSessionID(),
			Expires: time.Now().Add(time.Hour * 24 * 365),
		}
		http.SetCookie(w, cookie)
	}
}

func GetSessionID(req *http.Request) string {
	var cookie, err = req.Cookie(SESSION_COOKIE_NAME)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func generateSessionID() string {
	return uuid.New().String()
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ManageSession(w, req)
		next.ServeHTTP(w, req)
	})
}
