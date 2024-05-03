package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type Handler struct {
	database *sql.DB
}

func (h *Handler) GetEntries(c echo.Context) error {
	ret := make([]interface{}, 0)
	return c.JSON(200, ret)
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
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal(err)
	}
	h := Handler{
		database: db,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/entries", h.GetEntries)
	h.RegisterFeedsRoutes(e)

	e.Start(":1597")
}
