package ecode

type Code interface {
	Error() string
	Code() uint32
	Message() string
	Equal(Code) bool
}

func New(code uint32, msg string) Code {
	return &_Code{
		code: code,
		msg:  msg,
	}
}
