package system

import (
	"io/ioutil"
	"log"
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

func ImportBilling() BillingData {
	file, _ := os.Open("billing.data")
	mask, err := ioutil.ReadAll(file)
	err = file.Close()
	if err != nil {
		log.Panicln(err)
	}

	if len(mask) != 6 {
		return BillingData{}
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
	}
}
