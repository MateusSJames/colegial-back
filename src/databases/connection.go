package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
    
    connStr := "postgres://sa:12345@localhost:5432/colegial?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    // Verificar a conexão com o banco de dados
    if err := db.Ping(); err != nil {
        return nil, err
    }

    fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

    return db, nil
}