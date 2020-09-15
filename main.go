package main

import (
	"AUT-camp/controllers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/camps", http.StatusSeeOther) })
	http.HandleFunc("/camps", controllers.Camps)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/signup", controllers.Signup)
	http.ListenAndServe(":8080", nil)
}
