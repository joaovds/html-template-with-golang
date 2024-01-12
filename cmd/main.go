package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type GithubProfile struct {
  Name string
  Email string
  Link string
  Languages []string
}

func main() {
  profiles := []GithubProfile{
    {
      Name: "Gandalf",
      Email: "gandalf@email.com",
      Link: "https://github.com/gandalf",
      Languages: []string{"Go", "Elixir", "TypeScript"},
    },
    {
      Name: "Frodo",
      Email: "frodo@email.com",
      Link: "https://github.com/frodo",
      Languages: []string{"Crystal", "Rust", "Lua"},
    },
    {
      Name: "Sam",
      Email: "sam@email.com",
      Link: "https://github.com/sam",
      Languages: []string{"Go", "Python", "JavaScript"},
    },
  }

  templates := []string{
    "./public/head.html",
    "./public/content.html",
    "./public/index.html",
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    t := template.New("index.html")
    t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
    t = template.Must(t.ParseFiles(templates...))
    err := t.Execute(w, profiles)
    if err != nil {
      panic(err)
    }
  })

  fmt.Println("Server running on port 8080")
  http.ListenAndServe(":8080", nil)
}
