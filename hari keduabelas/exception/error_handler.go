package exception

import (
	"belajar-resfull-api/helper"
	"belajar-resfull-api/model/web"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(write http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(write, request, err) {
		return
	}

	if validatorError(write, request, err) {
		return
	}

	internalServerError(write, request, err)
}

func validatorError(write http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		write.Header().Set("Content-type", "application/json")
		write.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(write, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(write http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	fmt.Println(ok)
	if ok {
		write.Header().Set("Content-type", "application/json")
		write.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(write, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(write http.ResponseWriter, request *http.Request, err interface{}) {
	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Status Error",
		Data:   err,
	}

	helper.WriteToResponseBody(write, webResponse)
}
