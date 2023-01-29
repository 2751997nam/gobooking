package handlers

import (
	"net/http"

	"gobooking/pkg/config"
	"gobooking/pkg/models"
	"gobooking/pkg/render"
)

// repository pattern
var Repo *Repository

type Repository struct {
	AppConfig *config.AppConfig
}

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: ac,
	}
}

func NewHandlers(repo *Repository) {
	Repo = repo
}

func (m *Repository) DoNothing(res http.ResponseWriter, req *http.Request) {
}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	remoteIp := req.RemoteAddr
	m.AppConfig.Session.Put(req.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(res, "home.html", &models.TemplateData{})
}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	remoteIp := m.AppConfig.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(res, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
