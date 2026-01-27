package global

import (
	"context"
	"errors"
)

var (
	ctx    context.Context
	cancel context.CancelCauseFunc
)

func init() {
	ctx, cancel = context.WithCancelCause(context.Background())
}

func Context() context.Context {
	return ctx
}

func Shutdown() {
	cancel(errors.New("shutdown"))
}
