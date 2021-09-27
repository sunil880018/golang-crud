package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"todoApp/Conn"
	"todoApp/model"
)
func Getdata(w http.ResponseWriter, r *http.Request) {
	db := Conn.OpenConnection()
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}
	var people []model.Person
	for rows.Next() {
		var person model.Person
		rows.Scan(&person.Id ,&person.Name, &person.Course)
		people = append(people, person)
	}
	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}
func Create(w http.ResponseWriter, r *http.Request) {
	db := Conn.OpenConnection()
    w.Header().Set("Content-Type","application/json")
	var p model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	fmt.Println(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id:=p.Id
	sqlStatement := `INSERT INTO person (id,name,course) VALUES ($1, $2, $3)`
	fmt.Println(sqlStatement)
	_, err = db.Exec(sqlStatement,id, p.Name, p.Course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
func delete(w http.ResponseWriter, r *http.Request) {
	db := Conn.OpenConnection()
	w.Header().Set("Content-Type","application/json")
	var p model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := p.Id
	fmt.Println("my id ",id)
	sqlStatement := `DELETE FROM person WHERE id = $1;`
	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
func update(w http.ResponseWriter, r *http.Request) {
	db := Conn.OpenConnection()
	w.Header().Set("Content-Type","application/json")
	var p model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    id:= p.Id
	sqlStatement := `UPDATE person SET name = $2, course = $3 WHERE id = $1;`
	_, err = db.Exec(sqlStatement, id,p.Name,p.Course)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", Getdata).Methods("GET")
	r.HandleFunc("/insert", Create).Methods("POST")
	r.HandleFunc("/delete",delete).Methods("DELETE")
	r.HandleFunc("/update",update).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
