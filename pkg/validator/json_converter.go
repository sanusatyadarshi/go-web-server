package validator

import (
	"encoding/json"
	"fmt"
)

type CreditCardDetails struct {
	BankName   string `json:"bankName"`
	CardNumber string `json:"cardNumber"`
	Expiry     string `json:"expiry"`
}

func main() {
	jsonInput := `{
        "bankName": "SBI",
        "cardNumber": "6011111111111117",
        "expiry": "11/24"
    }`
	var c CreditCardDetails
	err := json.Unmarshal([]byte(jsonInput), &c)

	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	fmt.Printf("Bank Name: %s, Card Number: %s, Expiry: %s\n", c.BankName, c.CardNumber, c.Expiry)
}
