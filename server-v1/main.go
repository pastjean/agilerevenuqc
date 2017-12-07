package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	version      = "v1"
	htmlTemplate = `<html>
	<head><title>Agile Revenu Québec</title>
	<link href="https://fonts.googleapis.com/css?family=Palanquin+Dark" rel="stylesheet">
	</head>
	<body style="font-family: 'Palanquin Dark', sans-serif;font-size:2em;display:flex;align-items: center;flex-direction:column;justify-content: center;background-color:#efefef"><h1>Agile Revenu Québec</h1><div style="display:flex;align-items: center;justify-content: center;background-color:white;min-height:50%;min-width:50%;max-width:300px;padding:2em 2em">{{.Quote}}</div></body>
</html>`
)

func getText() string {
	return "Agile Revenu Québec 🎉"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("version", version)

		if r.Header.Get("accepts") == strings.ToLower("application/json") {
			w.Header().Set("content-type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(map[string]string{"text": getText()})
			return
		}

		w.Header().Set("content-type", "text/html; charset=utf-8")
		tmpl, _ := template.New("t").Parse(htmlTemplate)
		tmpl.Execute(w, map[string]string{"Quote": getText()})
	})

	port := "0.0.0.0:8080"
	log.Println("Listening on: ", port)
	log.Panic(http.ListenAndServe(port, nil))
}

var ()
