package main

import "fmt"

// Usuario representa a una persona registrada
// dentro del Sistema de Gestión de Streaming.
type Usuario struct {
	idUsuario       int
	nombre          string
	correo          string
	contrasena      string
	planSuscripcion string
}

// NuevoUsuario crea y devuelve un nuevo usuario
// con los datos recibidos como parámetros.
func NuevoUsuario(id int, nombre string, correo string, contrasena string, plan string) *Usuario {
	return &Usuario{
		idUsuario:       id,
		nombre:          nombre,
		correo:          correo,
		contrasena:      contrasena,
		planSuscripcion: plan,
	}
}

// MostrarInformacion muestra los datos básicos del usuario.
func (u Usuario) MostrarInformacion() {
	fmt.Println("ID:", u.idUsuario)
	fmt.Println("Nombre:", u.nombre)
	fmt.Println("Correo:", u.correo)
	fmt.Println("Plan:", u.planSuscripcion)
}

// GetNombre permite consultar el nombre del usuario.
func (u Usuario) GetNombre() string {
	return u.nombre
}

// SetNombre permite modificar el nombre del usuario.
func (u *Usuario) SetNombre(nombre string) error {
	if nombre == "" {
		return ErrNombreVacio
	}

	u.nombre = nombre
	return nil
}
