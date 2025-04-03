package main

import (
	"fmt"
	"net/http"
)

// Manejador para la ruta /saludo
func saludoHandler(w http.ResponseWriter, r *http.Request) {
	// Responder con un mensaje en formato JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"mensaje": "¡Hola, bienvenido al microservicio en Go!"}`)
	fmt.Fprintf(w, `{"mensaje2": "Así queda mejor??"}`)
}

func main() {
	// Asignar la función al endpoint /saludo
	http.HandleFunc("/saludo", saludoHandler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor ejecutándose en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
