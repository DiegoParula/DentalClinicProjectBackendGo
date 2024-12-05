package paciente

import (
	"errors"
	"log"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store/interfaces"
)

type PacienteRepository interface {
	Agregar(paciente domain.Paciente) (domain.Paciente, error)
	Listar() ([]domain.Paciente, error)
	Modificar(id int, paciente domain.Paciente) (domain.Paciente, error)
	Eliminar(id int) error
	BuscarPorID(id int) (domain.Paciente, error)
	BuscarPorDNI(dni string) (domain.Paciente, error)
}

type repository struct {
	storage interfaces.PacienteInterface
}
func NewRepository(storage interfaces.PacienteInterface) PacienteRepository {
	return &repository{storage}
}

// Agregar 
func (r *repository) Agregar(paciente domain.Paciente) (domain.Paciente, error) {
	// Verificar si el DNI ya existe
	if r.storage.ExisteDNI(paciente.DNI) {
		return domain.Paciente{}, errors.New("DNI de paciente ya existe")
	}
	/*if r.storage.ExisteDNI(paciente.DNI){
		return domain.Paciente{}, errors.New("DNI de paciente ya existe")
	}*/
	id, err:= r.storage.Agregar(paciente)
	if err!= nil {
		log.Printf("Error creando paciente: %v", err)
        return domain.Paciente{}, errors.New("Error creando paciente")
    }

	p, err:= r.storage.BuscarPorID(id)
	if err!= nil {
        log.Printf("Error recuperando paciente: %v", err)
        return domain.Paciente{}, errors.New("Error recuperando paciente")
    }
	return p, nil
}

// BuscarPorDNI 
func (r *repository) BuscarPorDNI(dni string) (domain.Paciente, error) {
	paciente, err := r.storage.BuscarPorDNI(dni)
	if err!= nil {
        return domain.Paciente{}, errors.New("Paciente no encontrado")
    }
	return paciente, nil
}

// BuscarPorID 
func (r *repository) BuscarPorID(id int) (domain.Paciente, error) {
	paciente, err:= r.storage.BuscarPorID(id)
	if err!= nil {
        return domain.Paciente{}, errors.New("Paciente no encontrado")
    }
	return paciente, nil
}

// Eliminar 
func (r *repository) Eliminar(id int) error {
	err := r.storage.Eliminar(id)
	if err!= nil {
        return errors.New("Error eliminando paciente")
    }
	return nil
}

// Listar 
func (r *repository) Listar() ([]domain.Paciente, error) {
	pacientes, err:= r.storage.Listar()
    if err!= nil {
        return nil, errors.New("Error listando pacientes")
    }
    return pacientes, nil
}

// Modificar 
func (r *repository) Modificar(id int, paciente domain.Paciente) (domain.Paciente, error) {
	if !r.storage.Existe(id){
		return domain.Paciente{}, errors.New("el paciente no existe")
	}
	err:= r.storage.Modificar(paciente)
	if err!= nil {
		log.Printf("Error modificando paciente: %v", err)
        return domain.Paciente{}, errors.New("Error modificando paciente")
    }
	return paciente, nil
}


