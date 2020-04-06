package main

import (
	handler "github.com/damogallagher/go-amazonses/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize handler
	h := &handler.Handler{}

	// routes - send email
	e.GET("/", h.SendEmail)

	e.Logger.Fatal(e.Start(":5000"))
}

