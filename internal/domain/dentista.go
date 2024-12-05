package domain

type Dentista struct {
	ID int `json:"id"`
	Nombre string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}