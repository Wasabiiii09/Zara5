package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// .env laden (falls vorhanden)
	_ = godotenv.Load()

	// Handler für die Startseite
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("public/index.html")
		if err != nil {
			http.Error(w, "HTML Datei nicht gefunden", http.StatusInternalServerError)
			return
		}

		data := map[string]string{
			"Title":   "Paparazie",
			"Message": "Willkommen auf meiner Go-Homepage!",
		}

		tmpl.Execute(w, data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server läuft auf http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}