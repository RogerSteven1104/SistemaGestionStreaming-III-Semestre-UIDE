package main

import "fmt"

// Categoria representa una categoría utilizada
// para organizar los contenidos de la plataforma.
type Categoria struct {
	idCategoria     int
	nombreCategoria string
}

// NuevaCategoria crea una nueva categoría con los datos recibidos.
func NuevaCategoria(id int, nombre string) *Categoria {
	return &Categoria{
		idCategoria:     id,
		nombreCategoria: nombre,
	}
}

// GetNombre permite consultar el nombre de la categoría.
func (c Categoria) GetNombre() string {
	return c.nombreCategoria
}

// SetNombre permite modificar el nombre de la categoría.
func (c *Categoria) SetNombre(nombre string) error {
	if nombre == "" {
		return fmt.Errorf("el nombre de la categoría no puede estar vacío")
	}

	c.nombreCategoria = nombre
	return nil
}

// categorias almacena las categorías registradas en el sistema.
var categorias []*Categoria

// AgregarCategoria incorpora una categoría al listado del sistema.
func AgregarCategoria(categoria *Categoria) error {
	if categoria == nil {
		return fmt.Errorf("no se puede agregar una categoría vacía")
	}

	if categoria.GetNombre() == "" {
		return fmt.Errorf("no se puede agregar una categoría sin nombre")
	}

	categorias = append(categorias, categoria)

	return nil
}

// ListarCategorias muestra todas las categorías registradas.
func ListarCategorias() {
	fmt.Println("=== CATEGORIAS REGISTRADAS ===")

	for _, categoria := range categorias {
		fmt.Println(
			categoria.idCategoria,
			"-",
			categoria.GetNombre(),
		)
	}
}
