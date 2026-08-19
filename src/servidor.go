package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// PlanRespuesta representa la información de un plan
// de suscripción disponible en el sistema.
type PlanRespuesta struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

// ActualizacionContenido representa los datos que pueden
// modificarse mediante el servicio PUT.
type ActualizacionContenido struct {
	Titulo   string `json:"titulo"`
	Duracion int    `json:"duracion"`
	Genero   string `json:"genero"`
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

// servicioCategorias devuelve las categorías registradas
// en el sistema utilizando formato JSON.
func servicioCategorias(w http.ResponseWriter, r *http.Request) {

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

// servicioPlanes devuelve los planes de suscripción
// disponibles en el sistema utilizando formato JSON.
func servicioPlanes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	planes := []PlanRespuesta{
		{
			ID:     1,
			Nombre: "Básico",
			Precio: 7.99,
		},
		{
			ID:     2,
			Nombre: "Estándar",
			Precio: 10.99,
		},
		{
			ID:     3,
			Nombre: "Premium",
			Precio: 14.99,
		},
	}

	err := json.NewEncoder(w).Encode(planes)

	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
}

// servicioBuscarContenido permite buscar contenidos
// utilizando el género recibido como parámetro.
func servicioBuscarContenido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	generoBuscado := strings.TrimSpace(r.URL.Query().Get("genero"))

	if generoBuscado == "" {
		http.Error(
			w,
			"Debe proporcionar el parámetro genero",
			http.StatusBadRequest,
		)
		return
	}

	contenidos := []ContenidoRespuesta{
		{
			ID:            1,
			Titulo:        "Avatar 2",
			Descripcion:   "Una historia de ciencia ficción en Pandora.",
			Genero:        "Ciencia ficción",
			Duracion:      180,
			Clasificacion: "PG-13",
		},
		{
			ID:            2,
			Titulo:        "John Wick",
			Descripcion:   "Un antiguo asesino regresa a la acción.",
			Genero:        "Acción",
			Duracion:      101,
			Clasificacion: "R",
		},
		{
			ID:            3,
			Titulo:        "Son como niños",
			Descripcion:   "Un grupo de amigos se reúne nuevamente.",
			Genero:        "Comedia",
			Duracion:      102,
			Clasificacion: "PG-13",
		},
	}

	resultados := []ContenidoRespuesta{}

	for _, contenido := range contenidos {
		if strings.EqualFold(contenido.Genero, generoBuscado) {
			resultados = append(resultados, contenido)
		}
	}

	if len(resultados) == 0 {
		http.Error(
			w,
			"No se encontraron contenidos para el género indicado",
			http.StatusNotFound,
		)
		return
	}

	err := json.NewEncoder(w).Encode(resultados)

	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
}

// servicioActualizarContenido permite modificar un contenido
// mediante una solicitud HTTP PUT.
func servicioActualizarContenido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido. Utilice PUT.", http.StatusMethodNotAllowed)
		return
	}

	partes := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(partes) != 3 {
		http.Error(w, "Identificador de contenido inválido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(partes[2])

	if err != nil {
		http.Error(w, "El ID debe ser numérico", http.StatusBadRequest)
		return
	}

	if id != 1 {
		http.Error(w, "Contenido no encontrado", http.StatusNotFound)
		return
	}

	var datos ActualizacionContenido

	err = json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	datos.Titulo = strings.TrimSpace(datos.Titulo)
	datos.Genero = strings.TrimSpace(datos.Genero)

	if datos.Titulo == "" {
		http.Error(w, "El título no puede estar vacío", http.StatusBadRequest)
		return
	}

	if datos.Genero == "" {
		http.Error(w, "El género no puede estar vacío", http.StatusBadRequest)
		return
	}

	if datos.Duracion <= 0 {
		http.Error(w, "La duración debe ser mayor que cero", http.StatusBadRequest)
		return
	}

	contenido := NuevoContenido(
		id,
		datos.Titulo,
		"Una historia de ciencia ficción en Pandora.",
		datos.Genero,
		datos.Duracion,
		"PG-13",
	)

	respuesta := map[string]interface{}{
		"mensaje": "Contenido actualizado correctamente",
		"contenido": ContenidoRespuesta{
			ID:            contenido.idContenido,
			Titulo:        contenido.GetTitulo(),
			Descripcion:   contenido.descripcion,
			Genero:        contenido.GetGenero(),
			Duracion:      contenido.GetDuracion(),
			Clasificacion: contenido.clasificacion,
		},
	}

	err = json.NewEncoder(w).Encode(respuesta)

	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
}

// servicioEliminarContenido permite eliminar un contenido
// mediante una solicitud HTTP DELETE.
func servicioEliminarContenido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido. Utilice DELETE.", http.StatusMethodNotAllowed)
		return
	}

	partes := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(partes) != 3 {
		http.Error(w, "Identificador de contenido inválido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(partes[2])

	if err != nil {
		http.Error(w, "El ID debe ser numérico", http.StatusBadRequest)
		return
	}

	// En esta versión académica se considera registrado
	// el contenido con ID 1.
	if id != 1 {
		http.Error(w, "Contenido no encontrado", http.StatusNotFound)
		return
	}

	respuesta := map[string]interface{}{
		"mensaje": "Contenido eliminado correctamente",
		"id":      id,
	}

	err = json.NewEncoder(w).Encode(respuesta)

	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
}

// servicioContenidoPorID permite dirigir las solicitudes
// PUT y DELETE al servicio correspondiente.
func servicioContenidoPorID(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPut:
		servicioActualizarContenido(w, r)

	case http.MethodDelete:
		servicioEliminarContenido(w, r)

	default:
		http.Error(
			w,
			"Método no permitido. Utilice PUT o DELETE.",
			http.StatusMethodNotAllowed,
		)
	}
}

// iniciarServidor configura y pone en funcionamiento
// el servidor web del Sistema de Gestión de Streaming.
func iniciarServidor() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		fmt.Fprintln(w, "Sistema de Gestión de Streaming")
		fmt.Fprintln(w, "Servidor web funcionando correctamente.")
	})

	// Servicio Web #1: consulta de información de contenidos.
	http.HandleFunc("/api/contenidos", servicioContenidos)

	// Servicio Web #2: consulta de categorías registradas.
	http.HandleFunc("/api/categorias", servicioCategorias)

	// Servicio Web #3: consulta de información de usuarios.
	http.HandleFunc("/api/usuarios", servicioUsuarios)

	// Servicio Web #4: reproducción de películas y series.
	http.HandleFunc("/api/reproduccion", servicioReproduccion)

	// Servicio Web #5: consulta de planes de suscripción.
	http.HandleFunc("/api/planes", servicioPlanes)

	// Servicio Web #6: búsqueda de contenidos por género.
	http.HandleFunc("/api/contenidos/buscar", servicioBuscarContenido)

	// Servicios Web #7 y #8:
	// actualización mediante PUT y eliminación mediante DELETE.
	http.HandleFunc("/api/contenidos/", servicioContenidoPorID)

	fmt.Println("=== SISTEMA DE GESTIÓN DE STREAMING ===")
	fmt.Println("Servidor iniciado en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
