package main

import (
	"math/rand"
	"net/http"
	"time"
	 "os"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	Title  string
	Year string
}

var movies []Movie = []Movie{
	{"KGF", "2018"},
	{"Bahubali", "2015"},
	{"Avengers Endgame", "2019"},
	{"Premaloka", "1987"},
	{"Pirates of the Carribean", "2003"},
}

func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()

	api.GET("/movies", getMovie)
	api.POST("/movies", getMovie)
	api.PUT("/movies", getMovie)

	port := os.Getenv("PORT")

	if port == "" {
		port = "25255"
	}


	api.Start(":"+port)

}

func getMovie(c echo.Context) error {
	index := rand.Intn(len(movies))
	return c.JSON(http.StatusOK, movies[index])
}