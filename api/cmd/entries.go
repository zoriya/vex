package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zoriya/vex"
)

type ChangeEntryStatusDto struct {
	Id           uuid.UUID `json:"id" validate:"required"`
	IsRead       bool      `json:"isRead"`
	IsBookmarked bool      `json:"isBookmarked"`
	IsReadLater  bool      `json:"isReadLater"`
	IsIgnored    bool      `json:"isIgnored"`
}

// @Tags         Entries
// @Summary      Get entries
// @Produce      json
// @Success      200    {object} vex.Entry
// @Router       /entries [get]
// @Security JWT
func (h *Handler) GetEntries(c echo.Context) error {
	user, err := GetCurrentUserId(c)
	if err != nil {
		return err
	}
	ret, err := h.entries.ListEntries(user)
	if err != nil {
		return err
	}
	return c.JSON(200, ret)
}

// @Tags         Entries
// @Summary      Change status
// @Produce      json
// @Param        DTO body ChangeEntryStatusDto true " "
// @Success      200    {object} vex.Entry
// @Router       /entries [patch]
// @Security JWT
func (h *Handler) ChangeUserStatus(c echo.Context) error {
	user, err := GetCurrentUserId(c)
	if err != nil {
		return err
	}

	var req ChangeEntryStatusDto
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&req); err != nil {
		return err
	}

	err = h.entries.ChangeStatus(vex.ChangeStatusDao{
		Id:           req.Id,
		User:         user,
		IsRead:       req.IsRead,
		IsBookmarked: req.IsBookmarked,
		IsReadLater:  req.IsReadLater,
		IsIgnored:    req.IsIgnored,
	})
	if err != nil {
		return err
	}
	ret, err := h.entries.GetEntry(req.Id, user)
	if err != nil {
		return err
	}
	return c.JSON(200, ret)
}

func (h *Handler) RegisterEntriesRoutes(e *echo.Echo, r *echo.Group) {
	r.GET("/entries", h.GetEntries)
	r.PATCH("/entries", h.ChangeUserStatus)
}
