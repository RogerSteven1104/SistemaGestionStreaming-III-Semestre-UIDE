package main

import "fmt"

// Pelicula representa una película y reutiliza
// los atributos y métodos de Contenido mediante incrustación.
type Pelicula struct {
	Contenido
}

// NuevaPelicula crea una nueva película utilizando
// la información general definida en Contenido.
func NuevaPelicula(
	titulo string,
	duracion int,
) *Pelicula {

	return &Pelicula{
		Contenido: Contenido{
			titulo:   titulo,
			duracion: duracion,
		},
	}
}

// Reproducir muestra el mensaje correspondiente a una película.
func (p Pelicula) Reproducir() {
	fmt.Println("Reproduciendo película:", p.titulo)
}
