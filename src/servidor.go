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

// UsuarioRespuesta representa la información pública
// de un usuario que será enviada mediante JSON.
type UsuarioRespuesta struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Plan   string `json:"plan"`
}

// ReproduccionRespuesta representa el resultado de reproducir
// un contenido mediante el servicio web.
type ReproduccionRespuesta struct {
	Tipo    string `json:"tipo"`
	Titulo  string `json:"titulo"`
	Mensaje string `json:"mensaje"`
}

// servicioContenidos devuelve los contenidos registrados
// en el sistema utilizando formato JSON.
func servicioContenidos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

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

// servicioUsuarios devuelve la información pública
// de los usuarios registrados en formato JSON.
func servicioUsuarios(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	usuario := NuevoUsuario(
		1,
		"Carlos",
		"roger@email.com",
		"123456",
		"Premium",
	)

	respuesta := UsuarioRespuesta{
		ID:     usuario.GetID(),
		Nombre: usuario.GetNombre(),
		Correo: usuario.GetCorreo(),
		Plan:   usuario.GetPlan(),
	}

	err := json.NewEncoder(w).Encode([]UsuarioRespuesta{respuesta})

	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
}

// servicioReproduccion demuestra la reproducción de diferentes
// tipos de contenido utilizando las estructuras Pelicula y Serie.
func servicioReproduccion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pelicula := NuevaPelicula(
		"Avatar 2",
		180,
	)

	serie := NuevaSerie(
		"Stranger Things",
		4,
	)

	respuestas := []ReproduccionRespuesta{
		{
			Tipo:    "pelicula",
			Titulo:  pelicula.GetTitulo(),
			Mensaje: "Reproduciendo película: " + pelicula.GetTitulo(),
		},
		{
			Tipo:    "serie",
			Titulo:  serie.GetTitulo(),
			Mensaje: "Reproduciendo serie: " + serie.GetTitulo(),
		},
	}

	err := json.NewEncoder(w).Encode(respuestas)

	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
}

// iniciarServidor configura y pone en funcionamiento
// el servidor web del Sistema de Gestión de Streaming.
func iniciarServidor() {

	// Ruta principal del servidor.
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

		err := json.NewEncoder(w).Encode(categoriasJSON)

		if err != nil {
			http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
			return
		}
	})

	// Servicio Web #3: consulta de información de usuarios.
	http.HandleFunc("/api/usuarios", servicioUsuarios)

	// Servicio Web #4: reproducción de películas y series.
	http.HandleFunc("/api/reproduccion", servicioReproduccion)

	fmt.Println("=== SISTEMA DE GESTIÓN DE STREAMING ===")
	fmt.Println("Servidor iniciado en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
