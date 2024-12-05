package dentista

import (
	"errors"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"fmt"
)

type DentistaService interface {
	//Obtener todos los dentistas
	GetAll() ([]domain.Dentista, error)
	// Busca dentista por dni
	GetByMatricula(matricula string) (domain.Dentista, error)
	// GetByID busca un dentista por su matricula
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(p domain.Dentista) (domain.Dentista, error)
	// Delete elimina un dentista
	Delete(id int) error
	// Update actualiza un dentista
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	// Actualiza dentista en forma parcial
	Patch(id int, p domain.Dentista) (domain.Dentista, error)
}

type dentistaService struct {
	dr DentistaRepository
}



// NewService crea un nuevo servicio
func NewService(dr DentistaRepository) DentistaService {
	return &dentistaService{dr}
}

// GetAll implements DentistaService.
func (d *dentistaService) GetAll() ([]domain.Dentista, error) {
	dentistas, err:= d.GetAll()
	if err!= nil {
        return nil, err
    }
	return dentistas, nil
}

// GetByDni implements DentistaService.
func (d *dentistaService) GetByMatricula(matricula string) (domain.Dentista, error) {
	dentista, err := d.dr.GetByMatricula(matricula)
	if err!= nil {
        return domain.Dentista{}, err
    }
	return dentista, nil
}

// GetByID busca un dentista por su id
func (d *dentistaService) GetByID(id int) (domain.Dentista, error) {
    dentista, err := d.dr.GetByID(id)
    if err != nil {
        if err.Error() == fmt.Sprintf("no se encontró dentista con id %d", id) {
            return domain.Dentista{}, errors.New("dentista not found")
        }
        return domain.Dentista{}, err
    }
    return dentista, nil
}


// Create agrega un nuevo dentista
func (d *dentistaService) Create(dentista domain.Dentista) (domain.Dentista, error) {
	// Verifica si el dentista ya existe
	existingDentista, err := d.dr.GetByID(dentista.ID)
	if err == nil && existingDentista.ID != 0 {
		return domain.Dentista{}, errors.New("dentista already exists")
	}

	// Crea el nuevo dentista
	newDentista, err := d.dr.Create(dentista)
	if err != nil {
		return domain.Dentista{}, err
	}
	return newDentista, nil
}

// Update actualiza un dentista
func (d *dentistaService) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	// Verifica si el dentista existe
	existingDentista, err := d.dr.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if existingDentista.ID == 0 {
		return domain.Dentista{}, errors.New("dentista does not exist")
	}

	// Aplica la actualización completa
	u.ID = id
	updatedDentista, err := d.dr.Update(id, u)
	if err != nil {
		return domain.Dentista{}, err
	}
	return updatedDentista, nil
}


// Delete elimina un dentista
func (d *dentistaService) Delete(id int) error {
	// Llamar al método Delete del repositorio
	err := d.dr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// Patch actualiza parcialmente un dentista
func (d *dentistaService) Patch(id int, p domain.Dentista) (domain.Dentista, error) {
	// Verifica si el dentista existe antes de hacer un parcheo
	existingDentista, err := d.dr.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if existingDentista.ID == 0 {
		return domain.Dentista{}, errors.New("dentista does not exist")
	}

	// Aplica la actualización parcial
	if p.Nombre != "" {
		existingDentista.Nombre = p.Nombre
	}
	if p.Apellido != "" {
		existingDentista.Apellido = p.Apellido
	}
	if p.Matricula != "" {
		existingDentista.Matricula = p.Matricula
	}

	updatedDentista, err := d.dr.Patch(id, existingDentista)
	if err != nil {
		return domain.Dentista{}, err
	}
	return updatedDentista, nil
}
