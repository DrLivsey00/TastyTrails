package handlers

import (
	"fmt"
	"github.com/DrLivsey00/TastyTrails/internal/config"
	"github.com/DrLivsey00/TastyTrails/internal/models"
	"github.com/DrLivsey00/TastyTrails/internal/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
	fmt.Println("Home Page opened")
}
func (m *Repository) Add(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"
	dishName := r.FormValue("dish_name")
	render.RenderTemplate(w, "add_dish.page.tmpl", &models.TemplateData{StringMap: stringMap})
	fmt.Println("Add dish opened")
	fmt.Println("блюдо: " + dishName)
}
