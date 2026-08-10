package main

import "fmt"

// Reproducible define el comportamiento que debe tener
// cualquier contenido que pueda ser reproducido.
type Reproducible interface {
	Reproducir()
}

// contenidosReproducibles almacena los diferentes tipos
// de contenido que pueden reproducirse en la plataforma.
var contenidosReproducibles []Reproducible

// AgregarReproducible incorpora un elemento al catálogo.
func AgregarReproducible(elemento Reproducible) error {
	if elemento == nil {
		return fmt.Errorf("no se puede agregar un contenido vacío")
	}

	contenidosReproducibles = append(contenidosReproducibles, elemento)

	return nil
}

// ReproducirCatalogo reproduce todos los elementos registrados.
func ReproducirCatalogo() {
	fmt.Println("=== CATÁLOGO DE REPRODUCCIÓN ===")

	for _, elemento := range contenidosReproducibles {
		elemento.Reproducir()
	}
}
