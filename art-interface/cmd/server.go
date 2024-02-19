package main

import (
	art_interf "art/art-interface/pkg/interface"
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Output string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("HTTP/1.1 200 OK")
		http.ServeFile(w, r, "index.html")

	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		// Get the value of the "input" field from the form
		userInput := r.Form.Get("input")
		isMalformed := false
		var output string

		// Decode the input
		output, isMalformed = art_interf.DecodeInput(userInput)

		if isMalformed {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		} else {
			fmt.Println("HTTP/1.1 202 Accepted")

			// Create a PageData struct with the decoded value
			data := PageData{
				Output: output,
			}

			// Render the HTML template with the PageData
			tmpl := template.Must(template.ParseFiles("index.html"))
			tmpl.Execute(w, data)

		}

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	// Register the indexHandler function to handle requests to the root URL ("/")
	http.HandleFunc("/", indexHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}