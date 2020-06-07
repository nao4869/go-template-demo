package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("temp.gohtml"))
}

func main() {
	error := tpl.ExecuteTemplate(os.Stdout, "temp.gohtml", "continuation")
	if error != nil {
		log.Fatalln(error)
	}
}
