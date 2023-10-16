package main

import (
	"fmt",
	"net/http",
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err Error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServer)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(http.status), status)
}
