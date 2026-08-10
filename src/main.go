package main

import (
	"errors"
	"fmt"
)

func main() {

	// Creamos un nuevo usuario utilizando el constructor.
	usuario := NuevoUsuario(
		1,
		"Roger",
		"roger@email.com",
		"123456",
		"Premium",
	)

	fmt.Println("Nombre inicial:", usuario.GetNombre())

	// Modificamos el nombre mediante el Setter.
	err := usuario.SetNombre("Carlos")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Nombre después del cambio:", usuario.GetNombre())

	usuario.MostrarInformacion()

	fmt.Println()

	// Creamos un contenido utilizando el constructor.
	contenido := NuevoContenido(
		1,
		"Avatar",
		"Una historia de ciencia ficción en Pandora.",
		"Ciencia ficción",
		162,
		"PG-13",
	)

	contenido.MostrarInformacion()

	fmt.Println()

	// Probamos la validación del título.
	// Actualizamos información del contenido utilizando los Setter.
	err = contenido.SetTitulo("Avatar 2")

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = contenido.SetDuracion(180)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Título actualizado:", contenido.GetTitulo())
	fmt.Println("Duración actualizada:", contenido.GetDuracion())

	fmt.Println()

	// Creamos las categorías del sistema.
	categoria1 := NuevaCategoria(1, "Acción")
	categoria2 := NuevaCategoria(2, "Comedia")
	categoria3 := NuevaCategoria(3, "Ciencia ficción")

	// Agregamos las categorías al sistema.
	err = AgregarCategoria(categoria1)

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AgregarCategoria(categoria2)

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AgregarCategoria(categoria3)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Probamos el registro de una categoría sin nombre.
	categoriaError := NuevaCategoria(4, "")

	err = AgregarCategoria(categoriaError)

	if err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, ErrCategoriaInvalida) {
			fmt.Println("Validación: la categoría ingresada no es válida")
		}
	}

	// Mostramos las categorías registradas.
	ListarCategorias()

	fmt.Println()

	// Creamos una película y una serie.
	pelicula := NuevaPelicula("Avatar 2", 180)
	serie := NuevaSerie("Stranger Things", 5)

	// Guardamos diferentes tipos en una colección de la interfaz.
	elementosReproducibles := []Reproducible{
		contenido,
		pelicula,
		serie,
	}

	// Recorremos la colección y ejecutamos el mismo método
	// sobre diferentes tipos de objetos.
	for _, elemento := range elementosReproducibles {
		elemento.Reproducir()
	}

	fmt.Println()

	// Comprobamos que Pelicula y Serie reutilizan
	// los métodos de Contenido mediante la incrustación.
	fmt.Println("Título de la película:", pelicula.GetTitulo())
	fmt.Println("Título de la serie:", serie.GetTitulo())

}
