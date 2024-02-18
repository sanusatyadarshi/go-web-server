// Test credit card numbers: https://www.paypalobjects.com/en_AU/vhelp/paypalmanager_help/credit_card_numbers.htm

package validator

import (
	"fmt"
	"os"
	"strconv"

	"github.com/theplant/luhn"
)

type CreditCardDetails struct {
	BankName   string
	CardNumber string
	Expiry     string
}

func (c CreditCardDetails) CheckCardLength() string {
	cardNumber := c.CardNumber
	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		fmt.Printf("Card number is invalid: %s\n", cardNumber)
		fmt.Println(len(cardNumber))
		os.Exit(1)
	}
	return cardNumber
}

func (c CreditCardDetails) LunhValidator() bool {
	cardNumber, err := strconv.Atoi(c.CardNumber)
	if err != nil {
		fmt.Printf("Failed to convert the card number %s in integer ", c.CardNumber)
	}
	luhn := luhn.Valid(cardNumber)

	//return true or false
	return luhn

}

// func main() {
// 	cardNumber := CreditCardDetails{BankName: "HDFC", CardNumber: "6011111111111117", Expiry: "11/24"}

// 	cardLength, validCardNumber := cardNumber.CheckCardLength(), cardNumber.LunhValidator()
// 	fmt.Println(cardLength, validCardNumber)

// }
