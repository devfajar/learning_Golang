package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type student struct {
	id		string
	name	string
	age		int
	grade	int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id, name, grade FROM tb_student WHERE age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = "E001"
	err = db.QueryRow("SELECT name, grade FROM tb_student where id = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name %s\n grade: %d\n", result.name, result.grade)
}


// Menggunakan Prepare
func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()


	stmt, err := db.Prepare("SELECT name, grade FROM tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\n grade: %d\n", result1.name, result1.grade)

	var result2 = student{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\n grade: %d\n", result2.name, result2.grade)

	var result3 = student{}
	stmt.QueryRow("B001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\n grade: %d\n", result3.name, result3.grade)
}

// Menggunakan exec
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tb_student VALUES (?, ?, ?, ?)", "G001", "Galahad", 29, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Insert Success")

	_, err = db.Exec("UPDATE tb_student set age = ? WHERE id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Update Success")


	_, err = db.Exec("DELETE FROM tb_student WHERE id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Delete Success")
}


func main() {
	// Menggunakan Query
	//sqlQuery()

	// Menggunakan QueryRow
	//sqlQueryRow()

	// Menggunakan Prepare
	//sqlPrepare()

	// Menggunakan Exec
	sqlExec()
}