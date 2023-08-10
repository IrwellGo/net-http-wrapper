package net_http_wrapper

import (
	"net/http"
	"strings"
)

type Cors struct {
	allowedMethods []string
	handler        http.Handler
}

func (c Cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := NewResponseFormatter(w)
	rw.AttachCors(strings.Join(c.allowedMethods, ","), "*")

	if r.Method == "OPTIONS" {
		rw.WithCode(http.StatusOK)
		return
	}

	c.handler.ServeHTTP(w, r)
}

func NewMiddlewareCors(allowedMethods []string, next http.Handler) *Cors {
	return &Cors{allowedMethods: allowedMethods, handler: next}
}
