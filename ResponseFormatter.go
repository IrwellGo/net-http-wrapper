package net_http_wrapper

import (
	"encoding/json"
	"net/http"
)

type ResponseFormatter struct {
	http.ResponseWriter
}

func (rf ResponseFormatter) WithCookie(name string, value string) ResponseFormatter {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(rf, &cookie)

	return rf
}

func (rf ResponseFormatter) WithCode(code int) ResponseFormatter {
	rf.WriteHeader(code)

	return rf
}

func (rf ResponseFormatter) WithBodyAsJson(body any) ResponseFormatter {
	jsoner := json.NewEncoder(rf)
	jsoner.Encode(body)

	return rf
}

func (rf ResponseFormatter) WithJsonResponse(code int, body any) ResponseFormatter {
	rf.WithCode(code).WithBodyAsJson(body)

	return rf
}

func (rf ResponseFormatter) WithResponse(code int, body []byte) ResponseFormatter {
	rf.WithCode(code).Write(body)

	return rf
}

func (rf ResponseFormatter) attachCors(methods string, origin string) {
	if origin == "" {
		origin = "*"
	}

	rf.Header().Set("Access-Control-Allow-Origin", origin)
	rf.Header().Set("Access-Control-Allow-Methods", methods)
	rf.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
