package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "./index.html")

	}))

	port := ":8080"
	fmt.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
