package handlers

import (
	"learning/pkg/config"
	"learning/pkg/models"
	"learning/pkg/render"
	"net/http"
)

// the repository used by the handlers
var Repo *Repository

// is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewHandlers(a *config.AppConfig) {
	Repo = &Repository{
		App: a,
	}
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// // create a new repository
// func NewRepo(a *config.AppConfig) *Repository {
// 	return &Repository{
// 		App: a,
// 	}
// }

// // sets the repository for the handlers
// func NewHandlers(r *Repository) {
// 	Repo = r
// }
