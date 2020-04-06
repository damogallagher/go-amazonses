package handler

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

//SendEmail is a handler method to send an email
//e.GET("/", h.SendEmail)
func (h *Handler) SendEmail(c echo.Context) (err error) {
	// build the query based on params passed in the request
	fmt.Println("SendEmail request body:", c.QueryParams())



	c.JSON(http.StatusOK, "Sending email")

	return nil
}