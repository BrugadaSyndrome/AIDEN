package main

import (
	"html/template"
	"log"
	"os"
)

type cssSettings struct {
	TextLightColor          string
	BodyBackgroundColor     string
	GameCardBackgroundColor string
	NavBackgroundColor      string
}

func generateStaticFiles() {
	css := cssSettings{
		TextLightColor:          "#FFFFFF",
		BodyBackgroundColor:     "#565656",
		GameCardBackgroundColor: "#33C0FE",
		NavBackgroundColor:      "#3E5EFC",
	}

	templates, err := template.New("").ParseFiles("web/css/input.css")
	if err != nil {
		log.Fatal("unable to parse index.css: " + err.Error())
	}

	f, err := os.Create("web/css/index.css")
	if err != nil {
		log.Fatal("unable to open index.css: " + err.Error())
	}

	err = templates.ExecuteTemplate(f, "index.css", css)
	if err != nil {
		log.Fatal("web/css parse failed: " + err.Error())
	}
}
