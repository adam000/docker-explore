package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func main() {
	tpl = template.New("T")
	template.Must(tpl.ParseGlob("templates/*"))

	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/run", runHandler).
		Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	ignore := map[string]interface{}{}
	err := tpl.ExecuteTemplate(w, "explore.html", ignore)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	// NOTE: this is generally something you should never do because
	// of the potential vulnerabilities. Imagine if someone ran `rm rf /`
	// against your machine. This is for demonstration purposes only, not
	// suitable for production use.
	command := r.FormValue("command")
	log.Print(command)
	args := strings.Split(command, " ")
	name := args[0]
	args = args[1:]

	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	fmt.Fprint(w, string(output))
}
