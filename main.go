package main

import (
	"fmt"
	"net/http"
	"time"
)

// Manejador para la ruta /saludo
func saludoHandler(w http.ResponseWriter, r *http.Request) {
	// Responder con un mensaje en formato JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"mensaje": "¡Hola, bienvenido al microservicio en Go!"}`)
}

// Manejador para la ruta /despedirse
func adiosHandler(w http.ResponseWriter, r *http.Request) {
	//responder con otro mensaje JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"mensaje2": "Adios!"}`)
}

// Manejador para la ruta /fecha
func fechaHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener la fecha y hora actual
	fechaActual := time.Now().Format("2006-01-02 15:04:05")

	// Responder con la fecha y hora en formato JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"fecha": "%s"}`, fechaActual)
}

func main() {
	// Asignar la función al endpoint /saludo
	http.HandleFunc("/saludo", saludoHandler)

	// Asignar la otra función al endpoint /adios
	http.HandleFunc("/adios", adiosHandler)

	// Asignar la función al endpoint /fecha
	http.HandleFunc("/fecha", fechaHandler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor ejecutándose en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
