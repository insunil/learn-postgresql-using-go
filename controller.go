package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

func insert(name string, price float64) {
	id := 0
	err := db.QueryRow("INSERT INTO PRODUCT (Name,Price) VALUES($1,$2) RETURNING Id", name, price).Scan(&id)
	if err != nil {
		slog.Error("insertion issue occurred")
		os.Exit(0)
	}
	fmt.Println(id)
}
func getSpecific(id int) {
	var p Product
	err := db.QueryRow("select *from PRODUCT where Id=$1", id).Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		fmt.Println("error occurred  during get")
		os.Exit(0)
	}
	fmt.Println(p)
}

func getAll() {
	rows, err := db.Query("SELECT id, name, price FROM product")
	if err != nil {
		fmt.Println("error occurred during getting")
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			fmt.Println("Error occurred during scanning")
			return
		}
		fmt.Println(p)
	}

	
}

func update(id int, name string, price float64) {
	res, err := db.Exec("update PRODUCT set Name=$1,Price=$2 where Id=$3", name, price, id)
	if err != nil {
		log.Fatal("error occurred during updating")
	}
	fmt.Println(res)
}

func deleteSpecific(id int) {
	res, err := db.Exec("delete from PRODUCT where Id=$1", id)
	if err != nil {
		log.Fatal("error occurred during deleting")
	}
	fmt.Println(res)
}
