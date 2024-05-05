package main

import (
	"fmt"
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

	feeds, err := h.feeds.GetFeedData(req.Link)
	if err != nil {
		return echo.NewHTTPError(400, fmt.Sprintf("Invalid feed link: %v", err))
	}
	if len(feeds) != 1 {
		return c.JSON(409, feeds)
	}
	feed := feeds[0]
	feed.SubmitterId = user
	feed.Tags = req.Tags
	feed, err = h.feeds.AddFeed(feed)
	if err != nil {
		log.Printf("Add feed error: %v", err)
		return echo.NewHTTPError(500, "internal server error")
	}
	h.sync.SyncFeed(feed)
	return c.JSON(201, feed)
}

func (h *Handler) RegisterFeedsRoutes(e *echo.Echo, r *echo.Group) {
	e.GET("/feeds", h.GetFeeds)
	r.POST("/feeds", h.AddFeed)
}
