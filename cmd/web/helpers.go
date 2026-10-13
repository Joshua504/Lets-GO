package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// the function writes a log entry at error level (including the request method and url as attribute), then sends a generic 500 internal server error to the user
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "url", url, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// this function sends specific status code and corresponding description to users
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData){
	ts, ok := app.templateCache[page]
	if !ok{
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}
	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil{
		app.serverError(w, r, err)
	}
}