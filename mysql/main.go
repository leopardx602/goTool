package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	username  string = "root"
	password  string = "password"
	addr      string = "127.0.0.1"
	port      int    = 3306
	database  string = "db01"
	parseTime bool   = true // time.time or string
)

// sql.ErrNoRows

type Product struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Price    int       `json:"age"`
	Image    string    `json:"image"`
	CreateAt string    `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func Create(conn *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS table01(
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(16) NOT NULL DEFAULT "",
		price INT DEFAULT 0,
		image VARCHAR(64) DEFAULT "",
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`
	// CONSTRAINT id_name PRIMARY KEY (id, name)
	_, err := conn.Query(sql)
	if err != nil {
		return err
	}
	return nil
}

func CreateEvent(conn *sql.DB) error {
	sql := `CREATE EVENT IF NOT EXISTS event01
		ON SCHEDULE EVERY 10 SECOND
		ON COMPLETION PRESERVE
		DO
	  	delete from table01 where TO_SECONDS(now())-TO_SECONDS(created_at)>120;`
	_, err := conn.Query(sql)
	if err != nil {
		return err
	}
	return nil
}

func Insert(conn *sql.DB, product Product) error {
	_, err := conn.Exec("INSERT INTO table01 (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				fmt.Println("already exists")
				return nil
			}
		}
	}
	return nil
}

func Select(conn *sql.DB, command string) error {
	res, err := conn.Query(command)
	if err != nil {
		return err
	}
	defer res.Close()

	var products []Product
	for res.Next() {
		var product Product
		if err := res.Scan(&product.ID, &product.Name, &product.Price, &product.Image, &product.CreateAt, &product.UpdateAt); err != nil {
			return err
		}
		products = append(products, product)
	}
	fmt.Println(products)
	return nil
}

func Update(conn *sql.DB, product Product) error {
	result, err := conn.Exec("UPDATE table01 SET name=?, price=?, image=? WHERE id=?", product.Name, product.Price, product.Image, product.ID) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		return err
	}
	fmt.Println(result.RowsAffected())
	return nil
}

func Delete(conn *sql.DB, id int) error {
	_, err := conn.Exec("DELETE FROM table01 WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v", username, password, addr, port, database, parseTime)
	conn, err := sql.Open("mysql", connInfo)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// create talble
	if err = Create(conn); err != nil {
		panic(err)
	}

	// // insert
	// if err = Insert(conn, Product{ID: 2, Name: "iphone2", Price: 2000}); err != nil {
	// 	fmt.Println(err)
	// }

	// update
	if err = Update(conn, Product{ID: 2, Name: "iphone2", Price: 2000}); err != nil {
		fmt.Println(err)
	}

	// // delete
	// if err = Delete(conn, 1); err != nil {
	// 	fmt.Println(err)
	// }

	// select
	if err = Select(conn, "SELECT * FROM table01"); err != nil { // SELECT COUNT(*)
		fmt.Println(err)
	}

	// // select one
	// var product Product
	// if err = conn.QueryRow("SELECT name, price FROM table01 where id = ?", 1).Scan(&product.Name, &product.Price); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(product.Name, product.Price)
}
