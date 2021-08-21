package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ReadBuyersJson() {

	jsonFile, err := os.Open("../Info/buyers.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	//Info
	var buyers []Buyer

	// we unmarshal our byteArray which contains our jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &buyers)
	fmt.Println(buyers)

}

func ReadProductsCsv() []Product {

	dFile, err := os.Open("../Info/products.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer dFile.Close()

	r := csv.NewReader(dFile)
	r.Comma = ','
	r.FieldsPerRecord = 3

	var products []Product
	//Leer linea por línea
	for {
		record, err := r.Read()
		//Error al final del archivo
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error", err)
		}

		//Mirar vacio
		productNew := Product{
			Id:   record[0],
			Name: record[1],
		}
		if record[2] != "" {
			price, err := strconv.Atoi(record[2])
			if err != nil {
				log.Println("Error price", err)
				continue
			}
			productNew.Price = price

		}

		products = append(products, productNew)

	}
	fmt.Println(products)
	return products

}

func readInfoTxt() {

	transactionsAll := Transaction{

		Id:       "#0000611ef080",
		Buyer_id: "c9288b31",
		Ip:       "126.238.179.254",
		device:   "android",
		Products: []string{"96028f6d", "7c6b9e30"},
	}

	fmt.Println(transactionsAll)

}
