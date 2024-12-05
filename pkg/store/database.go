package store

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql" // Driver para MySQL
)

//hacemos la conexión a la base de datos de esta manera para: 
//mejor organización del código, la separación de responsabilidades, y el mantenimiento a largo plazo.
//para no sobrecargar el main con más responsabilidades

// NewDatabaseConnection inicializa una conexión a la base de datos
func NewDatabaseConnection() (*sql.DB, error) {
    dsn := "root:diegoadmin1@tcp(localhost:3306)/clinica_odontologica"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Verifica que la conexión sea válida
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    log.Println("Conexión a la base de datos exitosa.")
    return db, nil
}