package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ContenidoRespuesta representa la información que se enviará
// al cliente mediante el servicio web en formato JSON.
type ContenidoRespuesta struct {
	ID            int    `json:"id"`
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	Genero        string `json:"genero"`
	Duracion      int    `json:"duracion"`
	Clasificacion string `json:"clasificacion"`
}

// servicioContenidos devuelve los contenidos registrados
// en el sistema utilizando formato JSON.
func servicioContenidos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	contenido := NuevoContenido(
		1,
		"Avatar 2",
		"Una historia de ciencia ficción en Pandora.",
		"Ciencia ficción",
		180,
		"PG-13",
	)

	respuesta := ContenidoRespuesta{
		ID:            1,
		Titulo:        contenido.GetTitulo(),
		Descripcion:   "Una historia de ciencia ficción en Pandora.",
		Genero:        contenido.GetGenero(),
		Duracion:      contenido.GetDuracion(),
		Clasificacion: "PG-13",
	}

	err := json.NewEncoder(w).Encode(respuesta)

	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
}

func iniciarServidor() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		fmt.Fprintln(w, "Sistema de Gestión de Streaming")
		fmt.Fprintln(w, "Servidor web funcionando correctamente.")
	})

	// Servicio Web #1: consulta de información de contenidos.
	http.HandleFunc("/api/contenidos", servicioContenidos)

	// Servicio Web #2: consulta de categorías registradas.
	http.HandleFunc("/api/categorias", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		categoriasJSON := []map[string]interface{}{}

		for _, categoria := range categorias {
			categoriasJSON = append(categoriasJSON, map[string]interface{}{
				"id":     categoria.idCategoria,
				"nombre": categoria.GetNombre(),
			})
		}

		json.NewEncoder(w).Encode(categoriasJSON)
	})

	fmt.Println("=== SISTEMA DE GESTIÓN DE STREAMING ===")
	fmt.Println("Servidor iniciado en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}

}
