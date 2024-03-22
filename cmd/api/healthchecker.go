package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "enviroment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}