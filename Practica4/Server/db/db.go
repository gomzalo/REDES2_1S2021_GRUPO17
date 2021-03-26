package db

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lib/pq"
)

const (
	HOST = "db"
	PORT = 3306
)

type Database struct {
	Conn *sql.DB
}

func Initialize(usuario, pass, db_name string) (Database, error) {
	db := Database{}
	// conn_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	HOST, PORT, usuario, pass, database)
	//Se utilizan variables de entorno para poder modificar valores desde el archivo docker-compose, y
	//asi evitar tener que modificar el codigo, y generar nuevamente la imagen.
	// usuario := os.Getenv("USER_NAME")
	// pass := os.Getenv("PASS")
	// host := os.Getenv("HOST")
	// port := os.Getenv("PORT")
	// db_name := os.Getenv("DB_NAME")
	conn_string := usuario + ":" + pass + "@tcp(" + HOST + ":" + strconv.Itoa(PORT) + ")/" + db_name
	conn, err := sql.Open("mysql", conn_string)
	if err != nil {
		// panic(err.Error())
		return db, err
	} else {
		db.Conn = conn
		log.Println("conectado a la base de datos")
	}

	// if there is an error opening the connection, handle it
	// db.Conn = conn
	// err = db.Conn.Ping()
	// if err != nil {
	// 	// panic(err.Error())
	// 	return db, err
	// }
	//  else {
	// 	db.Conn = conn
	// log.Println("conectado a la base de datos")
	// }
	return db, nil
}
