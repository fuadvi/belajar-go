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

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstname := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Teuku&lastName=Fuad")
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/users", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, string(body), "Hello Teuku Fuad")
}
