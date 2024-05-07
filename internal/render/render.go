package render

import (
	"bytes"
	"github.com/DrLivsey00/TastyTrails/internal/config"
	"github.com/DrLivsey00/TastyTrails/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// sets a new template cache
func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn`t get tamplate cache:", tmpl)
	}
	buf := new(bytes.Buffer)
	td = AddDefData(td)
	_ = t.Execute(buf, td)
	//render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}
		cache[name] = ts
	}
	return cache, nil
}
