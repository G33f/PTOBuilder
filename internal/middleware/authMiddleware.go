package middleware

import (
	"PTOBuilder/pkg/logging"
	"net/http"
	"strings"
)

type checkFunc func(string) bool

type AuthMiddleware struct {
	log       *logging.Logger
	checkFunc checkFunc
}

func NewMiddleware(checkFunc func(string) bool, log *logging.Logger) AuthMiddleware {
	return AuthMiddleware{
		checkFunc: checkFunc,
		log:       log,
	}
}

func (m *AuthMiddleware) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			m.log.Info("Unauthorized user")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		reqToken = splitToken[1]
		if !m.checkFunc(reqToken) {
			m.log.Info("Unauthorized user")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
