package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func insert(conn *sql.DB, command string) error {
	res, err := conn.Query(command) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}

func sqlSelect(conn *sql.DB, command string) error {
	res, err := conn.Query(command)
	if err != nil {
		return err
	}
	defer res.Close()

	for res.Next() {
		var person Person
		err = res.Scan(&person.Name, &person.Age)
		if err != nil {
			return err
		}
		fmt.Println(person.Name, person.Age)
	}
	return nil
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	conn, err := sql.Open("mysql", "root:951219@tcp(localhost:3306)/db01")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// insert
	// err = insert(conn, "INSERT INTO person(name, age) VALUES ('Ting', 25)")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// select
	err = sqlSelect(conn, "SELECT * FROM person")
	if err != nil {
		fmt.Println(err)
	}

	// select one
	var person Person
	err = conn.QueryRow("SELECT name, age FROM person where name = ?", "Chen").Scan(&person.Name, &person.Age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person.Age, person.Name)
}
