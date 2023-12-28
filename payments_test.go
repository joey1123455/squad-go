package squad

import (
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// go test -run Test_paymentObjectImp_convert
func Test_paymentObjectImp_convert(t *testing.T) {
	squad := paymentObjectImp{}
	amt1 := squad.convert(10)
	amt2 := squad.convert(10.00)
	amt3 := squad.convert(10000000.00)
	amt4 := squad.convert(145065.72)
	assert.Equal(t, 1000, amt1)
	assert.NotEqual(t, 100, amt2)
	assert.Equal(t, 1000000000, amt3)
	assert.Equal(t, 14506572, amt4)
}

// go test -run Test_paymentObjectImp_completeUrl_false
func Test_paymentObjectImp_completeUrl_false(t *testing.T) {
	squad := paymentObjectImp{
		live: false,
	}
	url := squad.completeUrl("transaction/initiate")
	notExpected := "https://api-d.squadco.com/transaction/initiate"
	expected := "https://sandbox-api-d.squadco.com/transaction/initiate"
	assert.Equal(t, expected, url)
	assert.NotEqual(t, notExpected, url)
}

// go test -run Test_paymentObjectImp_completeUrl_true
func Test_paymentObjectImp_completeUrl_true(t *testing.T) {
	squad := paymentObjectImp{
		live: true,
	}
	url := squad.completeUrl("transaction/initiate")
	expected := "https://api-d.squadco.com/transaction/initiate"
	notExpected := "https://sandbox-api-d.squadco.com/transaction/initiate"
	assert.Equal(t, expected, url)
	assert.NotEqual(t, notExpected, url)
}

// go test -run Test_paymentObjectImp_Initiate_success
func Test_paymentObjectImp_Initiate_success(t *testing.T) {
	var expectedAmount float64
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	cur := "usd"
	amount := float64(100)
	charge := false
	payChan := []string{"card", "bank", "ussd"}

	switch {
	case charge:
		if cur == "usd" || cur == "USD" {
			expectedAmount = amount*100 + float64(350)
		} else if cur == "ngn" || cur == "NGN" {
			expectedAmount = amount*100 + float64(100)
		}
	case !charge:
		if cur == "usd" || cur == "USD" {
			expectedAmount = amount * 100
		} else if cur == "ngn" || cur == "NGN" {
			expectedAmount = amount * 100
		}
	}

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)
	payObj := squad.CreatePaymentObject(charge, payChan)
	res, err := payObj.Initiate(amount, cur, "", map[string]string{"email": "eg@mail.com"}, nil, false)
	t.Log(res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, float64(200), res["status"])
	assert.Equal(t, true, res["success"])
	data := res["data"]
	myData := data.(map[string]any)
	t.Log(myData)
	assert.Equal(t, strings.ToUpper(cur), myData["currency"])
	assert.Equal(t, expectedAmount, myData["transaction_amount"])
	assert.Equal(t, false, myData["is_recurring"])
	expectedChans := []any{"card", "bank", "ussd"}
	assert.Equal(t, expectedChans, myData["authorized_channels"])
}

// go test -run Test_paymentObjectImp_Initiate_wrong_input_no_customer_map
func Test_paymentObjectImp_Initiate_wrong_input_no_customer_map(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	cur := "usd"
	amount := float64(100)
	charge := false
	payChan := []string{"card", "bank", "ussd"}
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)
	payObj := squad.CreatePaymentObject(charge, payChan)
	res, err := payObj.Initiate(amount, cur, "", nil, nil, false)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, "customer map must be passed")
}

// go test -run Test_paymentObjectImp_Initiate_wrong_input_currency
func Test_paymentObjectImp_Initiate_wrong_input_currency(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	cur := "abc"
	amount := float64(100)
	charge := false
	payChan := []string{"card", "bank", "ussd"}
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)
	payObj := squad.CreatePaymentObject(charge, payChan)
	res, err := payObj.Initiate(amount, cur, "", map[string]string{"email": "eg@mail.com"}, nil, false)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, "currency should be NGN or USD")
}

// go test -run Test_paymentObjectImp_Initiate_wrong_input_no_customer_email
func Test_paymentObjectImp_Initiate_wrong_no_customer_email(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	cur := "usd"
	amount := float64(100)
	charge := false
	payChan := []string{"card", "bank", "ussd"}
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)
	payObj := squad.CreatePaymentObject(charge, payChan)
	res, err := payObj.Initiate(amount, cur, "", map[string]string{"name": "joseph folayan"}, nil, false)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, "customer email must be provided")
}

func Test_paymentObjectImp_Initiate_wrong_input_amount(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	cur := "usd"
	amount := float64(0)
	charge := false
	payChan := []string{"card", "bank", "ussd"}
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)
	payObj := squad.CreatePaymentObject(charge, payChan)
	res, err := payObj.Initiate(amount, cur, "", map[string]string{"email": "eg@mail.com"}, nil, false)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, "amount is not provided")
}
