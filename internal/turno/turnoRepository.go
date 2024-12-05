package turno

import (
	"errors"
	"fmt"
	"time"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store/interfaces"
)

type TurnoRespository interface {
	Agregar(turno domain.Turno) (domain.Turno, error)
	AgregarPorDNIyMatricula(dni string, matricula string, fechaHora time.Time, descripcion string) (domain.Turno, error)
	BuscarPorID(id int) (domain.Turno, error)
	Actualizar(turno domain.Turno) (domain.Turno, error)
	Eliminar(id int) error
	BuscarPorDNIPaciente(dni string) ([]domain.Turno, error)
	AgregarPorDyM(turno domain.Turno) error
}

type repository struct {
	storage interfaces.TurnoInterface
}

// Actualizar implements TurnoRespository.
func (r *repository) Actualizar(turno domain.Turno) (domain.Turno, error) {
	//falta validar lo que se le pasa///////// 
	err := r.storage.Actualizar(turno)
	if err!= nil {
		fmt.Println("Error modificando el turno:", err)
        return domain.Turno{}, errors.New("Error modificando el turno")
    }
	return turno, nil
}

// Agregar implements TurnoRespository.
func (r *repository) Agregar(turno domain.Turno) (domain.Turno, error) {
	//falta validar lo que se le pasa///////// 
	id, err := r.storage.Agregar(turno)
	if err!= nil {
        return domain.Turno{}, errors.New("Error creando el turno")
    }

	t, err := r.storage.BuscarPorID(id)
	if err!= nil {
		fmt.Println("Error recuperando el turno:", err)
        return domain.Turno{}, errors.New("Error recuperando el turno")
    }
		
	

	
	return t, nil
}

//Agregar por DNI y Matricula 
func (r *repository) AgregarPorDyM(turno domain.Turno) error {
    //falta validar DNI y matricula
    _, err := r.storage.Agregar(turno)
	if err!= nil {
        return errors.New("Error creando el turno")
    }
	/*
	t, err := r.storage.BuscarPorID(id)
	if err!= nil {
		fmt.Println("Error recuperando el turno:", err)
        return errors.New("Error recuperando el turno")
    }*/
    return nil
}

// AgregarPorDNIyMatricula implements TurnoRespository.
func (r *repository) AgregarPorDNIyMatricula(dni string, matricula string, fechaHora time.Time, descripcion string) (domain.Turno, error) {
	//falta validar DNI y matricula
	id, err := r.storage.AgregarPorDNIyMatricula(dni, matricula, fechaHora, descripcion)
	if err!= nil {
		fmt.Println("Error creando el turno:", err)
        return domain.Turno{}, errors.New("Error creando el turno")
    }
	// Busca el turno recién creado para retornar la estructura completa
	turno, err := r.BuscarPorID(id)
	if err != nil {
		return domain.Turno{}, errors.New("Error recuperando el turno creado")
	}

	return turno, nil
}

// BuscarPorDNIPaciente implements TurnoRespository.
func (r *repository) BuscarPorDNIPaciente(dni string) ([]domain.Turno, error) {
	turnos, err:= r.storage.BuscarPorDNIPaciente(dni)
	if err!= nil {
        return nil, errors.New("Error buscando turnos por DNI del paciente")
	}
    
	return turnos, nil	
}

// BuscarPorID implements TurnoRespository.
func (r *repository) BuscarPorID(id int) (domain.Turno, error) {
	turno, err := r.storage.BuscarPorID(id)
	if err!= nil {
        return turno, errors.New("Turno no encontrado")
    }
	return turno, nil
}

// Eliminar implements TurnoRespository.
func (r *repository) Eliminar(id int) error {
	err := r.storage.Eliminar(id)
    if err!= nil {
        return errors.New("Error eliminando el turno")
    }
    return nil
}

func NewRepositoryTurno(storage interfaces.TurnoInterface) TurnoRespository {
	return &repository{storage}
}
