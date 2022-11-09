package main

import (
	"context"
	"log"
	"zenapi/swagger"

	"github.com/google/uuid"
)

func main(){

	conf := swagger.NewConfiguration()
	conf.AddDefaultHeader("Authorization","fd3e4eaadcca65c9a4cea18f2eb8d85a")
	c := swagger.NewAPIClient(conf) 
	// log.Println(c)


	result, _, err := c.TransactionsApi.CreateTransaction(
		context.Background(),	
		swagger.ActionCreateObject{
			MerchantTransactionId: uuid.NewString(),
			PaymentChannel:        "PCL_CARD",
			Amount:                "100",
			Currency:              "PLN",
			Customer:              &swagger.TransactionsCustomer{
				Email:     "neroslava@ya.ru",
				Ip:        "8.8.8.8",
			},
		},
		"Bearer a41a0314b76b4c0db12201a1fb914bb5",
		"ZEN12345",
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(result)
}