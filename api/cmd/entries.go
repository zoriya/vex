package main

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetEntries(c echo.Context) error {
	ret, err := h.entries.ListEntries()
	if err != nil {
		return err
	}
	return c.JSON(200, ret)
}

func (h *Handler) RegisterEntriesRoutes(e *echo.Echo, r *echo.Group) {
	e.GET("/entries", h.GetEntries)
}
