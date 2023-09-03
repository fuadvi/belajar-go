package middleware

import (
	"belajar-resfull-api/helper"
	"belajar-resfull-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handle http.Handler
}

func NewAuthMiddleware(handle http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handle: handle}
}

func (middlaware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "Rahasia" == request.Header.Get("X-API-KEY") {
		middlaware.Handle.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
