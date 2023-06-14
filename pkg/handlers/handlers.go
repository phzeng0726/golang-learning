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

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
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
