package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	tpl, err := template.ParseFiles("./views/login.gohtml")
	check(err)
	err = tpl.Execute(w, map[string]interface{}{})
	check(err)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	tpl, err := template.ParseFiles("./views/signup.gohtml")
	check(err)
	err = tpl.Execute(w, map[string]interface{}{})
	check(err)
}
