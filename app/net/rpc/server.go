package rpc

import (
	"context"
	"fmt"
	"github.com/cnlisea/ant/typex"
	"reflect"
	"runtime"
	"strconv"
	"sync"

	"github.com/smallnest/rpcx/server"
)

type Server struct {
	Ip     string
	Port   uint16
	Server *server.Server

	handlerMeta reflect.Value
	handlers    sync.Map
	closed      bool
	mutex       sync.RWMutex
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
	if err := s.HandlerCache(handler); err != nil {
		return err
	}
	return s.Server.Register(handler, "")
}

func (s *Server) Addr() (string, uint16) {
	return s.Ip, s.Port
}

func (s *Server) HandlerCache(handler interface{}) error {
	methods, err := typex.ReflectSuitableMethods(reflect.TypeOf(handler), true)
	if err != nil {
		return err
	}

	for name, method := range methods {
		s.handlers.Store(name, method)
	}
	s.handlerMeta = reflect.ValueOf(handler)
	return nil
}
func (s *Server) HandlerCall(ctx context.Context, method string, args interface{}, reply interface{}) (exist bool, err error) {
	m, _ := s.handlers.Load(method)
	if m == nil {
		return
	}

	handler, ok := m.(*typex.ReflectMethodType)
	if !ok {
		s.handlers.Delete(method)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			buf = buf[:n]

			err = fmt.Errorf("[service internal error]: %v, method: %s, argv: %+v, stack: %s",
				r, method, args, buf)
		}
	}()

	if reply == nil {
		reply = typex.ReflectTypePools.Get(handler.ReplyType)
	}
	returnValues := handler.MethodCall().Call(
		[]reflect.Value{
			s.handlerMeta,
			reflect.ValueOf(ctx),
			reflect.ValueOf(args),
			reflect.ValueOf(reply),
		})
	errInter := returnValues[0].Interface()
	if errInter != nil {
		err = errInter.(error)
		return
	}

	exist = true
	return
}
