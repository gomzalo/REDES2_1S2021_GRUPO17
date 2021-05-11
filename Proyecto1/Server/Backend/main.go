package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	datapb "grpctypes"

	// "github.com/gomzalo/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// type ESTUDIANTE struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }
type server struct{}

type Reporte struct {
	Carnet int    `json:"Carnet"`
	Nombre string `json:"Nombre"`
	Curso  string `json:"Curso"`
	Fecha  string `json:"Fecha"`
	Cuerpo string `json:"Cuerpo"`
}

type CarnetE struct {
	Carnet int `json:"Carnet"`
}

var carnets []CarnetE

// var reportes []Reporte

type reporteJSON struct {
	reportes []Reporte
	string   `json:"data"`
}

var nombre_servidor = "default"
var mensaje = "Mensaje default"

var db_handler *sql.DB
var db *mongo.Client

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

	results, err := db_handler.Query("SELECT Carnet, Nombre, Curso, Cuerpo FROM Reporte")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	registros := ""
	var reportes []Reporte
	var msj Reporte
	msj.Cuerpo = mensaje
	msj.Nombre = nombre_servidor
	reportes = append(reportes, msj)

	log.Println("Resultados: ", results)
	for results.Next() {
		var emp Reporte
		err = results.Scan(&emp.Carnet, &emp.Nombre, &emp.Curso, &emp.Cuerpo)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// registros = registros + strconv.Itoa(emp.Carnet) + " - " + emp.Nombre + " - " + emp.Curso + " - " + emp.Fecha + " - " + emp.Cuerpo + "\n"
		reportes = append(reportes, emp)
		json.NewEncoder(log.Writer()).Encode(emp)
		registros = registros + strconv.Itoa(emp.Carnet) + " - " + emp.Nombre + " - " + emp.Curso + " - " + emp.Cuerpo + "\n"
	}

	// json.NewEncoder(log.Writer()).Encode(reportes)
	json.NewEncoder(w).Encode(reportes)
	// registros = registros + mensaje

	log.Println("Registros: ", reportes)
	// w.Write([]byte("[{ \"mensaje\": \"" + mensaje + "\"}]"))
}

// ::::::::::::::::::: 	ENDPOINTS PRACTICA	 :::::::::::::::::::
func CrearReporteHandler(w http.ResponseWriter, request *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// (w).Header().Set("Access-Control-Allow-Origin", "*")
	// (w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	// (w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "appliction/json")
	// var reportes []Reporte
	carnet_, ok := request.URL.Query()["Carnet"]
	curso_, ok := request.URL.Query()["Curso"]
	nombre_, ok := request.URL.Query()["Nombre"]
	cuerpo_, ok := request.URL.Query()["Cuerpo"]
	fecha_, ok := request.URL.Query()["Fecha"]
	if !ok || len(carnet_[0]) < 1 {
		log.Println("Url Param 'Carnet' is missing")
		return
	}
	Carnet := carnet_[0]
	Curso := curso_[0]
	Nombre := nombre_[0]
	Cuerpo := cuerpo_[0]
	Fecha := fecha_[0]
	log.Println("Carnet: ", Carnet)
	log.Println("Curso: ", Curso)
	log.Println("Nombre: ", Nombre)
	log.Println("Cuerpo: ", Cuerpo)
	log.Println("Fecha: ", Fecha)
	// var nuevoReporte Reporte
	// json.NewDecoder(request.Body).Decode(&nuevoReporte)
	// // nuevoReporte.Carnet = (len(reportes) + 1)
	// reportes = append(reportes, nuevoReporte)
	// json.NewEncoder(w).Encode(nuevoReporte)
	// // Obteniendo datos de reporte
	// log.Println("Reportes: ", reportes)
	// Ingresando a BD

	// ******************************************
	// query := "INSERT INTO Reporte(Carnet, Nombre, Curso, Fecha, Cuerpo) VALUES (?, ?, ?, ?, ?);"
	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancelfunc()
	// stmt, err := db_handler.PrepareContext(ctx, query)
	// if err != nil {
	// 	log.Printf("Error %s when preparing SQL statement", err)
	// 	// return err
	// }
	// defer stmt.Close()
	// // current_time := time.Now().Local()
	// // current_time = current_time.Format("01-02-2006 15:04:05")
	// // nuevoReporte.Fecha = current_time.String()
	// res, err := stmt.ExecContext(ctx, Carnet, Nombre, Curso, Fecha, Cuerpo)
	// if err != nil {
	// 	log.Printf("Error %s when inserting row into products table", err)
	// 	// return err
	// }
	// rows, err := res.RowsAffected()
	// if err != nil {
	// 	log.Printf("Error %s when finding rows affected", err)
	// 	// return err
	// }
	// log.Printf("%d Reporte creado ", rows)
	// prdID, err := res.LastInsertId()
	// if err != nil {
	// 	log.Printf("Error %s when getting last inserted product", err)
	// 	// return err
	// }
	// log.Printf("Reporte del carnte %d creado", prdID)
	// // return nil
	// ******************************************
	results, err := db_handler.Query("INSERT INTO Reporte(Carnet, Nombre, Curso, Fecha, Cuerpo) VALUES (?, ?, ?, ?, ?);", Carnet, Nombre, Curso, Fecha, Cuerpo)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	registros := ""
	var reportes []Reporte
	var msj Reporte
	msj.Cuerpo = mensaje
	msj.Nombre = nombre_servidor
	reportes = append(reportes, msj)
	defer results.Close()
	log.Println("Resultados: ", results)
	for results.Next() {
		var emp Reporte
		err = results.Scan(&emp.Carnet, &emp.Nombre, &emp.Curso, &emp.Fecha, &emp.Cuerpo)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// registros = registros + strconv.Itoa(emp.Carnet) + " - " + emp.Nombre + " - " + emp.Curso + " - " + emp.Fecha + " - " + emp.Cuerpo + "\n"
		reportes = append(reportes, emp)
		json.NewEncoder(log.Writer()).Encode(emp)
		registros = registros + strconv.Itoa(emp.Carnet) + " - " + emp.Nombre + " - " + emp.Curso + " - " + emp.Cuerpo + "\n"
	}

	// json.NewEncoder(log.Writer()).Encode(reportes)
	json.NewEncoder(w).Encode(reportes)
	registros = registros + mensaje

	log.Println("Registros: ", reportes)

	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(`{"message": "Reporte ingresado"}`))
}

func BuscarReporteHandler(w http.ResponseWriter, request *http.Request) {
	////////////////////////////////////////////////////////////////////////////
	// enableCors(&w)
	// w.Header().Set("Content-Type", "appliction/json")
	// // var nuevoCarnet CarnetE
	// // json.NewDecoder(request.Body).Decode(&nuevoCarnet)
	// // // nuevoReporte.Carnet = (len(reportes) + 1)
	// // carnets = append(carnets, nuevoCarnet)
	// // json.NewEncoder(w).Encode(nuevoCarnet)
	// // // Obteniendo datos de reporte
	// // log.Println("Carnets: ", carnets)
	// params := mux.Vars(request)
	// carnet := params["carnet"]
	// fmt.Println(carnet)
	// var nuevoReporte Reporte
	// query := "SELECT * FROM Reporte WHERE Carnet = ?;"
	// row, err := db_handler.Query(query, carnet)
	// // *****************
	// if err != nil {
	// 	// panic(err.Error()) // proper error handling instead of panic in your app
	// 	log.Fatal(err)
	// }
	// defer row.Close()
	// for row.Next() {
	// 	row.Scan(&nuevoReporte.Carnet, &nuevoReporte.Nombre, &nuevoReporte.Curso, &nuevoReporte.Fecha, &nuevoReporte.Cuerpo)
	// 	json.NewEncoder(w).Encode(nuevoReporte)
	// }
	// w.WriteHeader(http.StatusCreated)
	// // nombre_servidor = os.Getenv("ID_SERVIDOR")
	// mensaje = "Solicitud atendida por el servidor: " + nombre_servidor
	// w.Write([]byte(`{"message":` + `"` + mensaje + `"}`))
	////////////////////////////////////////////////////////////////////////////
	enableCors(&w)
	w.Header().Set("Content-Type", "appliction/json")
	// log.Println("DatosHandler")
	params := mux.Vars(request)
	carnet := params["carnet"]
	fmt.Println("Carnet a buscar: ", carnet)

	results, err := db_handler.Query("SELECT * FROM Reporte WHERE Carnet = ?;", carnet)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	registros := ""
	var reportes []Reporte
	var msj Reporte
	msj.Cuerpo = mensaje
	msj.Nombre = nombre_servidor
	reportes = append(reportes, msj)
	defer results.Close()
	log.Println("Resultados: ", results)
	for results.Next() {
		var emp Reporte
		err = results.Scan(&emp.Carnet, &emp.Nombre, &emp.Curso, &emp.Fecha, &emp.Cuerpo)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// registros = registros + strconv.Itoa(emp.Carnet) + " - " + emp.Nombre + " - " + emp.Curso + " - " + emp.Fecha + " - " + emp.Cuerpo + "\n"
		reportes = append(reportes, emp)
		json.NewEncoder(log.Writer()).Encode(emp)
		registros = registros + strconv.Itoa(emp.Carnet) + " - " + emp.Nombre + " - " + emp.Curso + " - " + emp.Cuerpo + "\n"
	}

	// json.NewEncoder(log.Writer()).Encode(reportes)
	json.NewEncoder(w).Encode(reportes)
	registros = registros + mensaje

	log.Println("Registros: ", reportes)
	// w.Write([]byte(registros))
}

func MostrarReporteHandler(w http.ResponseWriter, request *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "appliction/json")
	var nuevoReporte Reporte
	json.NewDecoder(request.Body).Decode(&nuevoReporte)
	// nuevoReporte.Carnet = (len(reportes) + 1)
	var reportes []Reporte
	reportes = append(reportes, nuevoReporte)
	json.NewEncoder(w).Encode(nuevoReporte)
	// Obteniendo datos de reporte
	log.Println("Reportes: ", reportes)

	// Ingresando a BD

	query := "INSERT INTO Reporte(Carnet, Nombre, Curso, Fecha, Cuerpo) VALUES (?, ?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db_handler.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		// return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, nuevoReporte.Carnet, nuevoReporte.Nombre, nuevoReporte.Curso, nuevoReporte.Fecha, nuevoReporte.Cuerpo)
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

	// *****************
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func (*server) SaveReporte(ctx context.Context, req *datapb.Reporte) (*datapb.Response, error) {
	reportes := db.Database("redes").Collection("reportes")
	data := datapb.Reporte{}

	btes, err := json.Marshal(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	json.Unmarshal(btes, &data)

	fmt.Println(data)

	insertResult, err := reportes.InsertOne(context.TODO(), data)
	fmt.Println(insertResult)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result := &datapb.Response{
		Description: "Dato insertado",
	}

	return result, nil
}

func (*server) SaveAsistencia(ctx context.Context, req *datapb.Asistencia) (*datapb.Response, error) {

	asistencias := db.Database("redes").Collection("asistencias")
	data := datapb.Asistencia{}

	btes, err := json.Marshal(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	json.Unmarshal(btes, &data)

	fmt.Println(data)

	insertResult, err := asistencias.InsertOne(context.TODO(), data)
	fmt.Println(insertResult)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result := &datapb.Response{
		Description: "Dato insertado",
	}

	return result, nil
}

func main() {
	//Conexion a mongodb
	clientOptions := options.Client().ApplyURI("mongodb://db:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Connected to MongoDB!")
	db = client

	//Servidor GRPC
	host_grpc := ":3000"

	listener, err := net.Listen("tcp", host_grpc)

	if err != nil {
		log.Fatalf("El listener no funciono %v", err)
	}

	s := grpc.NewServer()

	datapb.RegisterDataServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("El server no funciono %v", err)
	}
	//----------------------------------

	router := mux.NewRouter()
	// Endpoints
	router.HandleFunc("/siu", PruebaHandler).Methods("GET")
	router.HandleFunc("/mensaje", MensajeHandler).Methods("GET")
	router.HandleFunc("/dato", DatosHandler).Methods("GET")
	router.HandleFunc("/crear", CrearReporteHandler).Methods("GET")
	router.HandleFunc("/buscar/{carnet}", BuscarReporteHandler).Methods("GET")
	router.HandleFunc("/mostrar", MostrarReporteHandler).Methods("GET")
	// BD
	nombre_servidor = os.Getenv("ID_SERVIDOR")
	mensaje = "Procesado por el servidor " + nombre_servidor
	log.Println("Estoy escuchando en el puerto 5000")

	//Se utilizan variables de entorno para poder modificar valores desde el archivo docker-compose, y
	//asi evitar tener que modificar el codigo, y generar nuevamente la imagen.
	usuario := os.Getenv("USER_NAME")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	db_name := os.Getenv("DB_NAME")

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
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
