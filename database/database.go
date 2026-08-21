package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Conectar abre la conexión con SQLite.
func Conectar() error {

	var err error

	DB, err = sql.Open("sqlite", "./streaming.db")

	if err != nil {
		return err
	}

	err = DB.Ping()

	if err != nil {
		return err
	}

	fmt.Println("Base de datos conectada correctamente")

	err = crearTablas()

	if err != nil {
		return err
	}

	return cargarDatosIniciales()
}

// crearTablas crea las tablas necesarias para el sistema.
func crearTablas() error {

	consulta := `
	CREATE TABLE IF NOT EXISTS contenidos (
		id INTEGER PRIMARY KEY,
		titulo TEXT NOT NULL,
		descripcion TEXT NOT NULL,
		genero TEXT NOT NULL,
		duracion INTEGER NOT NULL,
		clasificacion TEXT NOT NULL
	);
	`

	_, err := DB.Exec(consulta)

	if err != nil {
		return err
	}

	fmt.Println("Tabla contenidos creada correctamente")

	return nil
}

// cargarDatosIniciales agrega los contenidos iniciales
// únicamente si la tabla está vacía.
func cargarDatosIniciales() error {

	var cantidad int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM contenidos",
	).Scan(&cantidad)

	if err != nil {
		return err
	}

	// Si ya existen datos, no insertar nuevamente.
	if cantidad > 0 {
		fmt.Println("La base de datos ya contiene información")
		return nil
	}

	consulta := `
	INSERT INTO contenidos
	(id, titulo, descripcion, genero, duracion, clasificacion)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	contenidosIniciales := []struct {
		id            int
		titulo        string
		descripcion   string
		genero        string
		duracion      int
		clasificacion string
	}{
		{
			1,
			"Avatar 2",
			"Una historia de ciencia ficción en Pandora.",
			"Ciencia ficción",
			180,
			"PG-13",
		},
		{
			2,
			"John Wick",
			"Un antiguo asesino regresa a la acción.",
			"Acción",
			101,
			"R",
		},
		{
			3,
			"Son como niños",
			"Un grupo de amigos se reúne nuevamente.",
			"Comedia",
			102,
			"PG-13",
		},
	}

	for _, contenido := range contenidosIniciales {

		_, err := DB.Exec(
			consulta,
			contenido.id,
			contenido.titulo,
			contenido.descripcion,
			contenido.genero,
			contenido.duracion,
			contenido.clasificacion,
		)

		if err != nil {
			return err
		}
	}

	fmt.Println("Datos iniciales cargados correctamente")

	return nil
}
