package transport

import (
	"log"
	"net/http"
)

func Run(port string) {
	http.HandleFunc("/status", status)
	http.HandleFunc("/", index)
	http.HandleFunc("/new_member", newMember)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("[Fatal] ListenAndServe", err)
	}
}
