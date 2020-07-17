package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"
	"rtsa/aplus"
	"rtsa/goldenlotus"
	"rtsa/savenda"
)



func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}









// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

// ServeHTTP wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}


func main() {

	f, err := os.OpenFile("" +
		"activity.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/rtsa", aplus.RstaPost).Methods("Post")
	router.HandleFunc("/rtsa-savenda", savenda.RstaPostSavenda).Methods("Post")
	router.HandleFunc("/rtsa-golden-lotus", goldenlotus.RtsaPostGoldenLotus).Methods("Post")

	port := os.Getenv("PORT")

	log.Println(port)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8090", handler))
}


