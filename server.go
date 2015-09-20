package main

import (
  "net/http"
	"html/template"

  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"
)

func homepage(c web.C, w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/homepage.html")
  t.Execute(w, map[string] string {"title": "Vojtěch Bartoš, software engineer"})
}

func main() {
  goji.Get("/", homepage)
  goji.Serve()
}
