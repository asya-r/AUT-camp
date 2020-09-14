package main

import (
	"AUT-camp/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if os.Getenv("PORT") != "" {
		log.Print(os.Getenv("PORT"))
		port = os.Getenv("PORT")
	}
	log.Print("8080")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/camps", http.StatusSeeOther) })
	http.HandleFunc("/camps", controllers.Camps)
	http.ListenAndServe(":"+port, nil)
}
