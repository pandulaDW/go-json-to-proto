package transactions

func toProtoSingleTransaction(obj *SingleJsonTransaction) *SingleTransaction {
	return &SingleTransaction{
		Date:            obj.Date.Date.UnixNano(),
		Amount:          obj.Amount,
		TransactionCode: obj.TransactionCode,
		Price:           obj.Price,
		Total:           obj.Total,
	}
}

func toProtoTransaction(obj *TransactionJsonType) *Transaction {
	transactionList := make([]*SingleTransaction, len(obj.Transactions))
	for i, item := range obj.Transactions {
		transactionList[i] = toProtoSingleTransaction(&item)
	}

	return &Transaction{
		Id:               obj.ID.ID,
		AccountId:        obj.AccountID,
		TransactionCount: obj.TransactionCount,
		BucketStartDate:  obj.BucketStartDate.Date.UnixNano(),
		BucketEndDate:    obj.BucketEndDate.Date.UnixNano(),
		Transactions:     transactionList,
	}
}
