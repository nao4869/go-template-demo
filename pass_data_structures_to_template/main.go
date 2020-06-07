package main

// . is current piece of data
// range is entry/element of slice

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

// struct demo
type publicCloud struct {
	Name        string
	Description string
}

func init() {
	tpl = template.Must(template.ParseFiles("temp.gohtml"))
}

func main() {
	// pass slices to template
	//slices := []string{"AWS", "GCP", "Azure", "Alibaba"}

	// struct demo
	publicCloud := publicCloud{
		Name:        "AWS",
		Description: "Amazon Web Services",
	}

	error := tpl.ExecuteTemplate(os.Stdout, "temp.gohtml", publicCloud)
	if error != nil {
		log.Fatalln(error)
	}
}
