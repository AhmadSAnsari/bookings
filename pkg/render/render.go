package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AhmadSAnsari/bookings/pkg/config"
	"github.com/AhmadSAnsari/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for template page
func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td	
}

// RenderPage render pages to templates
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[html]
	//fmt.Println(t, ok)
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	buf := new(bytes.Buffer)

	td = addDefaultData(td)

	_ = t.Execute(buf, td)

	fmt.Println("Execution is done")
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writng template to browser", err)
	}
	fmt.Println("Buffer is written")
}

// CreateTemplateCache create template cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {
		myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("Page is currently", page) // Just to chek

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		// functions define aboue as var
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
