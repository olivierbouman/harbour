package render

import (
	"bytes"
	"fmt"
	"go-backend/pkg/config"
	"go-backend/pkg/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// FuncMap is a map of functions
var functions = template.FuncMap{}

var app *config.AppConfig

// Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Adds default data to your td that should be available on every page
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	// If you are in say development mode you dont want to use your cache
	if app.UseCache {
		// Get the templateCache from the appConfig
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Use the template name tmpl as key
	// ok is a variable that states whether the key exists or not
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from templateCache")
	}

	// This is a buffer that will hold information
	buf := new(bytes.Buffer)

	// Add default data
	td = AddDefaultData(td)

	// Execute your template and store its value in the buffer
	// Plus add our templateData td to the buffer
	_ = t.Execute(buf, td)

	// Now we have enough to write to our responseWriter
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// We dont want to reload this cache every time a new template is being called
	// or loaded in. To do this we need to set up an application wide
	// configurations settings.
	// The common approach is to create global variables; which isn't the way to go
	// Our configurations are stored in pkg/config/config.go

	// Map is a 'dict'
	myCache := map[string]*template.Template{}

	// Find all .page templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Initialize your ts
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Find the .layout templates in the same folder
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// If you find those, parse them
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
