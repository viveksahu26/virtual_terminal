package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/joho/godotenv"
)

// It process commands provided by the user and return back to the user.
func commands(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	// extracting values from form
	command := r.FormValue("command")
	log.Println("Command entered by the user is: ", command)

	cmds := exec.Command("bash", "-c", command)
	Stdout, err := cmds.Output()
	if err != nil {
		w.Write([]byte("Command is either incorrect or command is interactive in nature. \nPlease make sure command is correct and non-interactive in nature."))
		log.Println(err)
		return
	}

	data := struct {
		Output string
	}{
		Output: string(Stdout),
	}

	tpl.ExecuteTemplate(w, "extra.html", data)
}

// It is to check server is up and working.
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

// It is the user interface to enter it's commands.
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

// Init() initializes all * files from template folders
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// Load() --> reads the .env file and loads the set variables into the environment
	err := godotenv.Load()
	if err != nil {
		log.Println("Error in loading .env files..")
	}

	log.Println("Virtual Terminal Service Starts ...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	log.Println("PORT no. assigned to the app is: ", port)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// /home endpoint is mapped to home
	http.HandleFunc("/home", home)

	// /health endpoint is mapped to healthCheckUp
	http.HandleFunc("/health", healthCheckUp)

	// /cmd endpoint is mapped to cmd
	// It is internally invoked by form.
	http.HandleFunc("/cmd", commands)

	// Server Listening on localhost:<port>
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
