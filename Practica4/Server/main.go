package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	// "github.com/gomzalo/db"
	"github.com/gorilla/mux"
)

// type ESTUDIANTE struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

type Reporte struct {
	Carnet int    `json:"Carnet"`
	Nombre string `json:"Nombre"`
	Curso  string `json:"Curso"`
	Fecha  string `json:"Fecha"`
	Cuerpo string `json:"Cuerpo"`
}

var reportes []Reporte

var nombre_servidor = "default"
var mensaje = "Mensaje default"

var db_handler *sql.DB

//PruebaHandler funcion~
func PruebaHandler(w http.ResponseWriter, request *http.Request) {
	enableCors(&w)

	w.Write([]byte("AY MI MADRE EL BICHO SIIIIUUUUUUUU"))
}

//MensajeHandler funcion
func MensajeHandler(w http.ResponseWriter, request *http.Request) {
	enableCors(&w)

	w.Write([]byte(mensaje))
}

func DatosHandler(w http.ResponseWriter, request *http.Request) {
	enableCors(&w)
	log.Println("DatosHandler")
	results, err := db_handler.Query("SELECT Carnet, Nombre FROM Reporte")
	// results, err := db_handler.Query("SELECT EmployeeId, FirstName FROM EMPLOYEE")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	registros := ""
	log.Println("Resultados: ", results)
	for results.Next() {
		// var emp ESTUDIANTE
		// // for each row, scan the result into ESTUDIANTE object
		// err = results.Scan(&emp.Id, &emp.Name)
		// if err != nil {
		// 	panic(err.Error()) // proper error handling instead of panic in your app
		// }
		// registros = registros + strconv.Itoa(emp.Id) + " - " + emp.Name + "\n"
		var emp Reporte
		// for each row, scan the result into Reporte object
		err = results.Scan(&emp.Carnet, &emp.Nombre)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		registros = registros + strconv.Itoa(emp.Carnet) + " - " + emp.Nombre + "\n"
	}
	registros = registros + mensaje

	log.Println("Registros: ", registros)
	w.Write([]byte(registros))
}

// ::::::::::::::::::: 	ENDPOINTS PRACTICA	 :::::::::::::::::::
func CrearReporteHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "appliction/json")
	var r Reporte
	json.NewDecoder(request.Body).Decode(&r)
	r.Carnet = len(reportes) + 1
	reportes = append(reportes, r)
	json.NewEncoder(w).Encode(r)
	enableCors(&w)
	log.Println("DatosHandler")
	query := "INSERT INTO Reporte (Carnet, Nombre, Curso, Cuerpo) VALUES (?, ?, ?, ?, ?);"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	// results, err := db_handler.ExecContext("SELECT Carnet, Nombre FROM Reporte")
	stmt, err := db_handler.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		// return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, r.Carnet, r.Nombre, r.Curso, r.Fecha, r.Cuerpo)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		// return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		// return err
	}
	log.Printf("%d products created ", rows)
	prdID, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when getting last inserted product", err)
		// return err
	}
	log.Printf("Product with ID %d created", prdID)
	// return nil
	w.Write([]byte("registros"))
}
func main() {
	router := mux.NewRouter()
	// Endpoints
	router.HandleFunc("/siu", PruebaHandler).Methods("GET")
	router.HandleFunc("/mensaje", MensajeHandler).Methods("GET")
	router.HandleFunc("/dato", DatosHandler).Methods("GET")
	router.HandleFunc("/crear_reporte", CrearReporteHandler).Methods("POST")
	// BD
	nombre_servidor = os.Getenv("ID_SERVIDOR")
	mensaje = "Hola! te saluda el servidor " + nombre_servidor
	log.Println("Estoy escuchando en el puerto 5000")

	//Se utilizan variables de entorno para poder modificar valores desde el archivo docker-compose, y
	//asi evitar tener que modificar el codigo, y generar nuevamente la imagen.
	usuario := os.Getenv("USER_NAME")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	db_name := os.Getenv("DB_NAME")
	// database, err := db.Initialize(usuario, pass, db_name)

	// if err != nil {
	// 	log.Fatalf("No se pudo conectar a la BD: %v", err)
	// 	panic(err.Error())
	// } else {
	// 	db_handler = database.Conn
	// 	log.Println("aaaaaa")
	// }
	// defer database.Conn.Close()
	conn_string := usuario + ":" + pass + "@tcp(" + host + ":" + port + ")/" + db_name
	db, err := sql.Open("mysql", conn_string)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		db_handler = db
		log.Println("conectado a la base de datos")
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	log.Fatal(http.ListenAndServe(":5000", router))

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
