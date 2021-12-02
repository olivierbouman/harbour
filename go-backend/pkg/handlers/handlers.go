package handlers

import (
	"go-backend/pkg/config"
	"go-backend/pkg/models"
	"go-backend/pkg/render"
	"net/http"
)

// The repository used by our handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
// The "(m *Repository)" is a receiver
// The functions with a receiver have access to everything they receive
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// Juliette is the about page handler
func (m *Repository) Juliette(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "juliette.page.tmpl", &models.TemplateData{})
}

// Boats is the about page handler
func (m *Repository) Boats(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "boats.page.tmpl", &models.TemplateData{})
}

// Boats_overview is the about page handler
func (m *Repository) BoatsOverview(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "boats_overview.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Info per boat"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
