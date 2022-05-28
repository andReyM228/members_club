package transport

import (
	"log"
	"members_club/service"
	"net/http"
	"text/template"
)

func status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Println("[ERROR] Error from status")
	}

}

func index(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	tmpl.Execute(w, service.GetListOfMembers())
}

func newMember(w http.ResponseWriter, r *http.Request) {
	err := service.AddMember(r.PostFormValue("name"), r.PostFormValue("email"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
