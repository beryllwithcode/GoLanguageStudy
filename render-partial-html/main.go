package main

import (
    "net/http"
    "html/template"
    "fmt"
)

type M map[string]interface{}

func main() {
    var _, err = template.ParseGlob("views/*")
    if err != nil {
        panic(err.Error())
        
        return
    }

    http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
        var data = M{"name": "Beryll"}
        var tmpl = template.Must(template.ParseFiles(
            "views/index.html",
            "views/_header.html",
            "views/_message.html",
        ))

        err = tmpl.ExecuteTemplate(w, "index", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        var data = M{"fullName": "Beryll Pradana Ramadhan"}
        var tmpl = template.Must(template.ParseFiles(
            "views/about.html",
            "views/_header.html",
            "views/_message.html",
        ))

        err = tmpl.ExecuteTemplate(w, "about", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    fmt.Println("Server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}             