package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GeraTemplate)

	http.ListenAndServe(":8080", mux)
}
func GeraTemplate(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmp.Execute(w, []Curso{
		{"Go", 40},
		{"Java", 45},
		{"Pyton", 50},
	})
	if err != nil {
		panic(err)
	}
}
