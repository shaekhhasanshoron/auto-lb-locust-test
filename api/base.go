package api

import (
	"github.com/labstack/echo/v4"
	v1 "github.com/locust-test/api/v1"
	"net/http"
)

// Routes base router
func Routes(e *echo.Echo) {
	// Index Page
	e.GET("/", index)

	// Health Page
	e.GET("/health", health)
	v1.Router(e.Group("/api/v1"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is KloverCloud CI Service")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
