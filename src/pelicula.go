package main

import "fmt"

// Pelicula representa una película que puede ser reproducida
// dentro de la plataforma de streaming.
type Pelicula struct {
	titulo   string
	duracion int
}

// NuevaPelicula crea una nueva película con su título y duración.
func NuevaPelicula(titulo string, duracion int) *Pelicula {
	return &Pelicula{
		titulo:   titulo,
		duracion: duracion,
	}
}

// Reproducir muestra el mensaje correspondiente a una película.
func (p Pelicula) Reproducir() {
	fmt.Println("Reproduciendo película:", p.titulo)
}
