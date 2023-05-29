package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// the db server information
	config := mysql.Config{
		User:   "root",
		Passwd: "Joie_Vibre@0!9",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "students",
	}
	// get database handle
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("CONNECTED TO DB")

	// Call to db operation functions
	// Call to listAll
	stds, err := listAll()
	if err != nil {
		fmt.Println("some error")
	} else {
		fmt.Println(stds)
	}

	// Call to addStudent
	upStudent := Student{
		Name:  "Ron Ron",
		Email: "ron@gmail.com",
		Age:   21,
	}

	lastId, err := addStudent(upStudent)
	if err == nil {
		fmt.Println("Last insert id:", lastId)
	}

	stds, err = listAll()
	if err != nil {
		fmt.Println("some error")
	} else {
		fmt.Println(stds)
	}

	// Call to updateStudent
	newStudent := Student{
		Name:  "Ron Ron",
		Email: "ron@gmail.com",
		Age:   21,
	}
	lastId2, err := updateStudent(1, newStudent)
	if err == nil {
		fmt.Println("Last update id:", lastId2)
	}

	stds, err = listAll()
	if err != nil {
		fmt.Println("some error")
	} else {
		fmt.Println(stds)
	}

	// Call to deleteStudent

}

type Student struct {
	ID    int64
	Name  string
	Email string
	Age   int
}

func listAll() ([]Student, error) {
	var students []Student
	rows, err := db.Query("select * from students")
	if err != nil {
		return nil, fmt.Errorf("error in query all student: %v", err)
	}
	defer rows.Close()
	// loop through rows using scans to assign record to slice
	for rows.Next() {
		var std Student

		if err := rows.Scan(&std.ID, &std.Name, &std.Email, &std.Age); err != nil {
			return nil, fmt.Errorf("error in query all students: %v", err)
		}
		students = append(students, std)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in query all student: %v", err)
	}
	return students, nil

}

func addStudent(std Student) (int64, error) {
	result, err := db.Exec("INSERT into student (name,email,age) values (?,?,?)", std.Name, std.Email, std.Age)
	if err != nil {
		return 0, fmt.Errorf("Error in add student %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error in add student %v", err)
	}
	return id, nil
}

func updateStudent(stdId int, std Student) (int64, error) {
	result, err := db.Exec("UPDATE student SET name =?,  email= ?,  age= ? WHERE id=?", std.Name, std.Email, std.Age, stdId)
	if err != nil {
		return 0, fmt.Errorf("update student: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("update student: %v", err)
	}
	return id, nil
}
