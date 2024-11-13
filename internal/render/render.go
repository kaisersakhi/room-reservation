package render

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, templateName string) {

}

func BuildTemplateCache() (map[string]*template.Template, error) {

}
