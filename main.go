package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/chriswalz/Tablerich/database"
)

type Phone struct {
	Name   string
	Number string
}

func main() {

	//fmt.Println(users[0].First)
	//database.CreateAndOpen("chris")
	RunServer()
}
func RunServer() {
	fmt.Println("Running")
	database.Prepare()
	users := database.GetRows()
	tmp := template.Must(template.ParseFiles("static/index.html"))

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmp.Execute(w, struct {
			Phones []*database.PhoneUser
		}{
			users,
		})
	}).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8080", router)
}
