package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddFeedDto struct {
	Link string   `json:"link" validate:"required,url"`
	Tags []string `json:"tags" validate:"required"`
}

func (h *Handler) GetFeeds(c echo.Context) error {
	ret, err := h.feeds.ListFeeds()
	if err != nil {
		return err
	}
	return c.JSON(200, ret)
}

func (h *Handler) AddFeed(c echo.Context) error {
	user, err := GetCurrentUserId(c)
	if err != nil {
		return err
	}

	var req AddFeedDto
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&req); err != nil {
		return err
	}

	feed, err := h.feeds.AddFeed(req.Link, req.Tags, user)
	if err != nil {
		log.Printf("Add feed error: %v", err)
		return echo.NewHTTPError(500, "internal server error")
	}
	return c.JSON(201, feed)
}

func (h *Handler) RegisterFeedsRoutes(e *echo.Echo, r *echo.Group) {
	e.GET("/feeds", h.GetFeeds)
	r.POST("/feeds", h.AddFeed)
}
