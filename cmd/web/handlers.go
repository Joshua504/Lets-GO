package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joshua504/Lets-GO/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "home.html", templateData{
		Snippets: snippets,
	})

	//files := []string{
	//	"./ui/html/base.html",
	//	"./ui/html/partials/nav.html",
	//	"./ui/html/pages/home.html",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, r, err) // using the serverError() helper
	//	return
	//}
	//
	//data := templateData{
	//	Snippets: snippets,
	//}
	//
	//err = ts.ExecuteTemplate(w, "base", data)
	//if err != nil {
	//	app.serverError(w, r, err) // using the serverError() helper
	//} // use the serverError() helper
	//
	//w.Write([]byte("Welcome to my world!"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, http.StatusOK, "view.html", templateData{
		Snippet: snippet,
	})

	//files := []string{
	//	"./ui/html/base.html",
	//	"./ui/html/partials/nav.html",
	//	"./ui/html/pages/view.html",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, r, err)
	//	return
	//}
	//
	//data := templateData{
	//	Snippet: snippet,
	//}
	//
	//err = ts.ExecuteTemplate(w, "base", data)
	//if err != nil {
	//	app.serverError(w, r, err)
	//}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet ....."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\\nClimb Mount Fuji,\\nBut slowly, slowly!\\n\\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
