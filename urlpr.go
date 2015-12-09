package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

type Page struct {
	Title string
}

type PageError struct {
	Title     string
	ErrorCode int
	ErrorDesc string
}

var (
	templates *template.Template
	port      = flag.Int("port", 8383, "The port the pm_search app should listen on")
)

func main() {
	// setup template
	tpl_path, err := FSString(false, "/static/templates.html")
	if err != nil {
		log.Printf("Error creating templates: %s\n", err.Error())
	}
	templates = template.Must(template.New("").Parse(tpl_path))

	http.HandleFunc("/print", handlePrint)
	http.HandleFunc("/printers", handlePrinters)
	http.HandleFunc("/", handleIndex)
	http.Handle("/static/", http.FileServer(FS(false)))

	// serve single root level files
	//handleSingle("/favicon.ico", "/static/img/favicon.ico")

	log.Printf("urlpr server started - localhost:%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, r, 404, "The page you are looking for does not exist.")
		return
	}

	page := &Page{
		Title: "URL Printer",
	}

	// render the page...
	if err := templates.ExecuteTemplate(w, "index", page); err != nil {
		log.Printf("Error executing template: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func handlePrint(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		Title: "Printing",
	}

	url := r.FormValue("url")
	opts := r.FormValue("opts")
	res, err := http.Get(url)
	if err != nil {
		handleError(w, r, 500, err.Error())
		return
	}
	defer res.Body.Close()
	cmd := exec.Command("lpr", opts)
	cmd.Stdin = res.Body
	err = cmd.Run()
	if err != nil {
		handleError(w, r, 500, err.Error())
		return
	}

	// render the page...
	if err := templates.ExecuteTemplate(w, "print", page); err != nil {
		log.Printf("Error executing template: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func handlePrinters(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("lpstat", "-p", "-d")
	cmd.Stdout = w
	cmd.Run()
	return
}

func handleSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		b, err := FSByte(false, filename)
		if err != nil {
			log.Printf("Error serving single file: %s\n %s\n", filename, err.Error())
		}
		w.Write(b)
	})
}

func handleError(w http.ResponseWriter, r *http.Request, status int, desc string) {
	title := http.StatusText(status)
	if title == "" {
		title = "Unknown Error"
	}
	page := &PageError{
		Title:     title,
		ErrorCode: status,
		ErrorDesc: desc,
	}
	w.WriteHeader(page.ErrorCode)
	if err := templates.ExecuteTemplate(w, "error", page); err != nil {
		http.Error(w, page.ErrorDesc, page.ErrorCode)
	}
}
