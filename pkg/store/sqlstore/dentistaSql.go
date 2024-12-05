package sqlstore

import (
	"database/sql"
	"fmt"
    "log"
	//"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/dentista"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store/interfaces"
)

type dentistaSql struct {
	db *sql.DB
}



func NewSqlDentista(db *sql.DB) interfaces.DentistaInterface {
	return &dentistaSql{
		db: db,
}
}

// Read devuelve un dentista por su id
func (s *dentistaSql) Read(id int) (domain.Dentista, error) {
    var dentista domain.Dentista
    query := "SELECT id, nombre, apellido, matricula FROM dentistas WHERE id = ?"
    row := s.db.QueryRow(query, id)
    
    err := row.Scan(&dentista.ID, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("No se encontró dentista con id %d", id)
            return dentista, fmt.Errorf("no se encontró dentista con id %d", id)
        }
        log.Printf("Error escaneando dentista: %v", err)
        return dentista, fmt.Errorf("error escaneando dentista: %v", err)
    }
    
    return dentista, nil
}



// GetByDni implements interfaces.DentistaInterface.
func (s *dentistaSql) GetByMatricula(matricula string) (domain.Dentista, error) {
	var dentista domain.Dentista
	query := "SELECT * FROM dentistas WHERE matricula=?"
	row := s.db.QueryRow(query, matricula)
	err := row.Scan(&dentista.ID, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	
	if err != nil {
		return domain.Dentista{}, err
	}


	return dentista, nil
}

// GetAll implements interfaces.DentistaInterface.
func (s *dentistaSql) GetAll() ([]domain.Dentista, error) {
	listReturn := []domain.Dentista{}
	query := "SELECT * FROM dentistas"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var dentista domain.Dentista
	

		err := rows.Scan(&dentista.ID, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
		if err != nil {
			log.Fatal(err)
		}
		listReturn = append(listReturn, dentista) 
	}

	return listReturn, nil
}


// Create agrega un nuevo dentista
func (s *dentistaSql) Create(dentista domain.Dentista) error {
	query := "INSERT INTO dentistas(nombre, apellido, matricula) VALUES(?, ?, ?)"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(dentista.Nombre, dentista.Apellido, dentista.Matricula)
	if err != nil {
		return fmt.Errorf("error ejecutando query: %v", err)
	}
	return nil
}

// Delete elimina un dentista
func (s *dentistaSql) Delete(id int) error {
	query := "DELETE FROM dentistas WHERE id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando la consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error ejecutando la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error obteniendo el número de filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("dentista con id %d no encontrado", id)
	}

	return nil
}

// Patch actualiza parcialmente un dentista
func (s *dentistaSql) Patch(dentista domain.Dentista) error {
	query := "UPDATE dentistas SET nombre = ?, apellido = ? WHERE id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(dentista.Nombre, dentista.Apellido, dentista.ID)
	if err != nil {
		return fmt.Errorf("error ejecutando query: %v", err)
	}
	return nil
}

// Update actualiza un dentista
func (s *dentistaSql) Update(dentista domain.Dentista) error {
	query := "UPDATE dentistas SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(dentista.Nombre, dentista.Apellido, dentista.Matricula, dentista.ID)
	if err != nil {
		return fmt.Errorf("error ejecutando query: %v", err)
	}
	return nil
}

