package handlers

import (
	"github.com/okiprakasa/hello-world/pkg/config"
	"github.com/okiprakasa/hello-world/pkg/render"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSFRToken string //CSFRToken Cross Site Request Forgery Token (Security token for forms)
}

var Repo *Repository

// Repository create repository type consist of all AppConfig
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates Repository containing all AppConfig
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers create New Handler from reference of repository memory to activate other handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	//files := []string{"base.layout", "home.page"}
	//render.T(w, nil, files...)
	render.Template(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, _ *http.Request) {
	//files := []string{"base.layout", "about.page"}
	//render.T(w, nil, files...)
	render.Template(w, "about.page.tmpl")
}
