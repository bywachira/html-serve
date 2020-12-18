package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bywachira/html-server/helpers"
)

// Serve handles serving the directory
func Serve() {
	port := helpers.GeneratePortNumber()
	log.Println("We are running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir("."))))
}
