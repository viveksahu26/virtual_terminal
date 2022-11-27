package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

func commands(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	command := r.FormValue("command")
	fmt.Println("command by user is: ", command)

	// cmd := exec.Command(command[0])
	cmds := exec.Command("bash", "-c", command)

	Stdout, err := cmds.Output()
	if err != nil {
		w.Write([]byte("Command is Incorrect. Please enter correct command."))
		fmt.Println(err.Error())
		return
	}
	d := struct {
		Output string
	}{
		Output: string(Stdout),
	}

	tpl.ExecuteTemplate(w, "extra.html", d)
}

func healthCheckUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method other than GET not Supported...", http.StatusNotFound)
		return
	}

	w.Write([]byte("<h1>Health of Server is UP & Running... !!</h1>"))
}

func home(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	fmt.Println("GET")
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var tpl *template.Template

// Initialize it with all *.html files from template folders..
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("Virtual Terminal Service Starts ...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	fmt.Println("PORT is: ", port)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/home", home)

	// /health endpoint is mapped to healthCheckUp
	http.HandleFunc("/health", healthCheckUp)

	// /cmd endpoint is mapped to cmd
	// http.HandleFunc("/home", home)

	// /cmd endpoint is mapped to cmd
	http.HandleFunc("/cmd", commands)

	// Server Listening on localhost:9009
	err := http.ListenAndServe(":"+port, nil) // setting listening port
	if err != nil {
		panic(err)
	}
}
