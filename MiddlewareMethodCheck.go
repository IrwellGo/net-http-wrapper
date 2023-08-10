package net_http_wrapper

import "net/http"

type MiddlewareCors struct {
	method string
	next   http.Handler
}

func (c MiddlewareCors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := NewResponseFormatter(w)
	if r.Method != c.method {
		rw.WithCode(http.StatusMethodNotAllowed)
		return
	}

	c.next.ServeHTTP(w, r)
}

func NewMiddlewareMethodCheck(method string, next http.Handler) *MiddlewareCors {
	return &MiddlewareCors{method: method, next: next}
}
