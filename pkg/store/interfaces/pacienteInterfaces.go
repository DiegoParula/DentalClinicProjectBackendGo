package interfaces

import "github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"

type PacienteInterface interface {
	Agregar(paciente domain.Paciente) (int,error)
	Listar() ([]domain.Paciente, error)
	Modificar(paciente domain.Paciente) error
	Eliminar(id int) error
	BuscarPorID(id int) (domain.Paciente, error)
	BuscarPorDNI(dni string) (domain.Paciente, error)
	Existe(id int) bool
	ExisteDNI(dni string) bool
}