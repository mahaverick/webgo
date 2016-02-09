package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type Coupon struct {
	Name      string
	Completed bool
	StartDate time.Time
	EndDate   time.Time
}

type Coupons []Coupon

func couponApiHandler(w http.ResponseWriter, r *http.Request) {
	coupons := Coupons{
		Coupon{Name: "O2 Spa"},
		Coupon{Name: "Gold Gym"},
	}

	json.NewEncoder(w).Encode(coupons)
}

func renderTemplate(w http.ResponseWriter, tmpl string, def string) {
	t := template.Must(template.New("tele").ParseFiles("views/" + tmpl + ".html"))
	if err := t.ExecuteTemplate(w, def, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", "homepage")
}

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", homeHandler)
	s := router.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/coupons", couponApiHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", router)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
