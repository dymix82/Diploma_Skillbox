package system

import (
	"errors"
	"io/ioutil"
	"math"
	"os"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

const (
	BillingCreateCustomer = iota
	BillingPurchase
	BillingPayout
	BillingRecurring
	BillingFraudControl
	BillingCheckoutPage
)

// Импорт файла с маской для биллинга
func ImportBilling() (BillingData, error) {
	file, err := os.Open("billing.data")
	if err != nil {
		return BillingData{}, err
	}
	mask, err2 := ioutil.ReadAll(file)
	err2 = file.Close()
	if err2 != nil {
		return BillingData{}, err2
	}

	if len(mask) != 6 {
		return BillingData{}, errors.New("Billing: Данные в маске не полные")
	}

	var bitMask int8 = 0
	for i, bit := range mask {
		if bit == '1' {
			bitMask += int8(math.Pow(2, float64(len(mask)-i-1)))
		}
	}

	createCustomer := bitMask>>BillingCreateCustomer&1 == 1
	purchase := bitMask>>BillingPurchase&1 == 1
	payout := bitMask>>BillingPayout&1 == 1
	recurring := bitMask>>BillingRecurring&1 == 1
	fraudControl := bitMask>>BillingFraudControl&1 == 1
	checkoutPage := bitMask>>BillingCheckoutPage&1 == 1

	return BillingData{
		CreateCustomer: createCustomer,
		Purchase:       purchase,
		Payout:         payout,
		Recurring:      recurring,
		FraudControl:   fraudControl,
		CheckoutPage:   checkoutPage,
	}, nil
}
