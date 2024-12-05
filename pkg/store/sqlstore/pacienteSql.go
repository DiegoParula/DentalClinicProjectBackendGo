package sqlstore

import (
	"database/sql"
	"log"
	"time"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	storeinterface "github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store/interfaces" //cambio el alias del paquete, ya que "interface" es una palabra clave en Go y no debe ser usada como nombre de paquete.
)

type sqlStore struct {
	db *sql.DB
}
func NewSqlStorePaciente(db *sql.DB) storeinterface.PacienteInterface {
	return &sqlStore{db: db}

}
// Existe 
func (s *sqlStore) Existe(id int) bool {
	var exists bool
	//var id int
	query := "SELECT id FROM pacientes WHERE id=?"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {return false}
	if id >0{
		exists = true
	}
	return exists

}

// Existe DNI 
func (s *sqlStore) ExisteDNI(dni string) bool {
	/*var exists bool
	var id int
	query := "SELECT id FROM pacientes WHERE dni=?"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {return false}
	if id >0{
		 exists = true
	}
	return exists*/
	var id int
	query := "SELECT id FROM pacientes WHERE dni=?"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false // DNI no existe
		}
		// Manejo de error adicional si es necesario
		return false
	}
	return id > 0 // DNI existe

}

// BuscarPorDNI 
func (s *sqlStore) BuscarPorDNI(dni string) (domain.Paciente, error) {
	var paciente domain.Paciente
	var fechaAltaStr string
	query := "SELECT * FROM pacientes WHERE dni=?"
	row := s.db.QueryRow(query, dni) //ejecuta query y espera que retorne una sola fila de resultados, dni se pasa como segundo argumento para el ?, el resultado se almacena en row
	//asigna los valores de la fila resultante a las variables de la estructura paciente
	//oma punteros a los campos de paciente y los rellena con los valores obtenidos de la fila de la base de datos.
	err := row.Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Direccion, &paciente.DNI, &fechaAltaStr)
	//Si ocurre un error durante la operación de escaneo, como que no se encontró el paciente, se retorna un paciente vacío junto con el error correspondiente.
	//Esto permite que el código que llama a esta función maneje el error adecuadamente (mostrar "paciente no encontrado").
	if err != nil {
		return domain.Paciente{}, err
	}

	// Convertir fechaAltaStr a time.Time
	paciente.FechaAlta, err = time.Parse("2006-01-02", fechaAltaStr)
	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil

}

// Agregar
func (s *sqlStore) Agregar(paciente domain.Paciente) (int, error) {
	// Asignar la fecha actual a paciente.FechaAlta
	paciente.FechaAlta = time.Now()
	query := "INSERT INTO pacientes (nombre, apellido, direccion, dni, fecha_alta) VALUES (?,?,?,?,?)"
	stmt, err := s.db.Prepare(query) //preparo la consulta
	if err != nil {
		return 0,err //si falla la preparar la consulta devuelvo el error
	} //ejecutamos la consulta preparada y le poasamos los valores especificos, devolvemos el resultado (res)
	
	defer stmt.Close()

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Direccion, paciente.DNI, paciente.FechaAlta.Format("2006-01-02"))
	if err != nil {
		return 0,err
	} //Si hay un error al ejecutar la consulta, se devuelve ese error.
	
	

	_, err = res.RowsAffected() //devolvemos el número de filas afectadas sin utilizar ese valor, es para verificar si la ejecución fue exitosa
	if err != nil {
		return 0,err
	} //Si hay un error al intentar obtener el número de filas afectadas, se devuelve ese error.
	
	//Para capturar el id y al agregar el paciente devuelva el id correcto y ni 0
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}	

	return int(id), nil

}

// BuscarPorID
func (s *sqlStore) BuscarPorID(id int) (domain.Paciente, error) {

	var paciente domain.Paciente
	var fechaAltaStr string
	query := "SELECT * FROM pacientes WHERE id=?"
	row := s.db.QueryRow(query, id) //ejecuta query y espera que retorne una sola fila de resultados, id se pasa como segundo argumento para el ?, el resultado se almacena en row
	//asigna los valores de la fila resultante a las variables de la estructura paciente
	//toma punteros a los campos de paciente y los rellena con los valores obtenidos de la fila de la base de datos.
	err := row.Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Direccion, &paciente.DNI, &fechaAltaStr)
	//Si ocurre un error durante la operación de escaneo, como que no se encontró el paciente, se retorna un paciente vacío junto con el error correspondiente.
	//Esto permite que el código que llama a esta función maneje el error adecuadamente (mostrar "paciente no encontrado").
	if err != nil {
		return domain.Paciente{}, err
	}

	// Convertir fechaAltaStr a time.Time
	paciente.FechaAlta, err = time.Parse("2006-01-02", fechaAltaStr)
	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Eliminar //ver si eliminiamos por id o dni
func (s *sqlStore) Eliminar(id int) error {
	query := "DELETE FROM pacientes WHERE id = ?"
	stmt, err := s.db.Prepare(query) //preparo la consulta
	if err != nil {
		return err //si falla la preparar la consulta devuelvo el error
	}
	res, err := stmt.Exec(id) //ejecutamos la consulta preparada y le poasamos los valores especificos, devolvemos el resultado (res)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected() //devolvemos el número de filas afectadas sin utilizar ese valor, es para verificar si la ejecución fue exitosa
	if err != nil {
		return err
	}
	return nil
}

// Listar
func (s *sqlStore) Listar() ([]domain.Paciente, error) {
	listReturn := []domain.Paciente{}
	query := "SELECT * FROM pacientes"
	// s.db.Query devuelvo un conjunto de filas
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Verifico si hubo algún error con el conjunto de filas devuelto
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// Itero sobre cada fila devuelta por la consulta.
	for rows.Next() {
		var paciente domain.Paciente
		var fechaAltaStr string
		// Con rows.Scan copio los valores de la fila actual en los campos de paciente.

		err := rows.Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Direccion, &paciente.DNI, &fechaAltaStr)
		if err != nil {
			log.Fatal(err)
		}

		// Convertir el string al formato time.Time
        paciente.FechaAlta, err = time.Parse("2006-01-02", fechaAltaStr)
        if err != nil { 
            return nil, err
        }
		listReturn = append(listReturn, paciente) // Añado el paciente a la lista listReturn
	}

	return listReturn, nil
}

// Modificar //ver si permitimos modificar dni y fecha de alta
func (s *sqlStore) Modificar(paciente domain.Paciente) error {
	query := "UPDATE pacientes SET nombre = ?, apellido= ?, direccion=?, dni=?, fecha_alta=? WHERE id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Direccion, paciente.DNI, paciente.FechaAlta.Format("2006-01-02"), paciente.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}


