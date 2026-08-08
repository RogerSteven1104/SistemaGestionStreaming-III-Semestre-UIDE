package main

import "fmt"

func main() {

	usuario := NuevoUsuario(
		1,
		"Roger",
		"roger@gmail.com",
		"123456",
		"Premium",
	)

	fmt.Println("Nombre inicial:", usuario.GetNombre())

	err := usuario.SetNombre("Carlos")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Nombre después del cambio:", usuario.GetNombre())

	usuario.MostrarInformacion()
}
