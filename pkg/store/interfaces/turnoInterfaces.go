package interfaces

import (
	"time"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
)

type TurnoInterface interface {
	Agregar(turno domain.Turno) (int, error)
	AgregarPorDNIyMatricula(dni string, matricula string, fechaHora time.Time, descripcion string) (int, error)
	
	BuscarPorID(id int) (domain.Turno, error)
	Actualizar(turno domain.Turno) error
	Eliminar(id int) error
	BuscarPorDNIPaciente(dni string) ([]domain.Turno, error)

}
