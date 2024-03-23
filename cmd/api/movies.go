package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.c0utin.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		http.NotFound(w, r)

		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Frieren",
		Runtime:   6969,
		Genres:    []string{"anime", "drama", "fantasy"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)

	if err != nil {

		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and encountered a problem and could not process your request", http.StatusInternalServerError)

	}

}
