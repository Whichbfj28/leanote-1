package errcode

import (
	"context"

	"google.golang.org/grpc/codes"
)

func Mask(c int32, defaultMsg string) func(ctx context.Context, msg string, fields ...interface{}) CodeError {
	return func(ctx context.Context, msg string, fields ...interface{}) CodeError {
		if len(msg) == 0 {
			msg = defaultMsg
		}

		return Error(ctx, c, msg, fields...)
	}
}

func MaskStandard(c int32) func(ctx context.Context, msg string, fields ...interface{}) CodeError {
	return Mask(c, codes.Code(c).String())
}
