package domain

import "time"

type Paciente struct {
	ID int `json:"id"`
	Nombre string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Direccion string `json:"direccion" binding:"required"`
	DNI string `json:"dni" binding:"required"`
	FechaAlta time.Time `json:"fechaAlta"`
}