package controllers

import (
	"AUT-camp/models"
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
	camps := models.GetCamps()

	tpl, err := template.ParseFiles("src/views/camps.gohtml")
	check(err)
	err = tpl.Execute(w, map[string]interface{}{
		"camps": camps,
	})
	check(err)
}
