package main

import (
	"fmt"
	"net/http"
	"os"
)

func serveForm(w http.ResponseWriter, r *http.Request) {
	//Get method checking
	if r.Method != http.MethodGet {
		http.Error(w, "Mehtod not alloweed", http.StatusMethodNotAllowed)
		return
	}

	//Reading index.html
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "could not load index.html", http.StatusInternalServerError)
		return
	}

	//Set reponse type and send html content
	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	//Post method checking
	if r.Method != http.MethodPost {
		//Redirecting any get request to the form
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//Getting name from form
	name := r.FormValue("username")

	//Respone
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello, %v </h1>", name)
}

func main() {

	//Route for form page
	http.HandleFunc("/", serveForm)

	//Route for form submission
	http.HandleFunc("/submit", handleSubmit)

	//Starting the port
	fmt.Println("Starting port at-- https://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error in starting prot: ", err)
		return
	}

}
