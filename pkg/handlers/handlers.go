package handlers

import (
	"learning/pkg/config"
	"learning/pkg/render"
	"net/http"
)

// the repository used by the handlers
var Repo *Repository

// is the repository type
type Repository struct {
	App *config.AppConfig
}

// create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// func NewHandlers(a *config.AppConfig) {
// 	Repo = &Repository{
// 		App: a,
// 	}
// }

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")

}
