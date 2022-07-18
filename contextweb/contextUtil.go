package main

import "context"

type ctxKey int

const (
	userKey ctxKey = iota
)

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(userKey).(string)
	return v, ok
}
