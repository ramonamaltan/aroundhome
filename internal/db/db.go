package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "pguser"
	password = "localtest"
	dbname   = "aroundhome"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer DB.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	err = insertDummyData(db)
	if err != nil {
		panic(err)
	}
	return db
}

func insertDummyData(db *sql.DB) error {
	insertDynStmt := `insert into "partners"("servicename", "latitude", "longitude", "radius") values($1, $2, $3, $4)`
	_, err := db.Exec(insertDynStmt, "flooring", 59.12345, 59.12345, 50)
	if err != nil {
		panic(err)
	}
	return nil
}
