package main

import (
	"fmt"
	"github.com/autobaza/auto_catalog/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/car-makes", handlers.HandleCarMakes)
	http.HandleFunc("/car-models", handlers.HandleCarModels)
	http.HandleFunc("/car-serie", handlers.HandleCarSerie)

	fmt.Printf("Starting server at port 2222\n")
	if err := http.ListenAndServe(":2222", nil); err != nil {
		log.Fatal(err)
	}
}
