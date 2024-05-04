package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zoriya/vex"
)

func (h *Handler) AddFeed(c echo.Context) error {
	var feed vex.Feed
	err := c.Bind(&feed)
	if err != nil {
		return err
	}

	ret := make([]interface{}, 0)
	return c.JSON(201, ret)
}

func (h *Handler) RegisterFeedsRoutes(e *echo.Echo) {
	e.POST("/feed", h.AddFeed)
}
