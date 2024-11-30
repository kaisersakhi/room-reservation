package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"room-reservation/internal/config"
)

var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func Render(w http.ResponseWriter, templateName string) {
	var err error
	var tmpl *template.Template
	var ok bool

	if app.IsTemplateCacheEnabled() {
		tmpl, ok = app.GetTemplateCache()[templateName]
	} else {
		tempCache, err := BuildTemplateCache()

		if err == nil {
			ok = true
		}

		tmpl = tempCache[templateName]
	}

	if !ok {
		log.Fatal("could not find template cache.")
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, "")

	if err != nil {
		log.Println("failed to excute data on templates", err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

func BuildTemplateCache() (map[string]*template.Template, error) {
	log.Println("Building template cache...")

	cache := map[string]*template.Template{}

	htmlFiles, err := filepath.Glob("web/templates/*.html")

	if err != nil {
		return nil, err
	}

	for _, file := range htmlFiles {
		name := filepath.Base(file)

		templateSet, err := template.New(name).ParseFiles(file)

		if err != nil {
			break
		}

		layoutFiles, err := filepath.Glob("web/templates/*.layout.html")

		if err != nil {
			break
		}

		if len(layoutFiles) >= 1 {
			templateSet, err = templateSet.ParseGlob("web/templates/*.layout.html")

			if err != nil {
				break
			}
		}

		cache[name] = templateSet
	}

	return cache, nil
}
