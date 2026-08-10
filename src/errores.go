package main

import "errors"

// Errores utilizados para validar los datos del sistema.
var (
	ErrNombreVacio       = errors.New("el nombre no puede estar vacío")
	ErrTituloVacio       = errors.New("el título no puede estar vacío")
	ErrDuracionInvalida  = errors.New("la duración debe ser mayor que cero")
	ErrCategoriaInvalida = errors.New("la categoría no puede estar vacía")
)
