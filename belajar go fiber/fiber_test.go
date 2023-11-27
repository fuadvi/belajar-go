package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutingHelloWorld(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(byte))
}

func TestCtx(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		name := ctx.Query("name", "quest")
		return ctx.SendString("Hello " + name)
	})

	request := httptest.NewRequest("GET", "/?name=fuad", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello fuad", string(byte))

	request = httptest.NewRequest("GET", "/", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello quest", string(byte))
}

func TestRouteParam(t *testing.T) {
	app := fiber.New()
	app.Get("/order/:orderId/:userId", func(ctx *fiber.Ctx) error {
		orderId := ctx.Params("orderId")
		userId := ctx.Params("userId")
		return ctx.SendString("Transaksi order " + orderId + " dari user " + userId)
	})

	request := httptest.NewRequest("GET", "/order/1/2", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Transaksi order 1 dari user 2", string(byte))
}

func TestFormRequest(t *testing.T) {
	app := fiber.New()
	app.Post("/hello", func(ctx *fiber.Ctx) error {
		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name)
	})

	body := strings.NewReader("name=fuad")
	request := httptest.NewRequest("POST", "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello fuad", string(byte))
}
