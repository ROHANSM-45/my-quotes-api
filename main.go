package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	AUthor string
}

//var quotes [] quote=make([]quote, 0)

var quotes []quote = []quote{
	{"kaam karo helloooooooo ", "p"},
	{"kaam karo helloooooooo  ", "pawaffee"},
}

func main() {
	api := echo.New()

	api.GET("/quotes", getQuotes)
	api.GET("/quotes/random", getRandomQuote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}
	api.Start(":" + port)
}

func getQuotes(c echo.Context) error {
	return c.JSON(http.StatusOK, quotes)
}

func getRandomQuote(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])

}
