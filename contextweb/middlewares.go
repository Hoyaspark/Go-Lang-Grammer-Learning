package main

import (
	"net/http"
)

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}

	return userCookie.Value, nil
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)

		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = ContextWithUser(ctx, user)
		r = r.WithContext(ctx)

		h.ServeHTTP(rw, r)
	})
}
