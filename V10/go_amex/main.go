package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type ToDo struct {
	ID          int    `json:"id"`
	TASK_STRING string `json:"task_string"`
	DONE        bool   `json:"done"`
}

type Profile struct {
	Name    string
	Hobbies []string
}

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db_amex01"
)

var db *gorm.DB
var err error

var SQLBind string

func main() {

	SQLBind = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", SQLBind)

	if err != nil {
		fmt.Println("============   exiting ==========")
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	// err = db.Ping()
	if err != nil {
		fmt.Println("============   not pinging ==========")
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("======================")
	fmt.Println("Successfully connected!")

	// Migrate the schema
	db.AutoMigrate(&ToDo{})

	db.Create(&ToDo{ID: 1, TASK_STRING: "This is the first task just to setup", DONE: false})
	db.Create(&ToDo{ID: 2, TASK_STRING: "Second task", DONE: false})
	db.Create(&ToDo{ID: 3, TASK_STRING: "Third Task ", DONE: false})
	db.Create(&ToDo{ID: 4, TASK_STRING: "Fourth Task ", DONE: false})

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/double", doubleHandler)
	r.HandleFunc("/api01/initialize", InitializeHandler).Methods("GET")
	r.HandleFunc("/api01/createtodo/{id}/{todo_text}", CreateTodoHandler).Methods("GET")
	r.HandleFunc("/api01/gettodo/{id}", GetTodoHandler).Methods("GET")
	r.HandleFunc("/api01/deletetodo/{id}", DeleteTodoHandler).Methods("GET")
	r.HandleFunc("/api01/closetodo/{id}", CloseTodoHandler).Methods("GET")
	r.HandleFunc("/api01/Updatetodo/{id}/{todo_text}", UpdateTodoHandler).Methods("GET")
	r.HandleFunc("/api01/listalltodo", ListAllTodoHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/double", doubleHandler)
	return r
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("v")
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, v*2)
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// func foo(w http.ResponseWriter, r *http.Request) {
// 	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

// 	js, err := json.Marshal(profile)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// func foo(w http.ResponseWriter, r *http.Request) {
	profile := ToDo{999, "This is the root", false}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func InitializeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " setting up new todo DB..  or ressetting it....")
	w.Header().Set("Content-Type", "application/json")
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Create a new todo....")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	tt := vars["todo_text"]
	id, _ := strconv.Atoi(vars["id"])
	fmt.Fprintf(w, "xx   CREATE TASK .. %d  %s", id, tt)

	db, err := gorm.Open("postgres", SQLBind)

	fmt.Println("----------")
	fmt.Println(SQLBind)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(&ToDo{ID: id, TASK_STRING: tt, DONE: false})
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, " Get Todo.. %s", id)

	db, err := gorm.Open("postgres", SQLBind)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var todo []ToDo
	db.Where("id = ?", id).First(&todo)
	ct := strconv.Itoa(len(todo))
	fmt.Println("---------", ct)

	json.NewEncoder(w).Encode(todo)
	fmt.Println("{}", todo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " delete a new todo....")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, " Delete Todo.. %s", id)

	db, err := gorm.Open("postgres", SQLBind)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var todo []ToDo
	db.Where("id = ?", id).Delete(&todo)
	fmt.Println("{}", todo)
	json.NewEncoder(w).Encode(todo)
	fmt.Println("{}", todo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " update todo....")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	tt := vars["todo_text"]
	id, _ := strconv.Atoi(vars["id"])
	fmt.Fprintf(w, "xx   UPDATE TASK .. %d  %s", id, tt)
	db, err := gorm.Open("postgres", SQLBind)

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	td := &ToDo{ID: id}
	db.Find(&td)
	td.TASK_STRING = tt
	db.Save(&td)
}

func CloseTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " update todo....")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	tt := vars["todo_text"]
	id, _ := strconv.Atoi(vars["id"])
	fmt.Fprintf(w, "xx   UPDATE TASK .. %d  %s", id, tt)

	db, err := gorm.Open("postgres", SQLBind)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	td := &ToDo{ID: id}
	db.Find(&td)

	td.DONE = true
	db.Save(&td)

}

func ListAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " List All todo's....")
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", SQLBind)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var todo []ToDo
	db.Find(&todo)
	fmt.Println("{}", todo)

	json.NewEncoder(w).Encode(todo)
	fmt.Println("{}", todo)
}
