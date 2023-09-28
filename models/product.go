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
	productsTable, err := db.Query("select * from produtos order by id asc")

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

func DeleteProduct(productID string) {
	db := database.DatabaseConnect()

	deleteFromDatabase, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteFromDatabase.Exec(productID)
	defer db.Close()
}

func EditProduct(productId string) Product {
	db := database.DatabaseConnect()

	dbProduct, err := db.Query("select * from produtos where id=$1", productId)

	if err != nil {
		panic(err.Error())
	}

	attProduct := Product{}

	for dbProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		attProduct.Name = name
		attProduct.Description = description
		attProduct.Price = price
		attProduct.Quantity = quantity
	}
	defer db.Close()

	return attProduct
}

func UpdateProduct(name, description string, price float64, quantity int, id int) {
	db := database.DatabaseConnect()

	UpdateProduct, err := db.Prepare("update produtos set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	UpdateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
