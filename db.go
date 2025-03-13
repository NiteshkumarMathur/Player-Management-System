package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBGetAllPlayers() []Footballer {
	fmt.Println(" Reached DBGetAllPlayers")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/data")

	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM persons")
	fmt.Println(" Reached results", results)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	products := []Footballer{}
	for results.Next() {
		var prod Footballer
		// for each row, scan into the Product struct
		err = results.Scan(&prod.UUID, &prod.Name, &prod.Club, &prod.Country)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		fmt.Println(" Reached prod", prod)

		products = append(products, prod)
	}
	fmt.Println(" Reached products", products)

	return products
}

func DBAddPlayer(newplayer Footballer) error {
	fmt.Println(" Reached DBAddPlayers")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/data")

	if err != nil {
		return err
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()
	max, err := db.Query("SELECT Player_UUID FROM persons ORDER BY Player_UUID DESC LIMIT 0, 1")
	var UUID int

	for max.Next() {
		err = max.Scan(&UUID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	insert, err := db.Query(
		"INSERT INTO persons (Player_UUID,Player_Name,Club,Country) VALUES (?,?,?,?)",
		UUID+1, newplayer.Name, newplayer.Club, newplayer.Country)

	fmt.Println(" Reached results", newplayer)
	// if there is an error inserting, handle it
	if err != nil {
		return err
	}

	defer insert.Close()
	return nil
}
func DBDeleteByUUID(UUID int) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/data")
	if err != nil {
		fmt.Println(err)
	}

	// close database after all work is done
	defer db.Close()

	// delete data
	stmt, err := db.Prepare("delete from persons where Player_UUID=?")
	if err != nil {
		fmt.Println(err)
	}
	// delete 1st student
	_, err = stmt.Exec(UUID)
	if err != nil {
		fmt.Println(err)
	}
}

func DBUpdateByUuid(Player_UUID int, Player_Name string) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/data")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// close database after all work is done
	defer db.Close()

	stmt, err := db.Prepare("update persons set Player_Name=? where Player_UUID=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// delete 1st student
	_, err = stmt.Exec(Player_Name, Player_UUID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
