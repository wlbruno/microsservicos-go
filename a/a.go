package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9090",nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, "")
}

func process(w http.ResponseWriter, r *http.Request) {
	Log.Println(r.FormValue("coupon"))
	Log.Println(r.FormValue("cc-number"))

	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, "")
}