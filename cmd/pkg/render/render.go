package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderPage render pages to templates
func RenderTemplate(w http.ResponseWriter, html string) {
	// get the template cache from the app config
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[html]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
//	fmt.Println("Execution is done")
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writng template to browser", err)
	}
//	fmt.Println("Buffer is written")
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
