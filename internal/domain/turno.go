package domain

import "time"

type Turno struct {
	ID int `json:"id"`
	Paciente Paciente `json:"paciente"`
	Dentista Dentista `json:"dentista"`
	//PacienteID int `json:"pacienteId" binding:"required"`
	//DentistaID int `json:"dentistaId" binding:"required"`
	Fecha time.Time `json:"fecha"`
	Descripcion string `json:"descripcion" binding:"required"`
	
}
/*type TurnoInput struct {
    DNI         string    `json:"dni"`
    Matricula   string    `json:"matricula"`
    Fecha       time.Time `json:"fecha"`
    Descripcion string    `json:"descripcion"`
}*/