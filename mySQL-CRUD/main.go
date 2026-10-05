package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type Student struct {
	RollNo      int    `json:"rollNo"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

var db *sql.DB

func connectDB() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "student"

	fmt.Println(os.Getenv("DBUSER"))
	fmt.Println(os.Getenv("DBPASSWORD"))

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database Connected")
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file | ", err)
	}
}

func create(student Student) error {
	result, err := db.Exec("INSERT INTO class5 (FirstName, LastName, Gender, Age, PhoneNumber) VALUES (?,?,?,?,?)",
		student.FirstName,
		student.LastName,
		student.Gender,
		student.Age,
		student.PhoneNumber,
	)
	if err != nil {
		// panic(err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		// panic(err.Error())
		return err
	}

	if rows == 0 {
		// panic("Error while creating new row.")
		return errors.New("Error while creating student.")
	}

	fmt.Println("Data inserted successfully.")
	return nil
}

func readOne(rollNo int) (Student, error) {
	var student Student
	err := db.QueryRow("SELECT * FROM class5 where RollNo = ?", rollNo).Scan(&student.RollNo, &student.FirstName, &student.LastName, &student.Gender, &student.Age, &student.PhoneNumber)
	if err != nil {
		// panic(err.Error())
		return student, err
	}

	fmt.Println(student)
	return student, nil
}

func readAll() ([]Student, error) {
	var Students []Student
	results, err := db.Query("SELECT * FROM class5")
	if err != nil {
		// panic(err.Error())
		return Students, err
	}

	// result.Err()

	for results.Next() {
		var student Student

		err := results.Scan(&student.RollNo, &student.FirstName, &student.LastName, &student.Gender, &student.Age, &student.PhoneNumber)
		if err != nil {
			panic(err.Error())
		}

		Students = append(Students, student)
	}

	fmt.Println(Students)
	return Students, nil
}

// func update(rollNo int, firstName string, lastName string, gender string, age int, phoneNumber string) {
// 	result, err := db.Exec("UPDATE class5 set firstName = ?, lastName = ?, gender = ?, age = ?, phoneNumber = ? where rollNo = ?", firstName, lastName, gender, age, phoneNumber, rollNo)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	if rows == 0 {
// 		panic("Error while updating.")
// 	}

// 	fmt.Println("Data update successfully.")
// }

func update(rollNo int, student Student) error {
	result, err := db.Exec("UPDATE class5 set firstName = ?, lastName = ?, gender = ?, age = ?, phoneNumber = ? where rollNo = ?", student.FirstName, student.LastName, student.Gender, student.Age, student.PhoneNumber, rollNo)
	if err != nil {
		// panic(err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		// panic(err.Error())
		return err
	}

	if rows == 0 {
		// panic("Error while updating.")
		return errors.New("Please update some date.")
	}

	fmt.Println("Data update successfully.")
	return nil
}

func delete(rollNo int) error {
	result, err := db.Exec("DELETE FROM class5 WHERE RollNo = ?", rollNo)
	if err != nil {
		// panic(err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		// panic(err.Error())
		return err
	}

	if rows == 0 {
		// panic("Error while deleting data.")
		return errors.New("Invalid ID or student doesn't exist.")
	}

	fmt.Println("Data deleted successfully.")
	return nil
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	students, err := readAll()
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := struct {
		Data []Student `json:"data"`
	}{Data: students}

	json.NewEncoder(w).Encode(response)
}

func getOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	rollNo := vars["rollNo"]
	rollNo_INT, err := strconv.Atoi(rollNo)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	student, err := readOne(rollNo_INT)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: "Invalid ID"}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	response := struct {
		Data Student `json:"data"`
	}{Data: student}

	json.NewEncoder(w).Encode(response)

}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	rollNo := vars["rollNo"]
	rollNo_INT, err := strconv.Atoi(rollNo)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	var student Student

	err = json.Unmarshal(reqBody, &student)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: "Invalid JSON/BODY"}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	students, err := readAll()
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	for _, stu := range students {
		if stu.RollNo == rollNo_INT {
			if student.FirstName == "" {
				student.FirstName = stu.FirstName
			}
			if student.LastName == "" {
				student.LastName = stu.LastName
			}
			if student.Gender == "" {
				student.Gender = stu.Gender
			}
			if student.Age == 0 {
				student.Age = stu.Age
			}
			if student.PhoneNumber == "" {
				student.PhoneNumber = stu.PhoneNumber
			}
			if student.RollNo == 0 {
				student.RollNo = stu.RollNo
			}
		}
	}

	err = update(rollNo_INT, student)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	response := struct {
		Message string  `json:"message"`
		Data    Student `json:"data"`
	}{Message: "Student updated successfully", Data: student}

	json.NewEncoder(w).Encode(response)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create student route hit.")
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	var student Student
	err = json.Unmarshal(reqBody, &student)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	if student.FirstName == "" {
		err := errors.New("First Name is required.")
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	} else if student.LastName == "" {
		err := errors.New("Last Name is required.")
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	} else if student.Gender == "" {
		err := errors.New("Gender is required.")
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	} else if student.Age == 0 {
		err := errors.New("Age is required.")
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	} else if student.PhoneNumber == "" {
		err := errors.New("Phone number is required.")
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	err = create(student)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}
			fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{Message: "Student created successfully"}

	json.NewEncoder(w).Encode(response)

}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	rollNo := vars["rollNo"]
	rollNo_INT, err := strconv.Atoi(rollNo)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	err = delete(rollNo_INT)
	if err != nil {
		response := struct {
			Error string `json:"error"`
		}{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{Message: "Student data deleted successfully."}

	json.NewEncoder(w).Encode(response)
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/v1/students", getAllStudents).Methods("GET")
	myRouter.HandleFunc("/api/v1/student/{rollNo}", getOneStudent).Methods("GET")
	myRouter.HandleFunc("/api/v1/student/{rollNo}", updateStudent).Methods("PUT")
	myRouter.HandleFunc("/api/v1/student", createStudent).Methods("POST")
	myRouter.HandleFunc("/api/v1/student/{rollNo}", deleteStudent).Methods("DELETE")

	// CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"}, // frontend URLs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	fmt.Println("Server is listening on port 3001")

	// Wrap your router with the CORS handler
	log.Fatal(http.ListenAndServe(":3001", corsHandler.Handler(myRouter)))

	// fmt.Println("Server is listening on port 3000")
	// log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	loadENV()

	connectDB()
	defer db.Close()

	// create("Arya", "Stark", "Female", 18, "8639368924")
	// readOne(1)
	// readAll()
	// update(2, "Arya", "Stark", "Female", 20, "8639368924")
	// delete(1)
	// readAll()

	handleRequest()
}
