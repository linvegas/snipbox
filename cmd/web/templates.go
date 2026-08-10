package main

import (
	"time"
	"html/template"
	"path/filepath"

	"snipbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func humanDate(t time.Time) string {
	// You need to use this exact date and hour
	return t.Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, p := range pages {
		fileName := filepath.Base(p)

		baseTpl, err := template.New(fileName).Funcs(functions).ParseFiles("./ui/html/base.html")

		if err != nil {
			return nil, err
		}

		partialsTpl, err := baseTpl.ParseGlob("./ui/html/partials/*.html")

		if err != nil {
			return nil, err
		}

		finalTpl, err := partialsTpl.ParseFiles(p)

		if err != nil {
			return nil, err
		}

		cache[fileName] = finalTpl
	}

	return cache, nil
}
