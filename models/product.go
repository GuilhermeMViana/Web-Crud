package models

import (
	"web/database"
)

type Product struct {
	Name, Description string
	Id, Quantity      int
	Price             float64
}

func SearchProducts() []Product {
	db := database.DatabaseConnect()
	productsTable, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for productsTable.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsTable.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := database.DatabaseConnect()
	insertInDatabase, err := db.Prepare("insert into produtos(name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertInDatabase.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeletProduct(productID string) {
	db := database.DatabaseConnect()

	deleteFromDatabase, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteFromDatabase.Exec(productID)
	defer db.Close()
}
