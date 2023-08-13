package hari_kesembilan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Word")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, string(body), "Hello Word")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello World")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestHttpQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/say-hello", nil)
	recorder := httptest.NewRecorder()

	sayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(body), "Hello World")

	request = httptest.NewRequest(http.MethodGet, "localhost:8080/say-hello?name=fuad", nil)
	recorder = httptest.NewRecorder()

	sayHello(recorder, request)

	response = recorder.Result()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(body), "Hello fuad")
}

func MultipelparamValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	fmt.Fprint(w, strings.Join(names, " "))
}

func TestHttpMultipelValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/say-hello?name=teuku&name=fuad&name=maulana", nil)
	recorder := httptest.NewRecorder()

	MultipelparamValues(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(body), "teuku fuad maulana")
}
