package main

import "fmt"

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

}
