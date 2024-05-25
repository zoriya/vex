package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/zoriya/vex"
	"log"
	"net/http"
)

type AddFeedDto struct {
	Link string   `json:"link" validate:"required,url"`
	Tags []string `json:"tags" validate:"required"`
}

// @Tags         Feeds
// @Summary      Get Feeds
// @Produce      json
// @Success      200	{array} vex.Feed
// @Router       /feeds [get]
func (h *Handler) GetFeeds(c echo.Context) error {
	ret, err := h.feeds.ListFeeds()
	if err != nil {
		return err
	}
	return c.JSON(200, ret)
}

// @Tags         Feeds
// @Summary      Add a single Feed
// @Produce      json
// @Param        DTO body AddFeedDto true "The Feed to save"
// @Success      201    {object} vex.Feed
// @Failure      400    {object} string
// @Failure      409    {object} string
// @Router       /feeds [post]
// @Security JWT
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
