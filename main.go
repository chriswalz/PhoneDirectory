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
	fmt.Println("hello")


	router := mux.NewRouter()

	/*basiccssPath := "/css/basic.css"
	router.HandleFunc(basiccssPath, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static"+basiccssPath)
	})
	jsTableSorter := "/js/tablesorter.js"
	router.HandleFunc(jsTableSorter, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static"+jsTableSorter)
	}) */
	//s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	//router.PathPrefix("/static/").Handler(s)
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		database.Prepare()
		users := database.GetRows()
		indexTmp, err := template.ParseFiles("static/index.html") // todo move to outside of func
		if err != nil {
			fmt.Println(err)
		}
		tmp := template.Must(indexTmp, err)
		tmp.Execute(w, struct {
			Phones []*database.PhoneUser
		}{
			users,
		})
	}).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8080", router)
}
