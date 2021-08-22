package controllers

import (
	"apiStore/info"
	sc "apiStore/schema"
	"fmt"
)

//Structure the information related to its types
func RelationInfo() {
	//Buscar por cada usuario las transacciones relaciondas

	buyersall := info.ReadBuyersJson()

	var buyerStore []sc.InfoBuyer

	for i := 0; i < len(buyersall); i++ {
		transactionsbuyer := foundTransaction(buyersall[i].Id)

		newBuyer := sc.InfoBuyer{
			Id:           buyersall[i].Id,
			Name:         buyersall[i].Name,
			Age:          buyersall[i].Age,
			Transactions: transactionsbuyer,
		}

		buyerStore = append(buyerStore, newBuyer)

	}

	fmt.Println(buyerStore)

	//Hacer la mutacion
	MutationInfoBuyers(buyerStore)

}

//Funcion que actualiza la información de la transacción de los productos, agrega los id de la db para relacionar los nodos
func RelationProductos() []sc.Transaction {
	type idProduct struct {
		Uid string `json:"uid"`
	}
	//Traer las transacciones
	//Ese sería el nodo relacionado al usuario
	allTransactions := info.ReadInfoTxt()
	//Recorrer el arreglo y y comparar los id con los que se guardaron en l db

	//Recorre las transacciones y los productos de esas transacciones que es una array []
	for i := 0; i < len(allTransactions); i++ {
		for j := 0; j < len(allTransactions[i].Products); j++ {
			uid := foundIdProduct(allTransactions[i].Products[j])
			idproductr := idProduct{
				Uid: uid,
			}
			allTransactions[i].Products[j] = idproductr

		}

	}

	return (allTransactions)

	//fmt.Println(allTransactions)

}

func foundIdProduct(id string) string {

	//Trar los guardador en la db
	dbproducts := Getproducts()
	//fmt.Println(dbproducts)

	var idproductdb string

	for i := 0; i < len(dbproducts); i++ {
		if dbproducts[i].Id == id {
			idproductdb = dbproducts[i].Uid

		}
	}
	return (idproductdb)

}

func foundTransaction(idBuyer string) []sc.Transaction {

	transactions := RelationProductos()

	var transactionxbuyer []sc.Transaction

	for i := 0; i < len(transactions); i++ {

		if transactions[i].Buyer_id == idBuyer {
			transactionxbuyer = append(transactionxbuyer, transactions[i])

		}

	}

	return transactionxbuyer

}
