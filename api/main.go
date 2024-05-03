package api

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (h *Handler) GetEntries(c echo.Context) error {
	ret := make([]interface{}, 0)
	return c.JSON(200, ret)
}

func main() {
	h := Handler{}

	e := echo.New()
	e.GET("/entries", h.GetEntries)

	log.Print("Listening on :1597")
	e.Start(":1597")
}
