package domain

import "time"

// RequestData define los datos para crear un nuevo turno
type RequestData struct {
    DNI         string    `json:"dni" binding:"required"`
    Matricula   string    `json:"matricula" binding:"required"`
    Fecha       time.Time `json:"fecha" binding:"required"`
    Descripcion string    `json:"descripcion" binding:"required"`

}