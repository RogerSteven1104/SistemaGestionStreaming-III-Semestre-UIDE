package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"SistemaGestionStreaming/database"
)

// ContenidoRespuesta representa la información pública de un contenido
// que será enviada mediante JSON.
type ContenidoRespuesta struct {
	ID            int    `json:"id"`
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	Genero        string `json:"genero"`
	Duracion      int    `json:"duracion"`
	Clasificacion string `json:"clasificacion"`
}

// UsuarioRespuesta representa la información pública de un usuario.
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

// PlanRespuesta representa un plan de suscripción.
type PlanRespuesta struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

// NuevoContenidoRequest representa los datos recibidos
// para crear un nuevo contenido mediante POST.
type NuevoContenidoRequest struct {
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	Genero        string `json:"genero"`
	Duracion      int    `json:"duracion"`
	Clasificacion string `json:"clasificacion"`
}

// ActualizacionContenido representa los datos que pueden
// modificarse mediante PUT.
type ActualizacionContenido struct {
	Titulo   string `json:"titulo"`
	Duracion int    `json:"duracion"`
	Genero   string `json:"genero"`
}

// ============================================================
// CATÁLOGO CENTRAL DE CONTENIDOS
// ============================================================

// contenidos almacenará los contenidos registrados
// actualmente en el sistema.
var contenidos = []ContenidoRespuesta{
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

// ============================================================
// CATÁLOGO CENTRAL DE USUARIOS
// ============================================================

// usuarios almacena los usuarios registrados
// actualmente en el sistema.
var usuarios = []*Usuario{
	NuevoUsuario(
		1,
		"Carlos",
		"roger@email.com",
		"",
		"Premium",
	),
}

// ============================================================
// SERVICIO WEB 1
// GET /api/contenidos
// ============================================================

// servicioContenidos devuelve el catálogo de contenidos
// disponibles en el sistema.
func servicioContenidos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido. Utilice GET.",
			http.StatusMethodNotAllowed,
		)
		return
	}

	filas, err := database.DB.Query(`
		SELECT id, titulo, descripcion, genero, duracion, clasificacion
		FROM contenidos
		ORDER BY id
	`)

	if err != nil {
		http.Error(
			w,
			"Error al consultar los contenidos",
			http.StatusInternalServerError,
		)
		return
	}

	defer filas.Close()

	resultados := []ContenidoRespuesta{}

	for filas.Next() {

		var contenido ContenidoRespuesta

		err := filas.Scan(
			&contenido.ID,
			&contenido.Titulo,
			&contenido.Descripcion,
			&contenido.Genero,
			&contenido.Duracion,
			&contenido.Clasificacion,
		)

		if err != nil {
			http.Error(
				w,
				"Error al leer los contenidos",
				http.StatusInternalServerError,
			)
			return
		}

		resultados = append(resultados, contenido)
	}

	if err := filas.Err(); err != nil {
		http.Error(
			w,
			"Error al recorrer los contenidos",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(resultados)
}

// ============================================================
// SERVICIO WEB 2
// POST /api/contenidos
// ============================================================

// servicioCrearContenido permite registrar un nuevo contenido.
func servicioCrearContenido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Método no permitido. Utilice POST.",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var datos NuevoContenidoRequest

	err := json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	datos.Titulo = strings.TrimSpace(datos.Titulo)
	datos.Descripcion = strings.TrimSpace(datos.Descripcion)
	datos.Genero = strings.TrimSpace(datos.Genero)
	datos.Clasificacion = strings.TrimSpace(datos.Clasificacion)

	if datos.Titulo == "" {
		http.Error(
			w,
			"El título no puede estar vacío",
			http.StatusBadRequest,
		)
		return
	}

	if datos.Descripcion == "" {
		http.Error(
			w,
			"La descripción no puede estar vacía",
			http.StatusBadRequest,
		)
		return
	}

	if datos.Genero == "" {
		http.Error(
			w,
			"El género no puede estar vacío",
			http.StatusBadRequest,
		)
		return
	}

	if datos.Duracion <= 0 {
		http.Error(
			w,
			"La duración debe ser mayor que cero",
			http.StatusBadRequest,
		)
		return
	}

	if datos.Clasificacion == "" {
		http.Error(
			w,
			"La clasificación no puede estar vacía",
			http.StatusBadRequest,
		)
		return
	}

	// Guardar el contenido directamente en SQLite.
	resultado, err := database.DB.Exec(`
		INSERT INTO contenidos
		(titulo, descripcion, genero, duracion, clasificacion)
		VALUES (?, ?, ?, ?, ?)
	`,
		datos.Titulo,
		datos.Descripcion,
		datos.Genero,
		datos.Duracion,
		datos.Clasificacion,
	)

	if err != nil {
		http.Error(
			w,
			"Error al guardar el contenido en la base de datos",
			http.StatusInternalServerError,
		)
		return
	}

	// Obtener el ID generado por SQLite.
	nuevoID, err := resultado.LastInsertId()

	if err != nil {
		http.Error(
			w,
			"Error al obtener el ID del contenido",
			http.StatusInternalServerError,
		)
		return
	}

	nuevoContenido := ContenidoRespuesta{
		ID:            int(nuevoID),
		Titulo:        datos.Titulo,
		Descripcion:   datos.Descripcion,
		Genero:        datos.Genero,
		Duracion:      datos.Duracion,
		Clasificacion: datos.Clasificacion,
	}

	respuesta := map[string]interface{}{
		"contenido": nuevoContenido,
		"mensaje":   "Contenido creado correctamente",
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(respuesta)
}

// ============================================================
// SERVICIO WEB 3
// GET /api/categorias
// ============================================================

// servicioCategorias devuelve las categorías registradas.
func servicioCategorias(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido. Utilice GET.",
			http.StatusMethodNotAllowed,
		)
		return
	}

	categoriasJSON := []map[string]interface{}{}

	for _, categoria := range categorias {

		categoriasJSON = append(
			categoriasJSON,
			map[string]interface{}{
				"id":     categoria.idCategoria,
				"nombre": categoria.GetNombre(),
			},
		)
	}

	json.NewEncoder(w).Encode(categoriasJSON)
}

// ============================================================
// SERVICIO WEB 4
// GET /api/usuarios
// POST /api/usuarios
// PUT /api/usuarios/{id}
// DELETE /api/usuarios/{id}
// ============================================================

// servicioUsuarios gestiona las operaciones sobre los usuarios.
// Los datos se almacenan y consultan directamente desde SQLite.
func servicioUsuarios(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {

	case http.MethodGet:

		filas, err := database.DB.Query(`
			SELECT id, nombre, correo, plan
			FROM usuarios
			ORDER BY id
		`)

		if err != nil {
			http.Error(
				w,
				"Error al consultar los usuarios",
				http.StatusInternalServerError,
			)
			return
		}

		defer filas.Close()

		respuesta := []UsuarioRespuesta{}

		for filas.Next() {

			var usuario UsuarioRespuesta

			err := filas.Scan(
				&usuario.ID,
				&usuario.Nombre,
				&usuario.Correo,
				&usuario.Plan,
			)

			if err != nil {
				http.Error(
					w,
					"Error al leer los usuarios",
					http.StatusInternalServerError,
				)
				return
			}

			respuesta = append(respuesta, usuario)
		}

		if err := filas.Err(); err != nil {
			http.Error(
				w,
				"Error al recorrer los usuarios",
				http.StatusInternalServerError,
			)
			return
		}

		json.NewEncoder(w).Encode(respuesta)

	case http.MethodPost:

		var datos struct {
			Nombre string `json:"nombre"`
			Correo string `json:"correo"`
			Plan   string `json:"plan"`
		}

		err := json.NewDecoder(r.Body).Decode(&datos)

		if err != nil {
			http.Error(
				w,
				"JSON inválido",
				http.StatusBadRequest,
			)
			return
		}

		datos.Nombre = strings.TrimSpace(datos.Nombre)
		datos.Correo = strings.TrimSpace(datos.Correo)
		datos.Plan = strings.TrimSpace(datos.Plan)

		if datos.Nombre == "" {
			http.Error(
				w,
				"El nombre no puede estar vacío",
				http.StatusBadRequest,
			)
			return
		}

		if datos.Correo == "" {
			http.Error(
				w,
				"El correo no puede estar vacío",
				http.StatusBadRequest,
			)
			return
		}

		if datos.Plan == "" {
			http.Error(
				w,
				"El plan no puede estar vacío",
				http.StatusBadRequest,
			)
			return
		}

		resultado, err := database.DB.Exec(`
			INSERT INTO usuarios
			(nombre, correo, contrasena, plan)
			VALUES (?, ?, ?, ?)
		`,
			datos.Nombre,
			datos.Correo,
			"",
			datos.Plan,
		)

		if err != nil {

			if strings.Contains(
				strings.ToLower(err.Error()),
				"unique",
			) {
				http.Error(
					w,
					"El correo ya está registrado",
					http.StatusConflict,
				)
				return
			}

			http.Error(
				w,
				"Error al guardar el usuario en la base de datos",
				http.StatusInternalServerError,
			)
			return
		}

		nuevoID, err := resultado.LastInsertId()

		if err != nil {
			http.Error(
				w,
				"Error al obtener el ID del usuario",
				http.StatusInternalServerError,
			)
			return
		}

		nuevoUsuario := UsuarioRespuesta{
			ID:     int(nuevoID),
			Nombre: datos.Nombre,
			Correo: datos.Correo,
			Plan:   datos.Plan,
		}

		respuesta := map[string]interface{}{
			"mensaje": "Usuario creado correctamente",
			"usuario": nuevoUsuario,
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(respuesta)

	case http.MethodPut:

		partes := strings.Split(
			strings.Trim(r.URL.Path, "/"),
			"/",
		)

		if len(partes) != 3 {
			http.Error(
				w,
				"Identificador de usuario inválido",
				http.StatusBadRequest,
			)
			return
		}

		id, err := strconv.Atoi(partes[2])

		if err != nil {
			http.Error(
				w,
				"El ID debe ser numérico",
				http.StatusBadRequest,
			)
			return
		}

		var datos struct {
			Nombre string `json:"nombre"`
			Correo string `json:"correo"`
			Plan   string `json:"plan"`
		}

		err = json.NewDecoder(r.Body).Decode(&datos)

		if err != nil {
			http.Error(
				w,
				"JSON inválido",
				http.StatusBadRequest,
			)
			return
		}

		datos.Nombre = strings.TrimSpace(datos.Nombre)
		datos.Correo = strings.TrimSpace(datos.Correo)
		datos.Plan = strings.TrimSpace(datos.Plan)

		if datos.Nombre == "" {
			http.Error(
				w,
				"El nombre no puede estar vacío",
				http.StatusBadRequest,
			)
			return
		}

		if datos.Correo == "" {
			http.Error(
				w,
				"El correo no puede estar vacío",
				http.StatusBadRequest,
			)
			return
		}

		if datos.Plan == "" {
			http.Error(
				w,
				"El plan no puede estar vacío",
				http.StatusBadRequest,
			)
			return
		}

		resultado, err := database.DB.Exec(`
			UPDATE usuarios
			SET nombre = ?, correo = ?, plan = ?
			WHERE id = ?
		`,
			datos.Nombre,
			datos.Correo,
			datos.Plan,
			id,
		)

		if err != nil {

			if strings.Contains(
				strings.ToLower(err.Error()),
				"unique",
			) {
				http.Error(
					w,
					"El correo ya está registrado",
					http.StatusConflict,
				)
				return
			}

			http.Error(
				w,
				"Error al actualizar el usuario",
				http.StatusInternalServerError,
			)
			return
		}

		filasAfectadas, err := resultado.RowsAffected()

		if err != nil {
			http.Error(
				w,
				"Error al verificar la actualización",
				http.StatusInternalServerError,
			)
			return
		}

		if filasAfectadas == 0 {
			http.Error(
				w,
				"Usuario no encontrado",
				http.StatusNotFound,
			)
			return
		}

		usuarioActualizado := UsuarioRespuesta{
			ID:     id,
			Nombre: datos.Nombre,
			Correo: datos.Correo,
			Plan:   datos.Plan,
		}

		respuesta := map[string]interface{}{
			"mensaje": "Usuario actualizado correctamente",
			"usuario": usuarioActualizado,
		}

		json.NewEncoder(w).Encode(respuesta)

	case http.MethodDelete:

		partes := strings.Split(
			strings.Trim(r.URL.Path, "/"),
			"/",
		)

		if len(partes) != 3 {
			http.Error(
				w,
				"Identificador de usuario inválido",
				http.StatusBadRequest,
			)
			return
		}

		id, err := strconv.Atoi(partes[2])

		if err != nil {
			http.Error(
				w,
				"El ID debe ser numérico",
				http.StatusBadRequest,
			)
			return
		}

		resultado, err := database.DB.Exec(
			"DELETE FROM usuarios WHERE id = ?",
			id,
		)

		if err != nil {
			http.Error(
				w,
				"Error al eliminar el usuario",
				http.StatusInternalServerError,
			)
			return
		}

		filasAfectadas, err := resultado.RowsAffected()

		if err != nil {
			http.Error(
				w,
				"Error al verificar la eliminación",
				http.StatusInternalServerError,
			)
			return
		}

		if filasAfectadas == 0 {
			http.Error(
				w,
				"Usuario no encontrado",
				http.StatusNotFound,
			)
			return
		}

		respuesta := map[string]interface{}{
			"id":      id,
			"mensaje": "Usuario eliminado correctamente",
		}

		json.NewEncoder(w).Encode(respuesta)

	default:

		http.Error(
			w,
			"Método no permitido. Utilice GET, POST, PUT o DELETE.",
			http.StatusMethodNotAllowed,
		)
	}
}

// ============================================================
// SERVICIO WEB 5
// GET /api/reproduccion
// ============================================================

// servicioReproduccion demuestra el uso de interfaces,
// incrustación y polimorfismo mediante películas y series.
func servicioReproduccion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido. Utilice GET.",
			http.StatusMethodNotAllowed,
		)
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

	json.NewEncoder(w).Encode(respuestas)
}

// ============================================================
// SERVICIO WEB 6
// GET /api/planes
// ============================================================

// servicioPlanes devuelve los planes de suscripción disponibles.
func servicioPlanes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido. Utilice GET.",
			http.StatusMethodNotAllowed,
		)
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

	json.NewEncoder(w).Encode(planes)
}

// ============================================================
// SERVICIO WEB 7
// GET /api/contenidos/buscar?genero=...
// ============================================================

// servicioBuscarContenido permite buscar contenidos por género.
func servicioBuscarContenido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido. Utilice GET.",
			http.StatusMethodNotAllowed,
		)
		return
	}

	generoBuscado := strings.TrimSpace(
		r.URL.Query().Get("genero"),
	)

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

		if strings.EqualFold(
			contenido.Genero,
			generoBuscado,
		) {
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

	json.NewEncoder(w).Encode(resultados)
}

// ============================================================
// SERVICIO WEB 8
// PUT /api/contenidos/{id}
// ============================================================

// servicioActualizarContenido permite actualizar un contenido
// existente directamente en la base de datos SQLite.
func servicioActualizarContenido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodPut {
		http.Error(
			w,
			"Método no permitido. Utilice PUT.",
			http.StatusMethodNotAllowed,
		)
		return
	}

	partes := strings.Split(
		strings.Trim(r.URL.Path, "/"),
		"/",
	)

	if len(partes) != 3 {
		http.Error(
			w,
			"Identificador de contenido inválido",
			http.StatusBadRequest,
		)
		return
	}

	id, err := strconv.Atoi(partes[2])

	if err != nil {
		http.Error(
			w,
			"El ID debe ser numérico",
			http.StatusBadRequest,
		)
		return
	}

	var datos ActualizacionContenido

	err = json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(
			w,
			"JSON inválido",
			http.StatusBadRequest,
		)
		return
	}

	datos.Titulo = strings.TrimSpace(datos.Titulo)
	datos.Genero = strings.TrimSpace(datos.Genero)

	if datos.Titulo == "" {
		http.Error(
			w,
			"El título no puede estar vacío",
			http.StatusBadRequest,
		)
		return
	}

	if datos.Genero == "" {
		http.Error(
			w,
			"El género no puede estar vacío",
			http.StatusBadRequest,
		)
		return
	}

	if datos.Duracion <= 0 {
		http.Error(
			w,
			"La duración debe ser mayor que cero",
			http.StatusBadRequest,
		)
		return
	}

	// Actualizar directamente el registro en SQLite.
	resultado, err := database.DB.Exec(`
		UPDATE contenidos
		SET titulo = ?,
		    genero = ?,
		    duracion = ?
		WHERE id = ?
	`,
		datos.Titulo,
		datos.Genero,
		datos.Duracion,
		id,
	)

	if err != nil {
		http.Error(
			w,
			"Error al actualizar el contenido en la base de datos",
			http.StatusInternalServerError,
		)
		return
	}

	filasAfectadas, err := resultado.RowsAffected()

	if err != nil {
		http.Error(
			w,
			"Error al comprobar la actualización",
			http.StatusInternalServerError,
		)
		return
	}

	if filasAfectadas == 0 {
		http.Error(
			w,
			"Contenido no encontrado",
			http.StatusNotFound,
		)
		return
	}

	// Consultar nuevamente el contenido actualizado.
	var contenido ContenidoRespuesta

	err = database.DB.QueryRow(`
		SELECT id, titulo, descripcion, genero, duracion, clasificacion
		FROM contenidos
		WHERE id = ?
	`,
		id,
	).Scan(
		&contenido.ID,
		&contenido.Titulo,
		&contenido.Descripcion,
		&contenido.Genero,
		&contenido.Duracion,
		&contenido.Clasificacion,
	)

	if err != nil {
		http.Error(
			w,
			"Error al consultar el contenido actualizado",
			http.StatusInternalServerError,
		)
		return
	}

	respuesta := map[string]interface{}{
		"mensaje":   "Contenido actualizado correctamente",
		"contenido": contenido,
	}

	json.NewEncoder(w).Encode(respuesta)
}

// ============================================================
// SERVICIO WEB 9
// DELETE /api/contenidos/{id}
// ============================================================

// servicioEliminarContenido permite eliminar un contenido
// directamente desde la base de datos SQLite.
func servicioEliminarContenido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodDelete {
		http.Error(
			w,
			"Método no permitido. Utilice DELETE.",
			http.StatusMethodNotAllowed,
		)
		return
	}

	partes := strings.Split(
		strings.Trim(r.URL.Path, "/"),
		"/",
	)

	if len(partes) != 3 {
		http.Error(
			w,
			"Identificador de contenido inválido",
			http.StatusBadRequest,
		)
		return
	}

	id, err := strconv.Atoi(partes[2])

	if err != nil {
		http.Error(
			w,
			"El ID debe ser numérico",
			http.StatusBadRequest,
		)
		return
	}

	// Eliminar directamente el contenido de SQLite.
	resultado, err := database.DB.Exec(`
		DELETE FROM contenidos
		WHERE id = ?
	`,
		id,
	)

	if err != nil {
		http.Error(
			w,
			"Error al eliminar el contenido de la base de datos",
			http.StatusInternalServerError,
		)
		return
	}

	filasAfectadas, err := resultado.RowsAffected()

	if err != nil {
		http.Error(
			w,
			"Error al comprobar la eliminación",
			http.StatusInternalServerError,
		)
		return
	}

	if filasAfectadas == 0 {
		http.Error(
			w,
			"Contenido no encontrado",
			http.StatusNotFound,
		)
		return
	}

	respuesta := map[string]interface{}{
		"mensaje": "Contenido eliminado correctamente",
		"id":      id,
	}

	json.NewEncoder(w).Encode(respuesta)
}

// ============================================================
// CONTROLADOR DE CONTENIDOS POR ID
// ============================================================

// servicioContenidoPorID dirige las solicitudes PUT y DELETE
// hacia el servicio correspondiente.
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

// ============================================================
// SERVIDOR PRINCIPAL
// ============================================================

// iniciarServidor configura y pone en funcionamiento
// el servidor web del Sistema de Gestión de Streaming.
func iniciarServidor() {

	// Página principal.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set(
			"Content-Type",
			"text/plain; charset=utf-8",
		)

		fmt.Fprintln(
			w,
			"Sistema de Gestión de Streaming",
		)

		fmt.Fprintln(
			w,
			"Servidor web funcionando correctamente.",
		)

		fmt.Fprintln(
			w,
			"API disponible en /api",
		)
	})

	// Servicio Web #1 y #2.
	// GET  -> consultar contenidos.
	// POST -> crear contenido.
	http.HandleFunc(
		"/api/contenidos",
		func(w http.ResponseWriter, r *http.Request) {

			switch r.Method {

			case http.MethodGet:
				servicioContenidos(w, r)

			case http.MethodPost:
				servicioCrearContenido(w, r)

			default:
				http.Error(
					w,
					"Método no permitido. Utilice GET o POST.",
					http.StatusMethodNotAllowed,
				)
			}
		},
	)

	// Servicio Web #3.
	http.HandleFunc(
		"/api/categorias",
		servicioCategorias,
	)

	// Servicio Web #4.
	// GET  -> consultar usuarios.
	// POST -> crear usuarios.
	http.HandleFunc(
		"/api/usuarios",
		servicioUsuarios,
	)

	// Servicio Web #4.1.
	// PUT    -> actualizar usuario.
	// DELETE -> eliminar usuario.
	http.HandleFunc(
		"/api/usuarios/",
		servicioUsuarios,
	)

	// Servicio Web #5.
	http.HandleFunc(
		"/api/reproduccion",
		servicioReproduccion,
	)

	// Servicio Web #6.
	http.HandleFunc(
		"/api/planes",
		servicioPlanes,
	)

	// Servicio Web #7.
	http.HandleFunc(
		"/api/contenidos/buscar",
		servicioBuscarContenido,
	)

	// Servicios Web #8 y #9.
	// PUT y DELETE.
	http.HandleFunc(
		"/api/contenidos/",
		servicioContenidoPorID,
	)

	fmt.Println(
		"=== SISTEMA DE GESTIÓN DE STREAMING ===",
	)

	fmt.Println(
		"Servidor iniciado en http://localhost:8080",
	)

	err := http.ListenAndServe(
		":8080",
		nil,
	)

	if err != nil {
		fmt.Println(
			"Error al iniciar el servidor:",
			err,
		)
	}
}
