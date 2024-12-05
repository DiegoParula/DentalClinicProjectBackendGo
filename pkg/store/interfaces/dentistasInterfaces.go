package interfaces

import "github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"

type DentistaInterface interface {
	// Read devuelve un dentista por su id
	Read(id int) (domain.Dentista, error)
	//Obtener todos los dentistas
	GetAll() ([]domain.Dentista, error)
	// Busca dentista por matrícula
	GetByMatricula(matricula string) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(dentista domain.Dentista) error
	// Update actualiza un dentista
	Update(dentista domain.Dentista) error
	// Actualiza dentista en forma parcial
	Patch(dentista domain.Dentista) error
	// Delete elimina un dentista
	Delete(id int) error
}
