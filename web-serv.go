package main

import (
	"html/template"
	"net/http"
	"time"
)

type User struct {
	Index int
	Name  string
	Email string
	Date  time.Time
	Dates string
}
type Users []User

var u Users

func home_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		u.add(r.FormValue("name"), r.FormValue("email"))
	}

	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, u)
}

func (u *Users) add(name string, email string) {

	t := time.Now()
	i := len(*u) + 1
	*u = append(*u, User{Index: i, Name: name, Email: email, Date: t, Dates: t.Format("02.01.2006")})
}

func main() {

	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
}
