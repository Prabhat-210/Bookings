package render

import (
	"bytes"
	"github.com/Prabhat-210/Bookings/pkg/config"
	"github.com/Prabhat-210/Bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td //I created this so that we can use same data in multiple web page
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		//get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td) //pass data here which is in td

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} //empty map
	//get all of the files named *.page.tmpl from ./template
	pages, err := filepath.Glob("/home/prabhat.rawat@npci.org.in/Documents/API_Learnings/GO_by_Trevor Sawler/Day2/templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page) //When I parse the file named page and store that in a template called name right here.
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("/home/prabhat.rawat@npci.org.in/Documents/API_Learnings/GO_by_Trevor Sawler/Day2/templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("/home/prabhat.rawat@npci.org.in/Documents/API_Learnings/GO_by_Trevor Sawler/Day2/templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
