package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const weatherPath = "/weather"
	var routes = []string{weatherPath}
	port := ":8080"

	http.HandleFunc(weatherPath, weatherHandler)

	fmt.Printf("Server starting (%s)\n", port)
	fmt.Println("Routes:")
	for i := range routes {
		fmt.Printf("\t %s \n", routes[i])
	}

	fmt.Println("Ready")
	log.Fatal(http.ListenAndServe(port, nil))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))
}
