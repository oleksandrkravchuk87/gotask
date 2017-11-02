package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)





//--------------gorilla------------------
type Page struct {
	Title string
	Body  []byte
}


func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{name}", index).Methods("GET")
	router.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to request from gorilla")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<div>Hello: %s from Gorilla</div>", name)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to request")
	log.Println(r.UserAgent())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<div>%s</div>", "Hello world")


}

//------------------------------------------------------------------------------------------------
//----------------net/http------------------------------------------------------------------------
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}
//
//
//func viewHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path[1:] != "" {
//		fmt.Fprintf(w, "<div>%s</div>",  "Hello world " + r.URL.Path[1:] +"!")
//	} else {
//		fmt.Fprintf(w, "<div>%s</div>", "Hello world")
//	}
//
//}
//
//func main() {
//	http.HandleFunc("/", viewHandler)
//	http.ListenAndServe(":8080", nil)
//}
//------------------------------------------------------------------------------------------------
//func (p *Page) save() error {
//	filename := p.Title + ".txt"
//	return ioutil.WriteFile(filename, p.Body, 0600)
//}
//
//func loadPage(title string) (*Page, error) {
//	filename := title + ".txt"
//	body, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return nil, err
//	}
//	return &Page{Title: title, Body: body}, nil
//}