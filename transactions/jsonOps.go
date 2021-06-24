package transactions

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"
)

type SingleJsonTransaction struct {
	Date struct {
		Date time.Time `json:"$date"`
	} `json:"date"`
	Amount          int32   `json:"amount"`
	TransactionCode string  `json:"buy"`
	Symbol          string  `json:"symbol"`
	Price           float32 `json:"price,string"`
	Total           float32 `json:"total,string"`
}

type TransactionJsonType struct {
	ID struct {
		ID string `json:"$oid"`
	} `json:"_id"`
	AccountID        int32 `json:"account_id"`
	TransactionCount int32 `json:"transaction_count"`
	BucketStartDate  struct {
		Date time.Time `json:"$date"`
	} `json:"bucket_start_date"`
	BucketEndDate struct {
		Date time.Time `json:"$date"`
	} `json:"bucket_end_date"`
	Transactions []SingleJsonTransaction `json:"transactions"`
}

func ReadJsonLines() []*TransactionJsonType {
	file, err := os.Open("data/transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	transactions := make([]*TransactionJsonType, 0)

	for scanner.Scan() {
		line := scanner.Bytes()
		jsonObj := new(TransactionJsonType)
		err = json.Unmarshal(line, jsonObj)
		if err != nil {
			continue
		}

		transactions = append(transactions, jsonObj)
	}
	return transactions
}
