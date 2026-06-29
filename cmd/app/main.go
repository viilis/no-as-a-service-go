package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Println(port)

	e := echo.New()

	limiterStore := middleware.NewRateLimiterMemoryStore(20)

	e.Use(middleware.Recover())
	e.Use(middleware.CORS("*"))
	e.Use(middleware.RateLimiter(limiterStore))

	e.GET("/no", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	if err := e.Start(":" + port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
