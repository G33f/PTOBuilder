package middleware

import (
	"PTOBuilder/pkg/logging"
	"net/http"
	"strings"
)

type checkAuth func(string) bool
type checkAdmin func(string) bool

type AuthMiddleware struct {
	log        *logging.Logger
	checkAuth  checkAuth
	checkAdmin checkAdmin
}

func NewMiddleware(checkAuth func(string) bool, checkAdmin func(string) bool, log *logging.Logger) AuthMiddleware {
	return AuthMiddleware{
		checkAuth:  checkAuth,
		checkAdmin: checkAdmin,
		log:        log,
	}
}

func (m *AuthMiddleware) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			m.log.Info("Unauthorized user")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		reqToken = splitToken[1]
		if !m.checkAuth(reqToken) {
			m.log.Info("Unauthorized user")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) CheckAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			m.log.Info("Unauthorized user")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		reqToken = splitToken[1]
		if !m.checkAdmin(reqToken) {
			m.log.Info("The user does not have admin rights")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
