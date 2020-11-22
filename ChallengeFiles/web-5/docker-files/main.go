package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gorilla/mux"
)

func jsonError(w http.ResponseWriter, err interface{}, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	pdf := new(pdfDoc)
	err := json.NewDecoder(r.Body).Decode(pdf)

	if err != nil {
		log.Println("Failed to decode json with error: ", err)
		w.WriteHeader(400)
		fmt.Fprint(w, `{"result": "Error, why you do this"}`)
		return
	}

	if !pdf.sanitize() {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"result": "Failed to validate fields, why you do this"}`)
		return
	}

	fillTemplate(*pdf)
	exampleNewPDFGenerator()

	fmt.Fprint(w, `{"result": "success"}`)
}

func flagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, `<html><p>Here's a Flag for ya WACTF{59d7dd1cfd283ead53a8c9e28719903a}</p></html>`)
}

type pdfDoc struct {
	CompanyName string `json:"companyName"`
	Service     string `json:"service"`
	Description string `json:"description"`
	PriceCents  int32  `json:"priceCents"`
	Quantity    int32  `json:"quantity"`
	Total       int32  `json:"total"`
	Tax         int32  `json:"tax"`
	GrandTotal  int32  `json:"grandTotal"`
}

func (r *pdfDoc) sanitize() bool {
	if len(r.CompanyName) > 30 || len(r.Description) > 60 || len(r.Service) > 42 {
		return false
	}

	r.CompanyName = html.EscapeString(r.CompanyName)
	r.Description = html.EscapeString(r.Description)

	return true
}

func main() {
	flagMux := http.NewServeMux()
	flagMux.HandleFunc("/", flagHandler)

	go func() {
		http.ListenAndServe("127.0.0.1:9090", flagMux)
	}()

	staticMux := mux.NewRouter()
	staticMux.HandleFunc("/api", apiHandler)

	// Create the route
	staticMux.PathPrefix("/").Handler(http.FileServer(http.Dir("/var/www/html/")))

	go func() {
		srv := &http.Server{
			Handler:      staticMux,
			Addr:         ":9000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Println(srv.ListenAndServe())
	}()

	serverMux := mux.NewRouter()
	serverMux.HandleFunc("/api", apiHandler)
	serverMux.PathPrefix("/").Handler(http.FileServer(http.Dir("/var/www/serve/")))

	srv := &http.Server{
		Handler:      serverMux,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println(srv.ListenAndServe())
}

func fillTemplate(pdf pdfDoc) {
	tmpl, err := template.New("pdf.html.tpl").ParseFiles("/var/www/template/pdf.html.tpl")
	if err != nil {
		log.Println(err)
	}

	f, err := os.Create("/var/www/html/index.html")

	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, pdf)
	// err = tmpl.Execute(w, document)
	if err != nil {
		log.Println(err)
	}
	w.Flush()

}

func exampleNewPDFGenerator() {

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println(err)
	}

	pdfg.Dpi.Set(100)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage("http://127.0.0.1:9000/index.html")
	page.JavascriptDelay.Set(6000)
	page.NoStopSlowScripts.Set(true)
	// page.DebugJavascript.Set(true)
	// pdfg.SetStderr(os.Stdout)
	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.EnableLocalFileAccess.Set(false)
	page.LoadMediaErrorHandling.Set("skip")
	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Println(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("/var/www/serve/output/invoice.pdf")
	if err != nil {
		log.Println(err)
	}

	log.Println("PDF Created Successfully")
}
