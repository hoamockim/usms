package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"usms/client/dtos"
)

func Test_Transaction(t *testing.T) {
	transactionUrl := "https://api-payment-uat.int.vinid.dev/internal/transaction-history/v1"
	apiKey := "123456aA@"

	requestTran := dtos.TransactionRequest{
		Filter: &dtos.TransactionFilter{
			ActivityType: struct {
				In []string `json:"in"`
			}{
				In: []string{"WALLET_GROUP_IN", "WALLET_GROUP_OUT"},
			},
			GroupId: struct {
				In []string `json:"in"`
			}{
				In: []string{"63176", "63199", "63239", "63221", "63250", "63280", "63293", "63168", "63176"},
			},
			TransactionStatus: struct {
				In []string `json:"in"`
			}{
				In: []string{"SUCCESS", "FAIL"},
			},
		},
		Pageable: struct {
			Limit  int `json:"limit"`
			Offset int `json:"offset"`
		}{
			Limit:  50,
			Offset: 0,
		},
		Sort: struct {
			TransactionCreateTime string `json:"transaction_create_time"`
		}{TransactionCreateTime: "Desc"},
	}
	var responseData dtos.TransactionResponse
	ctx := context.Background()
	client := NewClient(ctx, "9c9c0097-4969-4a5c-8879-30aa0ae5be17", JsonContentType, "", apiKey, transactionUrl, &responseData)

	transactionEndpoint := client.makeEndpoint(ctx, "transaction-filters")
	res, err := transactionEndpoint(ctx, &requestTran)
	if err != nil {
		fmt.Println("err: ", err)
	}
	response, _ := res.(*dtos.TransactionResponse)
	resStr, _ := json.Marshal(&response)
	fmt.Println(string(resStr))
}
