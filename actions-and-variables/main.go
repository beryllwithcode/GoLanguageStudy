package main

import "fmt"
import "net/http"
import "html/template"

type Info struct {
	Affiliation string
	Address    	string
}

type Person struct {
	Name 	string
	Gender 	string
	Hobbies []string
	Info 	Info
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name: "Beryll Pradana Ramadhan",
			Gender: "Male",
			Hobbies: []string{"Learning", "Coding", "SkuyLiving"},
			Info: Info{"Indonesia", "Jakarta"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}