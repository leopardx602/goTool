package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username  string = "root"
	password  string = "password"
	addr      string = "127.0.0.1"
	port      int    = 3306
	database  string = "db01"
	parseTime bool   = true // time.time or string
)

func createTable(conn *sql.DB) error {
	sql := `CREATE TABlE table01(
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(16) NOT NULL DEFAULT "",
		price INT DEFAULT 0,
		image VARCHAR(64) DEFAULT "",
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`
	_, err := conn.Query(sql)
	if err != nil {
		return err
	}
	return nil
}

func insert(conn *sql.DB, command string) error {
	_, err := conn.Exec(command) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		return err
	}
	//fmt.Println(res)
	return nil
}

func sqlSelect(conn *sql.DB, command string) error {
	res, err := conn.Query(command)
	if err != nil {
		return err
	}
	defer res.Close()

	for res.Next() {
		var product Product
		err = res.Scan(&product.ID, &product.Name, &product.Price, &product.Image, &product.CreateAt, &product.UpdateAt)
		if err != nil {
			return err
		}
		fmt.Println(product)
	}
	return nil
}

type Product struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Price    int       `json:"age"`
	Image    string    `json:"image"`
	CreateAt string    `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func main() {
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v", username, password, addr, port, database, parseTime)
	conn, err := sql.Open("mysql", connInfo)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// create talble
	// err = createTable(conn)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// insert
	// err = insert(conn, "INSERT INTO table01(name, price, image) VALUES ('apple12', 1000, '')")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// select
	err = sqlSelect(conn, "SELECT * FROM table01")
	if err != nil {
		fmt.Println(err)
	}

	// select one
	// var product Product
	// err = conn.QueryRow("SELECT name, price FROM table01 where name = ?", "apple12").Scan(&product.Name, &product.Price)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(product.Price, product.Name)
}
