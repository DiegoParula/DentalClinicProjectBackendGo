package paciente

import (
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	
)

type Service interface {
	Agregar(paciente domain.Paciente) (domain.Paciente, error)
	Listar() ([]domain.Paciente, error)
	Modificar(id int, paciente domain.Paciente) (domain.Paciente, error)
	Eliminar(id int) error
	BuscarPorID(id int) (domain.Paciente, error)
	BuscarPorDNI(dni string) (domain.Paciente, error)
}

type service struct {
	r PacienteRepository
}

func NuevoServicio(r PacienteRepository) Service {
	return &service{r}
}


// Agregar 
func (s *service) Agregar(paciente domain.Paciente) (domain.Paciente, error) {
	paciente, err := s.r.Agregar(paciente)
	if err!= nil {
        return domain.Paciente{}, err
    }
	return paciente, nil
}

// BuscarPorDNI 
func (s *service) BuscarPorDNI(dni string) (domain.Paciente, error) {
	paciente, err := s.r.BuscarPorDNI(dni)
	if err!= nil {
        return domain.Paciente{}, err
    }
	return paciente, nil
}

// BuscarPorID 
func (s *service) BuscarPorID(id int) (domain.Paciente, error) {
	paciente, err := s.r.BuscarPorID(id)
	if err!= nil {
        return domain.Paciente{}, err
    }
	return paciente, nil
}

// Eliminar 
func (s *service) Eliminar(id int) error {
	_, err1 := s.r.BuscarPorID(id)
	if err1!= nil {
        return err1
}

	err := s.r.Eliminar(id)
    if err!= nil {
        return err
    }
    return nil
}

// Listar 
func (s *service) Listar() ([]domain.Paciente, error) {
	pacientes, err:= s.r.Listar()
	if err!= nil {
        return nil, err
    }
	return pacientes, nil
}

// Modificar 
func (s *service) Modificar(id int, u domain.Paciente) (domain.Paciente, error) {
	paciente, err:= s.r.BuscarPorID(id)
	if err!= nil {
        return domain.Paciente{}, err
    }

	if u.Nombre != ""{
		paciente.Nombre = u.Nombre
	}
	if u.Apellido!= ""{
        paciente.Apellido = u.Apellido
    }
	if u.Direccion!= ""{
        paciente.Direccion = u.Direccion
    }
	if u.DNI!= ""{
        paciente.DNI = u.DNI
    }
	/*if u.FechaAlta!= ""{
        paciente.FechaAlta = u.FechaAlta
    }*///ver que hacer con la fecha si se deja actualizar 
	paciente, err = s.r.Modificar(id, paciente)
	if err!= nil {
		return domain.Paciente{}, err
    }
	return paciente, nil	
}

