package main

import (
	"hexagonal-example/internal/di"
	"log"
)

func main() {
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello from port 8080!")
		})
		http.ListenAndServe(":8080", nil)
	*/

	srv, err := di.InitializeApp()
	if err != nil {
		log.Fatal("Failed to initialize application:", err)
	}
	log.Println("Listening on 8080...")
	log.Fatal(srv.ListenAndServe())

}
