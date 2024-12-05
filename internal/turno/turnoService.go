package turno

import (
	//"errors"
	"errors"
	"fmt"
	"time"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/dentista"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/paciente"
)

type ServiceTurno interface {
	Agregar(turno domain.Turno) (domain.Turno, error)
	AgregarPorDNIyMatricula(dni string, matricula string, fechaHora time.Time, descripcion string) (domain.Turno, error)
	BuscarPorID(id int) (domain.Turno, error)
	Actualizar(id int, turno domain.Turno) (domain.Turno, error)
	Eliminar(id int) error
	BuscarPorDNIPaciente(dni string) ([]domain.Turno, error)
	AgregarPorDyM(string, matricula string, fechaHora time.Time, descripcion string) (domain.Turno, error)
}

type serviceTurno struct {
	repo TurnoRespository
	pacienteRepo paciente.PacienteRepository
	dentistaRepo dentista.DentistaRepository
	
}
func NewServiceTurno(repo TurnoRespository, pacienteRepo paciente.PacienteRepository, dentistaRepo dentista.DentistaRepository) ServiceTurno {
	return &serviceTurno{repo, pacienteRepo, dentistaRepo}
}

// Actualizar implements ServiceTurno.
func (s *serviceTurno) Actualizar(id int, t domain.Turno) (domain.Turno, error) {
	turno, err := s.repo.BuscarPorID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if t.Paciente.ID > 0 {
		turno.Paciente = t.Paciente
	}
	if t.Dentista.ID > 0 {
        turno.Dentista = t.Dentista
    }
	if t.Descripcion!= "" {
        turno.Descripcion = t.Descripcion
    }
	// Actualizar la fecha si es válida
	if !t.Fecha.IsZero() {
		turno.Fecha = t.Fecha
	}
	///ver que hacer con la fecha si se deja actualizar 
	turno, err = s.repo.Actualizar(turno)
	if err!= nil {
        return domain.Turno{}, err
    }

	return turno, nil
}

// Agregar implements ServiceTurno.
func (s *serviceTurno) Agregar(turno domain.Turno) (domain.Turno, error) {
	  // Validar que el paciente existe
	  _, err := s.pacienteRepo.BuscarPorID(turno.Paciente.ID)
	  if err != nil {
		  // Devolver un error si no se encuentra el paciente
		  return domain.Turno{}, fmt.Errorf("el paciente con ID %d no existe", turno.Paciente.ID)
		  //return domain.Turno{}, errors.New("el paciente no existe")
	  }
	  _, err = s.pacienteRepo.BuscarPorDNI(turno.Paciente.DNI)
	  if err != nil {
		  // Devolver un error si no se encuentra el paciente
		  return domain.Turno{}, fmt.Errorf("el paciente con DNI %s no existe", turno.Paciente.DNI)
		  //return domain.Turno{}, errors.New("El paciente con no existe con id ")
		
	  }
	
	turno, err = s.repo.Agregar(turno)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *serviceTurno) AgregarPorDyM(dni string, matricula string, fechaHora time.Time, descripcion string) (domain.Turno, error) {
	
	// Validar que el paciente existe
	// Sustituir _ pod paciente cuadno este dentista
	paciente, err := s.pacienteRepo.BuscarPorDNI(dni)
	if err != nil {
		return domain.Turno{}, err
		
	}

	//Validar que el dentista exista
	// Sustituir _ con el dentista cuando este paciente
    dentista, err := s.dentistaRepo.GetByMatricula(matricula)
    if err!= nil {
        return domain.Turno{}, err
        
    }
    
    // Validar que la fecha es válida
    if fechaHora.Before(time.Now()) {
        return domain.Turno{}, errors.New("La fecha de la cita no puede ser anterior a la fecha actual")
    }
    
    // Validar que la descripción no está vacía
    if descripcion == "" {
        return domain.Turno{}, errors.New("La descripción de la cita no puede estar vacía")
    }
    //cambiar dentista y paciente
	var turno2 = domain.Turno{
		Paciente: paciente,
        Dentista: dentista,
        Descripcion: descripcion,
        Fecha: fechaHora,
	}
	// cuando este dentista dejar este y sacar el de abajo
	turno, err := s.repo.Agregar(turno2)
	if err!= nil {
        return domain.Turno{}, err
    }
	/*
	turno, err := s.repo.AgregarPorDNIyMatricula(dni, matricula, fechaHora, descripcion)
	
	
	if err!= nil {
        return domain.Turno{}, err
    }*/
    return turno, nil
}



// AgregarPorDNIyMatricula implements ServiceTurno.
func (s *serviceTurno) AgregarPorDNIyMatricula(dni string, matricula string, fechaHora time.Time, descripcion string) (domain.Turno, error) {
	_, err := s.pacienteRepo.BuscarPorDNI(dni)
	if err != nil {
		return domain.Turno{}, err
		
	}
	
	turno, err := s.repo.AgregarPorDNIyMatricula(dni, matricula, fechaHora, descripcion)
    if err!= nil {
        return domain.Turno{}, err
    }
    return turno, nil
}

// BuscarPorDNIPaciente implements ServiceTurno.
func (s *serviceTurno) BuscarPorDNIPaciente(dni string) ([]domain.Turno, error) {
	turnos, err := s.repo.BuscarPorDNIPaciente(dni)
    if err!= nil {
        return nil, err
    }
    return turnos, nil
}

// BuscarPorID implements ServiceTurno.
func (s *serviceTurno) BuscarPorID(id int) (domain.Turno, error) {
	turno, err := s.repo.BuscarPorID(id)
    if err!= nil {
        return domain.Turno{}, err
    }
    return turno, nil
}

// Eliminar implements ServiceTurno.
func (s *serviceTurno) Eliminar(id int) error {
	_, err1 := s.repo.BuscarPorID(id)
    if err1!= nil {
        return errors.New("Turno no encontrado")
    }
	
	err := s.repo.Eliminar(id)
	if err!= nil {
        return err
    }
	return nil
}

