package rpc

import (
	"strconv"
	"sync"

	"github.com/smallnest/rpcx/server"
)

type Server struct {
	Ip     string
	Port   uint16
	Server *server.Server

	closed bool
	mutex  sync.RWMutex
}

func NewServer(ip string, port uint16) *Server {
	return &Server{
		Ip:     ip,
		Port:   port,
		Server: server.NewServer(),
	}
}

func (s *Server) Run() error {
	if err := s.Server.Serve("tcp", s.Ip+":"+strconv.FormatUint(uint64(s.Port), 10)); err != nil && err != server.ErrServerClosed {
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
	return s.Server.Register(handler, "")
}

func (s *Server) Addr() (string, uint16) {
	return s.Ip, s.Port
}
