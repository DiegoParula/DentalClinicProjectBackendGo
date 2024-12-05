package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"

	"time"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"

	storeinterface "github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store/interfaces"
)

type sqlStoreTurno struct {
	db *sql.DB
}

// Actualizar implements interfaces.TurnoInterface.
func (s *sqlStoreTurno) Actualizar(turno domain.Turno) error {
	query := "UPDATE turnos SET paciente_id = ?, dentista_id = ?, fecha_hora = ?, descripcion = ? WHERE id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	fechaHoraStr := turno.Fecha.Format("2006-01-02T15:04:05")
	res, err := stmt.Exec(turno.Paciente.ID, turno.Dentista.ID, fechaHoraStr, turno.Descripcion, turno.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return err
}

// Agregar
func (s *sqlStoreTurno) Agregar(turno domain.Turno) (int, error) {

	query := "INSERT INTO turnos (fecha_hora, descripcion, paciente_id, dentista_id) VALUES (?, ?, ?, ?)"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	fechaHoraStr := turno.Fecha.Format("2006-01-02T15:04:05")
	res, err := stmt.Exec(fechaHoraStr, turno.Descripcion, turno.Paciente.ID, turno.Dentista.ID)
	if err != nil {
		//fmt.Printf("Insertando turno con pacienteID: %d, dentistaID: %d, fecha: %s, descripcion: %s", turno.Paciente.ID, turno.Dentista.ID, turno.Fecha.Format("2006-01-02T15:04:05"), turno.Descripcion)
		return 0, err
	}
	_, err = res.RowsAffected()
	if err != nil {

		return 0, err
	}
	
		id, err := res.LastInsertId()
		if err != nil {
			return 0,err
		}
	return int(id),nil
}

// AgregarPorDNIyMatricula
func (s *sqlStoreTurno) AgregarPorDNIyMatricula(dni string, matricula string, fecha time.Time, descripcion string) (int, error) {

	/*var pacienteID int
	query := "SELECT id FROM pacientes WHERE dni = ?"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&pacienteID)
	if err != nil {
		return 0, errors.New("Paciente no encontrado")
	}

	var dentistaID int
	queryD := "SELECT id FROM dentistas WHERE matricula = ?"
	rowD := s.db.QueryRow(queryD, matricula)
	errD := rowD.Scan(&dentistaID)
	if errD != nil {
		return 0, errors.New("Paciente no encontrado")
	}*/

	// Primero, buscar el paciente por DNI
	var paciente domain.Paciente
	var fechaHoraStr string
	err := s.db.QueryRow("SELECT id, nombre, apellido, direccion, dni, fecha_alta FROM pacientes WHERE dni = ?", dni).
		Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Direccion, &paciente.DNI, &fechaHoraStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("paciente no encontrado")
		}
		return 0, err
	}

	// Luego, buscar el dentista por matrícula
	var dentista domain.Dentista
	err = s.db.QueryRow("SELECT id, nombre, apellido, matricula FROM dentistas WHERE matricula = ?", matricula).
		Scan(&dentista.ID, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("dentista no encontrado")
		}
		return 0, err
	}

	// Crear el turno
	/*
		var turno domain.Turno
		turno.Paciente = paciente
		turno.Dentista = dentista
		turno.Fecha = fecha
		turno.Descripcion = descripcion*/
	/*
		turno := &domain.Turno{
			Paciente:    paciente,
			Dentista:    dentista,
			Fecha:       fecha,
			Descripcion: descripcion,
		}*/
	// Llamar al método Agregar y obtener el ID del nuevo turno
	/*err = s.Agregar(turno)
	if err != nil {
		return 0, err
	}

	return 0,nil*/

	queryt := "INSERT INTO turnos (fecha_hora, descripcion, paciente_id, dentista_id) VALUES (?, ?, ?, ?)"
	stmt, err := s.db.Prepare(queryt)
	if err != nil {
		return 0, err
	}

	// Parsear la fecha y hora
	var fecha2 time.Time
	fecha2, err = time.Parse("2006-01-02", fechaHoraStr)
	if err != nil {
		return 0, err
	}

	//fechaHoraStr := fecha.Format("2006-01-02T15:04:05")
	res, err := stmt.Exec(fecha2, descripcion, paciente.ID, dentista.ID)
	if err != nil {
		//fmt.Printf("Insertando turno con pacienteID: %d, dentistaID: %d, fecha: %s, descripcion: %s", turno.Paciente.ID, turno.Dentista.ID, turno.Fecha.Format("2006-01-02T15:04:05"), turno.Descripcion)
		return 0, err
	}
	_, err = res.RowsAffected()
	if err != nil {

		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// BuscarPorDNIPaciente
func (s *sqlStoreTurno) BuscarPorDNIPaciente(dni string) ([]domain.Turno, error) {
	query := `
	SELECT t.id, t.paciente_id, t.dentista_id, t.fecha_hora, t.descripcion,
	       p.nombre AS paciente_nombre, p.apellido AS paciente_apellido, p.dni,
	       d.nombre AS dentista_nombre, d.apellido AS dentista_apellido, d.matricula
	FROM turnos t
	JOIN pacientes p ON t.paciente_id = p.id
	JOIN dentistas d ON t.dentista_id = d.id
	WHERE p.dni = ?`

	rows, err := s.db.Query(query, dni)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var turnos []domain.Turno
	for rows.Next() {
		var turno domain.Turno
		var fechaHoraStr string

		// Escaneamos los datos, incluyendo los subcampos de Paciente y Dentista
		err := rows.Scan(
			&turno.ID, &turno.Paciente.ID, &turno.Dentista.ID,
			&fechaHoraStr, &turno.Descripcion,
			&turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.DNI,
			&turno.Dentista.Nombre, &turno.Dentista.Apellido, &turno.Dentista.Matricula)
		if err != nil {
			fmt.Printf("Error escaneando fila: %v\n", err) // para debug error
			return nil, err
		}

		// Parseamos la fecha
		turno.Fecha, err = time.Parse("2006-01-02 15:04:05", fechaHoraStr)
		if err != nil {
			return nil, err
		}

		turnos = append(turnos, turno)
	}

	return turnos, nil
}

// BuscarPorID
func (s *sqlStoreTurno) BuscarPorID(id int) (domain.Turno, error) {
	query := `
		SELECT t.id, t.fecha_hora, t.descripcion,
			   p.id, p.nombre, p.apellido, p.dni,
			   d.id, d.nombre, d.apellido, d.matricula
		FROM turnos t
		JOIN pacientes p ON t.paciente_id = p.id
		JOIN dentistas d ON t.dentista_id = d.id
		WHERE t.id = ?`
	row := s.db.QueryRow(query, id)

	var turno domain.Turno
	var fechaHoraStr string

	err := row.Scan(
		&turno.ID, &fechaHoraStr, &turno.Descripcion,
		&turno.Paciente.ID, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.DNI,
		&turno.Dentista.ID, &turno.Dentista.Nombre, &turno.Dentista.Apellido, &turno.Dentista.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Turno{}, fmt.Errorf("no se encontró el turno con el ID %d", id)
		}
		return domain.Turno{}, err
	}

	// Parsear la fecha y hora
	turno.Fecha, err = time.Parse("2006-01-02 15:04:05", fechaHoraStr)
	if err != nil {
		return domain.Turno{}, err
	}

	return turno, nil
}

// Eliminar
func (s *sqlStoreTurno) Eliminar(id int) error {
	_, err := s.db.Exec("DELETE FROM turnos WHERE id = ?", id)
	return err
}

func NewSqlStoreTurno(db *sql.DB) storeinterface.TurnoInterface {
	return &sqlStoreTurno{db: db}
}
