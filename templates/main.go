package main

import (
	"log"
	"os"
	"text/template"
)

// Excute.template will excute the file which specified

func main() {
	template, error := template.ParseFiles("template.gohtml")
	if error != nil {
		log.Fatalln(error)
	}

	// create a new html file
	newFile, error := os.Create("index.html")
	if error != nil {
		log.Println("error createing file", error)
	}
	defer newFile.Close()

	// print out html content to stdout
	error = template.Execute(os.Stdout, nil)
	if error != nil {
		log.Fatalln(error)
	}

	// white into new file
	error = template.Execute(newFile, nil)
	if error != nil {
		log.Fatalln(error)
	}
}
