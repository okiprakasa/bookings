package handlers

import (
	"github.com/okiprakasa/hello-world/models"
	"github.com/okiprakasa/hello-world/pkg/config"
	"github.com/okiprakasa/hello-world/pkg/render"
	"net/http"
)

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
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//files := []string{"base.layout", "home.page"}
	//render.T(w, nil, files...)

	m.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr) //Save the user IP address to session
	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//files := []string{"base.layout", "about.page"}
	//render.T(w, nil, files...)

	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip") //Access the saved user IP from session when accessing home page
	stringMap["remote_ip"] = remoteIP

	// Send the data to the template
	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
