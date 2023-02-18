package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sp-yduck/kubedin/api"
)

func main() {
	// new instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// root
	e.Static("/", "dist")

	// api
	e.GET("/api/job", job)

	// serve
	// to do : use port number from env var
	e.Logger.Fatal(e.Start(":1234"))
}

func job(c echo.Context) error {
	jobs, _ := api.ListJobs()
	return c.JSON(http.StatusOK, jobs)
}
