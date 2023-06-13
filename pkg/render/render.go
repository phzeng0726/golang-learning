package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// var functions = template.FuncMap{}

// var app *config.AppConfig

//	func NewTemplates(a *config.AppConfig) {
//		app = a
//	}
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Panicln(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all of the files name *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl") // get full path
	if err != nil {
		return myCache, err
	}
	// range through all fils ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page) // only split last file name
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
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
