package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/OuSatoru/qcloud/grab"
	"github.com/OuSatoru/qcloud/runner"
	"github.com/OuSatoru/qcloud/wechat"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	mainpage = "www/tmpl/home.html"
	four04   = "www/tmpl/404.html"
	image    = "www/tmpl/image.html"
)

func main() {
	dbUser := flag.String("du", "postgres", "Database User")
	dbPwd := flag.String("dp", "", "Database Password")
	appId := flag.String("ai", "", "AppId")
	appSecret := flag.String("as", "", "AppSecret")
	flag.Parse()
	if *dbPwd == "" {
		fmt.Println("Going to have no access to db nor access-token get.")
	} else {
		//dbLogin := runner.DbLogin{DbUser: *dbUser, DbPwd: *dbPwd}
		dbConnect, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/wechat?sslmode=disable", *dbUser, *dbPwd))
		if err != nil {
			log.Println(err)
		}
		if *appId != "" && *appSecret != "" {
			accessToken := wechat.AccessToken{AppId: *appId, AppSecret: *appSecret}
			go runner.RunningGetAccToken(dbConnect, accessToken)
		} else {
			fmt.Println("Going to have no access-token get.")
		}
		//TODO: largest num of post in yande, go func
	}

	//mux := http.NewServeMux()
	//fmt.Println(grab.Grab("https://yande.re/post/show/374899"))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/random", randomImg)

	http.HandleFunc("/wx", wechat.Handle)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static"))))
	http.ListenAndServe(":80", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, _ := template.ParseFiles(mainpage)
		t.Execute(w, nil)
		log.Println(r.RemoteAddr, "entering main page.")
	} else {
		fourOFour(w)
	}
}

func fourOFour(w http.ResponseWriter) {
	t, _ := template.ParseFiles(four04)
	t.Execute(w, nil)
}

func randomImg(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Path)
	//lk := struct {
	//	Link string
	//}{}
	// temporarily 376275
	//bigger than 4

	ur := grab.YandeHead + strconv.Itoa(random(376275))
	//link := grab.Grab(ur)
	log.Println(ur)
	http.Redirect(w, r, ur, http.StatusMovedPermanently)

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
