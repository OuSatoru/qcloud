package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	mainpage = "www/tmpl/home.html"
	four04 = "www/tmpl/404.html"
)

func main() {
	//mux := http.NewServeMux()

	http.HandleFunc("/", mainPage)
	//http.HandleFunc("/a", mainPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static"))))
	http.ListenAndServe(":8088", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, _ := template.ParseFiles(mainpage)
		t.Execute(w, nil)
	} else {
		fourOFour(w)
	}

}

func printUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func fourOFour(w http.ResponseWriter)  {
	t, _ := template.ParseFiles(four04)
	t.Execute(w, nil)
}
