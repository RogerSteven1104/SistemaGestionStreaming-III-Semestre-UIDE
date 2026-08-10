package main

import "fmt"

// Serie representa una serie disponible para reproducirse
// dentro de la plataforma de streaming.
type Serie struct {
	titulo     string
	temporadas int
}

// NuevaSerie crea una nueva serie con su título y número de temporadas.
func NuevaSerie(titulo string, temporadas int) *Serie {
	return &Serie{
		titulo:     titulo,
		temporadas: temporadas,
	}
}

// Reproducir muestra el mensaje correspondiente a una serie.
func (s Serie) Reproducir() {
	fmt.Println("Reproduciendo serie:", s.titulo)
}
