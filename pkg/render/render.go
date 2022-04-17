package render

import (
	"bytes"
	"fmt"
	"github.com/okiprakasa/hello-world/models"
	"github.com/okiprakasa/hello-world/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig
var tcError error

// TemplatePointer transfer reference of TemplateCache and the error associated to render package
func TemplatePointer(a *config.AppConfig, err error) {
	app = a
	tcError = err
}

func AddTemplateData(td *models.TemplateData) *models.TemplateData {

	return td
}

//Template renders template from bytes buff of template cache
func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//Get the template cache from the AppConfig (Production Mode)
		tc = app.TemplateCache
	} else {
		//Create the template cache everytime func Template is loaded (Development Mode)
		tc, tcError = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Cannot get template from template cache")
	}

	templates := template.Must(t, tcError)

	td = AddTemplateData(td)

	buf := new(bytes.Buffer)
	_ = templates.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

//CreateTemplateCache output map of template cache memory
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	files := make([]string, 2)
	files[0], _ = filepath.Abs("./templates/base.layout.tmpl")

	for _, page := range pages {
		name := filepath.Base(page)
		files[1], _ = filepath.Abs(page)

		ts, err := template.New("base.layout.tmpl").Funcs(functions).ParseFiles(files...)
		if err != nil {
			return myCache, err
		}

		myCache[name] = ts
	}

	return myCache, nil
}

//T func create template from the base up
func _(w http.ResponseWriter, data interface{}, tmpl ...string) {
	//https://stackoverflow.com/questions/35906683/go-render-html-template-with-inheritance
	//tmpl example [base.layout home.page]
	//cwd, _ := os.Getwd()
	files := make([]string, len(tmpl))
	fmt.Println(tmpl)
	for i, file := range tmpl {
		files[i], _ = filepath.Abs("./templates/" + file + ".tmpl")
		//files[i] = filepath.Join(cwd, "./templates/"+file+".tmpl")
	}

	//https://stackoverflow.com/questions/49043292/error-template-is-an-incomplete-or-empty-template
	t, err := template.New("base.layout.tmpl").Funcs(functions).ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templates := template.Must(t, err)
	err = templates.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
