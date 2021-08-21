package controllers

import (
	"apiStore/db"
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo/protos/api"
)

func Mutationproducts() {

	dgraphClient := db.ConnectionDgraph()
	ctx := context.Background()

	//Data Products

	products := ReadProductsCsv()
	println(products)

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	response, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	println(response)

}
