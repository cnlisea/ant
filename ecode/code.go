package ecode

type _Code struct {
	code uint32
	msg  string
}

func (s _Code) Error() string {
	return s.msg
}

func (s _Code) Code() uint32 {
	return s.code
}

func (s _Code) Message() string {
	return s.msg
}

func (s _Code) Equal(err Code) bool {
	return s.code == err.Code()
}
