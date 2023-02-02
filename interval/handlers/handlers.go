package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gobooking/interval/config"
	"gobooking/interval/forms"
	"gobooking/interval/models"
	"gobooking/interval/render"
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
	render.RenderTemplate(res, req, "home.html", &models.TemplateData{})
}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	remoteIp := m.AppConfig.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(res, req, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) General(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "generals.html", &models.TemplateData{})
}

func (m *Repository) Major(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "majors.html", &models.TemplateData{})
}

func (m *Repository) MakeReservation(res http.ResponseWriter, req *http.Request) {
	var emtpyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emtpyReservation
	render.RenderTemplate(res, req, "make-reservation.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostMakeReservation(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: req.Form.Get("first_name"),
		LastName:  req.Form.Get("last_name"),
		Email:     req.Form.Get("email"),
		Phone:     req.Form.Get("phone"),
	}

	form := forms.New(req.PostForm)
	form.Required(req, "first_name", "last_name", "email", "phone")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(res, req, "make-reservation.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}

func (m *Repository) Contact(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "contact.html", &models.TemplateData{})
}

func (m *Repository) SearchAvailbilty(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "reservation.html", &models.TemplateData{})
}

func (m *Repository) PostSearchAvailbilty(res http.ResponseWriter, req *http.Request) {
	start := req.Form.Get("start")
	end := req.Form.Get("end")

	res.Header().Set("Content-Type", "application/json")

	result := struct {
		Start string
		End   string
	}{
		Start: start,
		End:   end,
	}
	json.NewEncoder(res).Encode(result)
}
