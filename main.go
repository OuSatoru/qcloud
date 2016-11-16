package main

import (
	"github.com/OuSatoru/qcloud/grab"
	"fmt"
	"html/template"
	"net/http"
	"math/rand"
	"time"
	"strconv"
)

const (
	mainpage = "www/tmpl/home.html"
	four04 = "www/tmpl/404.html"
	image = "www/tmpl/image.html"
)

func main() {
	//mux := http.NewServeMux()
	//fmt.Println(grab.Grab("https://yande.re/post/show/374899"))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/random", randomImg)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static"))))
	http.ListenAndServe(":8087", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, _ := template.ParseFiles(mainpage)
		t.Execute(w, nil)
	} else {
		fourOFour(w)
	}
}

func fourOFour(w http.ResponseWriter)  {
	t, _ := template.ParseFiles(four04)
	t.Execute(w, nil)
}

func randomImg(w http.ResponseWriter, r *http.Request)  {
	//fmt.Println(r.URL.Path)
	//lk := struct {
	//	Link string
	//}{}
	//TODO: largest num of post in yande temporarily 374903
	//bigger than 4

	ur := grab.YandeHead + strconv.Itoa(random(374904))
	link := grab.Grab(ur)
	fmt.Println(link)
	http.Redirect(w, r, link, http.StatusMovedPermanently)

	//lk.Link = grab.Grab(ur)
	//t, _ := template.ParseFiles(image)
	//t.Execute(w, lk)
	////t.Execute(os.Stdout, lk)
	//http.Redirect(w, r, )
}

func random(largest int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	temp := rand.Intn(largest)
	if temp < 5 {
		temp = random(largest)
	}
	return temp
}