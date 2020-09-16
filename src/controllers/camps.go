package controllers

import (
	"html/template"
	"log"
	"net/http"
)

// Camps is home page handler
func Camps(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	tpl, err := template.ParseFiles("src/views/camps.gohtml")
	check(err)
	text := "Main page"
	err = tpl.Execute(w, map[string]interface{}{
		"text": text,
	})
	check(err)
}
