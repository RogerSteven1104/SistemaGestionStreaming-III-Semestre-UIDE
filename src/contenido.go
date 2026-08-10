package main

import "fmt"

// Contenido representa una película o serie disponible
// dentro del Sistema de Gestión de Streaming.
type Contenido struct {
	idContenido   int
	titulo        string
	descripcion   string
	genero        string
	duracion      int
	clasificacion string
}

// NuevoContenido crea un nuevo contenido con los datos recibidos.
func NuevoContenido(
	id int,
	titulo string,
	descripcion string,
	genero string,
	duracion int,
	clasificacion string,
) *Contenido {

	return &Contenido{
		idContenido:   id,
		titulo:        titulo,
		descripcion:   descripcion,
		genero:        genero,
		duracion:      duracion,
		clasificacion: clasificacion,
	}
}

// MostrarInformacion muestra la información principal del contenido.
func (c Contenido) MostrarInformacion() {
	fmt.Println("ID:", c.idContenido)
	fmt.Println("Título:", c.titulo)
	fmt.Println("Descripción:", c.descripcion)
	fmt.Println("Género:", c.genero)
	fmt.Println("Duración:", c.duracion, "minutos")
	fmt.Println("Clasificación:", c.clasificacion)
}

// GetTitulo permite consultar el título del contenido.
func (c Contenido) GetTitulo() string {
	return c.titulo
}

// SetTitulo permite modificar el título validando que no esté vacío.
func (c *Contenido) SetTitulo(titulo string) error {
	if titulo == "" {
		return fmt.Errorf("el título no puede estar vacío")
	}

	c.titulo = titulo
	return nil
}

// GetGenero permite consultar el género del contenido.
func (c Contenido) GetGenero() string {
	return c.genero
}

// SetGenero permite modificar el género validando que no esté vacío.
func (c *Contenido) SetGenero(genero string) error {
	if genero == "" {
		return fmt.Errorf("el género no puede estar vacío")
	}

	c.genero = genero
	return nil
}

// GetDuracion permite consultar la duración del contenido.
func (c Contenido) GetDuracion() int {
	return c.duracion
}

// SetDuracion permite modificar la duración validando que sea mayor que cero.
func (c *Contenido) SetDuracion(duracion int) error {
	if duracion <= 0 {
		return fmt.Errorf("la duración debe ser mayor que cero")
	}

	c.duracion = duracion
	return nil
}

// Reproducir muestra un mensaje indicando que el contenido está siendo reproducido.
func (c Contenido) Reproducir() {
	fmt.Println("Reproduciendo:", c.titulo)
}
