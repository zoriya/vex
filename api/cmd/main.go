package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/zoriya/vex"
)

type Handler struct {
	feeds     vex.FeedService
	entries   vex.EntryService
	users     vex.UserService
	sync      vex.SyncService
	jwtSecret []byte
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

	reader := vex.NewRssReader(http.DefaultClient)
	h := Handler{
		feeds:     vex.NewFeedService(db, &reader),
		entries:   vex.NewEntryService(db),
		users:     vex.NewUserService(db),
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}
	h.sync = vex.NewSyncService(&reader, &h.feeds, &h.entries)
	go h.sync.SyncFeedsForever()

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Use(middleware.Logger())

	r := e.Group("")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: h.jwtSecret,
	}))

	e.GET("/entries", h.GetEntries)
	h.RegisterLoginRoutes(e, r)
	h.RegisterEntriesRoutes(e, r)
	h.RegisterFeedsRoutes(e, r)

	e.Start(":1597")
}
