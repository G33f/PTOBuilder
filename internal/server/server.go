package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	ip           string
	port         string
	writeTimeout time.Duration
	readTimeout  time.Duration
	serverHTTP   http.Server
}

func NewServer(router *httprouter.Router) *Server {
	s := Server{
		ip:           viper.GetString("WebServer.host"),
		port:         viper.GetString("WebServer.port"),
		writeTimeout: viper.GetDuration("WebServer.writeTimeout") * time.Second,
		readTimeout:  viper.GetDuration("WebServer.readTimeout") * time.Second,
	}
	s.createHTTPServer(router)
	return &s
}

func (s *Server) createHTTPServer(router *httprouter.Router) {
	s.serverHTTP = http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.ip, s.port),
		Handler:      router,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
	}
}

func (s *Server) Run() error {
	if err := s.serverHTTP.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
