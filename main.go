package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"protobuf_using_go/transactions"
	"time"
)

func main() {
	start := time.Now()

	allTransactions := transactions.ReadJsonLines()
	err := transactions.WriteProtoToFile("data/encoded.bin", allTransactions)
	if err != nil {
		log.Fatal(err)
	}

	content, err := transactions.ReadProtoFromFile("data/encoded.bin")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range content {
		t := new(transactions.Transaction)
		err = proto.Unmarshal(item, t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("t: %s with account id: %d and t_count: %d, with first transaction: %v\n",
			t.Id, t.AccountId, t.TransactionCount, t.Transactions[0])
	}

	fmt.Println("took", time.Since(start))
}
