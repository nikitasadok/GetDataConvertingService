package main

import (
	"GETDataConvertService/routes"
	"flag"
	"log"
	"net/http"
)

var endpoint = flag.String("endpoint", "http://localhost:29999", "Enter valid endpoint of "+
	"DB service")

func main() {
	flag.Parse()
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":30001", router))

}
