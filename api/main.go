package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct{}

func (h *Handler) GetEntries(c echo.Context) error {
	ret := make([]interface{}, 0)
	return c.JSON(200, ret)
}

func main() {
	h := Handler{}

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/entries", h.GetEntries)

	e.Start(":1597")
}
