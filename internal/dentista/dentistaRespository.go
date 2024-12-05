package dentista

import (
	"errors"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store/interfaces"
)

type DentistaRepository interface {
	//Obtener todos los dentistas
	GetAll() ([]domain.Dentista, error)
	// Busca dentista por dni
	GetByMatricula(matricula string) (domain.Dentista, error)
	// GetByID busca un dentista por su matricula
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(d domain.Dentista) (domain.Dentista, error)
	// Update actualiza completamente un dentista
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	// Patch actualiza parcialmente un dentista
	Patch(id int, d domain.Dentista) (domain.Dentista, error)
	// Delete elimina un dentista
	Delete(id int) error
}

type dentistaRepository struct {
	storage interfaces.DentistaInterface
}



func NewDentistaRespository(storage interfaces.DentistaInterface) DentistaRepository {
	return &dentistaRepository{storage}
}

// GetAll implements DentistaRepository.
func (r *dentistaRepository) GetAll() ([]domain.Dentista, error) {
	dentistas, err:= r.storage.GetAll()
    if err!= nil {
        return nil, errors.New("Error listando dentistas")
    }
    return dentistas, nil
}

// GetByDni implements DentistaRepository.
func (r *dentistaRepository) GetByMatricula(matricula string) (domain.Dentista, error) {
	dentista, err := r.storage.GetByMatricula(matricula)
	if err!= nil {
        return domain.Dentista{}, errors.New("Dentista no encontrado")
    }
	return dentista, nil
}

// Create implements DentistaRepository.
func (r *dentistaRepository) Create(d domain.Dentista) (domain.Dentista, error) {
	existingDentista, err := r.storage.Read(d.ID)
	if err == nil && existingDentista.ID != 0 {
		return domain.Dentista{}, errors.New("dentista already exists")
	}

	// Si no existe, procede a crear el nuevo dentista
	err = r.storage.Create(d)
	if err != nil {
		return domain.Dentista{}, errors.New("error creating dentista")
	}
	return d, nil
}

// Delete implements DentistaRepository.
func (d *dentistaRepository) Delete(id int) error {
	err := d.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetByID implements DentistaRepository.
func (d *dentistaRepository) GetByID(id int) (domain.Dentista, error) {
    dentista, err := d.storage.Read(id)
    if err != nil {
        return domain.Dentista{}, err
    }
    return dentista, nil
}


// Patch implements DentistaRepository.
func (r *dentistaRepository) Patch(id int, d domain.Dentista) (domain.Dentista, error) {
	// Verificar si el dentista existe antes de intentar hacer un parcheo
	existingDentista, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if existingDentista.ID == 0 {
		return domain.Dentista{}, errors.New("dentista does not exist")
	}
	// Aplicar actualización parcial
	err = r.storage.Patch(d)
	if err != nil {
		return domain.Dentista{}, errors.New("error patching dentista")
	}
	return d, nil

}

// Update implements DentistaRepository.
func (r *dentistaRepository) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	// Verificar si el dentista existe antes de intentar hacer una actualización
	existingDentista, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if existingDentista.ID == 0 {
		return domain.Dentista{}, errors.New("dentista does not exist")
	}

	// Aplicar actualización completa
	d.ID = id
	err = r.storage.Update(d)
	if err != nil {
		return domain.Dentista{}, errors.New("error updating dentista")
	}
	return d, nil
}
