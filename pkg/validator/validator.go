// Test credit card numbers: https://www.paypalobjects.com/en_AU/vhelp/paypalmanager_help/credit_card_numbers.htm

package validator

import (
	"fmt"
	"os"
	"strconv"

	"github.com/theplant/luhn"
)

type creditCardDetails struct {
	bankName   string
	cardNumber string
	expiry     string
}

func (c creditCardDetails) checkCardLength() string {
	cardNumber := c.cardNumber
	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		fmt.Printf("Card number is invalid: %s\n", cardNumber)
		fmt.Println(len(cardNumber))
		os.Exit(1)
	}
	return cardNumber
}

func (c creditCardDetails) lunhValidator() bool {
	cardNumber, err := strconv.Atoi(c.cardNumber)
	if err != nil {
		fmt.Printf("Failed to convert the card number %s in integer ", c.cardNumber)
	}
	luhn := luhn.Valid(cardNumber)

	//return true or false
	return luhn

}

// func main() {
// 	cardNumber := creditCardDetails{bankName: "HDFC", cardNumber: "6011111111111117", expiry: "12/30"}

// 	cardLength, validCardNumber := cardNumber.checkCardLength(), cardNumber.lunhValidator()
// 	fmt.Println(cardLength, validCardNumber)

// }
