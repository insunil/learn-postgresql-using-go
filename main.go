package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var con_err error
	db, con_err = sql.Open("postgres", os.Getenv("POSTGRESQL_URI"))
	if err != nil {
		log.Fatal("Failed to connect:", con_err)
	}

}
func main() {
	for {
		fmt.Println("\n--- Product Menu ---")
		fmt.Println("1. Insert product")
		fmt.Println("2. Show all products")
		fmt.Println("3. Show specific")
		fmt.Println("4. Update product")
		fmt.Println("5. Delete product")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			var price float64
			fmt.Print("Enter product name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter product price: ")
			fmt.Scanln(&price)
			insert(name, price)
		case 2:
			getAll()

		case 3:
			var id int
			fmt.Print("Enter product ID to get: ")
			fmt.Scanln(&id)
			getSpecific(id)

		case 4:
			var id int
			var name string
			var price float64
			fmt.Print("Enter product ID to update: ")
			fmt.Scanln(&id)
			fmt.Print("Enter new name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter new price: ")
			fmt.Scanln(&price)
			update(id, name, price)
		case 5:
			var id int
			fmt.Print("Enter product ID to delete: ")
			fmt.Scanln(&id)
			deleteSpecific(id)
		case 6:
			fmt.Println("Exiting.")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
