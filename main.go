package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", root)
	e.GET("/callback", hello)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// Handler
func hello(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")
	url := "https://wear.googleapis.com/3p_auth/app.html?full_suffix_from_redirect=" +
		"com.iwmh.spotifywearapp" + "?code=" + code + "&state=" + state
	return c.Redirect(http.StatusMovedPermanently, url)
}

func root(c echo.Context) error {
	value := c.QueryParams()
	_ = value
	return c.String(http.StatusOK, "default")
}
