package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/zoriya/vex"
)

type Handler struct {
	feeds vex.FeedService
}

func (h *Handler) GetEntries(c echo.Context) error {
	ret := make([]interface{}, 0)
	return c.JSON(200, ret)
}

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	con := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_SERVER"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sqlx.Open("postgres", con)
	if err != nil {
		log.Fatal(err)
	}
	h := Handler{
		feeds: vex.NewFeedService(db),
	}

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.GET("/entries", h.GetEntries)
	h.RegisterFeedsRoutes(e)

	e.Start(":1597")
}
