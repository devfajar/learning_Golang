package main

import (
	"net/http"
	"fmt"
	"html/template"
)

type Info struct {
	Affiliation string
	Address		string
}

type Person struct {
	Name		string
	Gender 		string
	Hobbies		[]string
	Info		Info
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:		"Bruce Wayne",
			Gender:		"Male",
			Hobbies:	[]string{"Reading Books", "Travelling", "Buying Things"},
			Info:		Info{"Wayne Enterprises", "Gotham City"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 Divisions"
}