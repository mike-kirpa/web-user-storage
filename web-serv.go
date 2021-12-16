package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"time"
)

type Message struct {
	Name   string
	Email  string
	Errors map[string]string
}

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

	tmpl, err := template.ParseFiles("index.html", "error.html")
	if err != nil {
		fmt.Print(w, err.Error())
	}

	if r.Method == http.MethodPost {
		msg := &Message{"", "", make(map[string]string)}
		if r.FormValue("name") == "" {
			msg.Errors["Name"] = "Please input name."
		}

		if r.FormValue("email") == "" {
			msg.Errors["Email"] = "Please input email."
		}

		if m, _ := regexp.MatchString(`^([\w\.\_]{2,254})@(\w{1,}).([a-z]{2,4})$`, r.FormValue("email")); !m {
			msg.Errors["Email"] = "Please input valid email."
		}

		if m, _ := regexp.MatchString(`[a-zA-Z .]`, r.FormValue("name")); !m {
			msg.Errors["Name"] = "Please input valid name(English letters, dots, spaces)."
		}

		for k := range u {
			if r.FormValue("email") == u[k].Email {
				msg.Errors["Email"] = "This email is already added."
			}
		}

		if len(msg.Errors) != 0 {
			//tmpl, _ = template.ParseFiles("error.html")
			tmpl.ExecuteTemplate(w, "error", msg)
		} else {
			u.add(r.FormValue("name"), r.FormValue("email"))
			tmpl.ExecuteTemplate(w, "index", u)
		}
	} else {
		tmpl.ExecuteTemplate(w, "index", u)
	}
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

/*
func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Println(fmt.Sprintf("%q", x))
		rec := httptest.NewRecorder()
		fn(rec, r)
		log.Println(fmt.Sprintf("%q", rec.Body))
	}
}
*/
