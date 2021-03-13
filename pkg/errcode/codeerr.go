package errcode

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CodeError interface {
	Code() int32
	Error() string
}

type codeErr struct {
	code    int32
	fields  interface{}
	message string
}

func (c codeErr) Code() int32 {
	return c.code
}

func (c codeErr) Error() string {
	return status.Errorf(codes.Code(c.code), fmt.Sprint(map[string]interface{}{"message": c.message, "fields": c.fields})).Error()
}

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(_ context.Context, c int32, msg string, fields ...interface{}) CodeError {
	if codes.Code(c) == codes.OK {
		return nil
	}

	return codeErr{code: c, message: msg, fields: fields}
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(ctx context.Context, c int32, format string, a ...interface{}) error {
	return Error(ctx, c, fmt.Sprintf(format, a...))
}
