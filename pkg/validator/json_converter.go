package validator

type CreditCardDetails struct {
	bankName   string `json:"bankName"`
	cardNumber string `json:"cardNumber"`
	expiry     string `json:"expiry"`
}
