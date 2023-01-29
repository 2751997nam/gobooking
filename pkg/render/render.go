package render

import (
	"bytes"
	"gobooking/pkg/config"
	"gobooking/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var appConfig *config.AppConfig

// NewTemplates sets  the config for the template package
func NewTemplates(ac *config.AppConfig) {
	appConfig = ac

}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(res http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template
	if appConfig.UseCache {
		//get template cache from app config
		templateCache = appConfig.TemplateCache

	} else {
		templateCache, _ = CreateTemplateCache()
	}

	//get requested template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
		return
	}
	buf := new(bytes.Buffer)
	templateData = AddDefaultData(templateData)
	err := template.Execute(buf, templateData)
	if err != nil {
		log.Println(err)
		return
	}
	//render template

	_, err = buf.WriteTo(res)
	if err != nil {
		log.Println(err)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.html from ./templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	//range through all files endinng with .html
	for _, page := range pages {
		fileName := filepath.Base(page)
		templateSet, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
		}
		myCache[fileName] = templateSet
	}

	return myCache, err
}
