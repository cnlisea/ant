package http

import (
	"errors"
	"net/http"
	"strconv"
	"sync"
)

type Server struct {
	Ip     string
	Port   uint16
	Server *http.Server

	closed bool
	mutex  sync.RWMutex
}

func NewServer(ip string, port uint16) *Server {
	return &Server{
		Ip:     ip,
		Port:   port,
		Server: new(http.Server),
	}
}

func (s *Server) Run() error {
	s.Server.Addr = s.Ip + ":" + strconv.FormatUint(uint64(s.Port), 10)
	if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Close() error {
	var err error
	s.mutex.Lock()
	if !s.closed {
		s.closed = true
		err = s.Server.Close()
	}
	s.mutex.Unlock()
	return err
}

func (s *Server) SetHandler(handler interface{}) error {
	h, ok := handler.(http.Handler)
	if !ok {
		return errors.New("handler invalid")
	}

	s.Server.Handler = h
	return nil
}

func (s *Server) Addr() (string, uint16) {
	return s.Ip, s.Port
}
