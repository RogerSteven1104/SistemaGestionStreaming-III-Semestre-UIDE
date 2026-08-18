package main

import (
	"fmt"
	"net/http"
)

func iniciarServidor() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		fmt.Fprintln(w, "Sistema de Gestión de Streaming")
		fmt.Fprintln(w, "Servidor web funcionando correctamente.")
	})

	fmt.Println("=== SISTEMA DE GESTIÓN DE STREAMING ===")
	fmt.Println("Servidor iniciado en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
