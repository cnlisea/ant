package typex

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unicode"
	"unicode/utf8"
)

// Precompute the reflect type for error. Can't use error directly
// because Typeof takes an empty interface value. This is annoying.
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

// Precompute the reflect type for context.
var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()

type ReflectMethodType struct {
	sync.Mutex // protects counters
	method     reflect.Method
	ArgType    reflect.Type
	ReplyType  reflect.Type
}

func (rmt *ReflectMethodType) MethodCall() reflect.Value {
	return rmt.method.Func
}

func ReflectIsExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

func ReflectIsExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return ReflectIsExported(t.Name()) || t.PkgPath() == ""
}

// suitableMethods returns suitable Rpc methods of typ, it will report
// error using log if reportErr is true.
func ReflectSuitableMethods(typ reflect.Type, reportErr bool) (map[string]*ReflectMethodType, error) {
	methods := make(map[string]*ReflectMethodType)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mType := method.Type
		mName := method.Name
		// Method must be exported.
		if method.PkgPath != "" {
			continue
		}
		// Method needs four ins: receiver, context.Context, *args, *reply.
		if mType.NumIn() != 4 {
			if reportErr {
				return nil, errors.New(fmt.Sprintln("method ", mName, " has wrong number of ins:", mType.NumIn()))
			}
			continue
		}
		// First arg must be context.Context
		ctxType := mType.In(1)
		if !ctxType.Implements(typeOfContext) {
			if reportErr {
				return nil, errors.New(fmt.Sprintln("method ", mName, " must use context.Context as the first parameter"))
			}
			continue
		}

		// Second arg need not be a pointer.
		argType := mType.In(2)
		if !ReflectIsExportedOrBuiltinType(argType) {
			if reportErr {
				return nil, errors.New(fmt.Sprintln(mName, " parameter type not exported: ", argType))
			}
			continue
		}
		// Third arg must be a pointer.
		replyType := mType.In(3)
		if replyType.Kind() != reflect.Ptr {
			if reportErr {
				return nil, errors.New(fmt.Sprintln("method", mName, " reply type not a pointer:", replyType))
			}
			continue
		}
		// Reply type must be exported.
		if !ReflectIsExportedOrBuiltinType(replyType) {
			if reportErr {
				return nil, errors.New(fmt.Sprintln("method", mName, " reply type not exported:", replyType))
			}
			continue
		}
		// Method needs one out.
		if mType.NumOut() != 1 {
			if reportErr {
				return nil, errors.New(fmt.Sprintln("method", mName, " has wrong number of outs:", mType.NumOut()))
			}
			continue
		}
		// The return type of the method must be error.
		if returnType := mType.Out(0); returnType != typeOfError {
			if reportErr {
				return nil, errors.New(fmt.Sprintln("method", mName, " returns ", returnType.String(), " not error"))
			}
			continue
		}
		methods[mName] = &ReflectMethodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}

		// init pool for reflect.Type of args and reply
		ReflectTypePools.Init(argType)
		ReflectTypePools.Init(replyType)
	}
	return methods, nil
}
