package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Envs and inital setup
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// Data reading
	b, err := os.ReadFile("reasons.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
	}

	var reasons = []string{}

	if err := json.Unmarshal(b, &reasons); err != nil {
		log.Fatalf("Couldn't unmarshal file: %v\n", err)
	}

	// Configure and start server
	e := echo.New()

	limiterStore := middleware.NewRateLimiterMemoryStore(20)

	e.Use(middleware.Recover())
	e.Use(middleware.CORS("*"))
	e.Use(middleware.RateLimiter(limiterStore))

	e.GET("/no", func(c *echo.Context) error {
		reason := reasons[rand.IntN(len(reasons)-1)]

		return c.JSON(http.StatusOK, map[string]string{"reason": reason})
	})

	if err := e.Start(":" + port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
