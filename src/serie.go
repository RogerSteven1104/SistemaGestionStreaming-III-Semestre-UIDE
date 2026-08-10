package main

import "fmt"

// Serie representa una serie y reutiliza
// los atributos y métodos de Contenido mediante incrustación.
type Serie struct {
	Contenido
	temporadas int
}

// NuevaSerie crea una nueva serie utilizando
// la información general definida en Contenido.
func NuevaSerie(titulo string, temporadas int) *Serie {
	return &Serie{
		Contenido: Contenido{
			titulo: titulo,
		},
		temporadas: temporadas,
	}
}

// Reproducir muestra el mensaje correspondiente a una serie.
func (s Serie) Reproducir() {
	fmt.Println("Reproduciendo serie:", s.titulo)
}
