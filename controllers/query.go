package controllers

import (
	"apiStore/db"
	sc "apiStore/schema"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func Getproducts() {

	dgraphClient := db.ConnectionDgraph()
	txn := dgraphClient.NewTxn()

	const q = `
	{
		products(func:has(Name)){
		uid
		Id
		Name
		Price
	  }
	}
	`

	type Root struct {
		Products []sc.Product `json:"products"`
	}

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Products[0])

	/*out, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", out)*/

}
