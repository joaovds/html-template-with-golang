package main

import (
	"html/template"
	"os"
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
      Link: "github.com/gandalf",
      Languages: []string{"Go", "Elixir", "TypeScript"},
    },
    {
      Name: "Frodo",
      Email: "frodo@email.com",
      Link: "github.com/frodo",
      Languages: []string{"Crystal", "Rust", "Lua"},
    },
    {
      Name: "Sam",
      Email: "sam@email.com",
      Link: "github.com/sam",
      Languages: []string{"Go", "Python", "JavaScript"},
    },
  }

  templates := []string{
    "./public/head.html",
    "./public/content.html",
    "./public/index.html",
  }

  t := template.New("index.html")
  t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
  t = template.Must(t.ParseFiles(templates...))
  err := t.Execute(os.Stdout, profiles)
  if err != nil {
    panic(err)
  }
}
