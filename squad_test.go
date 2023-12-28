package squad

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_NewSquadObj_success(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	squadImp := squad.(*squadBaseACC)
	assert.NotNil(t, squad)
	assert.Nil(t, err)
	assert.Equal(t, name, squadImp.AccountName)
	assert.Equal(t, url, squadImp.CallBack)
	assert.Equal(t, live, squadImp.Live)
}

func Test_NewSquadObj_wrong_input_url(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "calback/wrong.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, squad)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid callback url, should have a valid protocol 'http://' or 'https://'")
}

func Test_NewSquadObj_wrong_apiKey_live(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := true
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, squad)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for account must start with 'sk'")
}

func Test_NewSquadObj_wrong_apiKey_test(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("TEST_WRONG_LIVE_API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, squad)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for test account must start with 'sandbox_sk'")
}

func Test_NewSquadObj_wrong_input_no_key(t *testing.T) {
	apiKey := ""
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, squad)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key must be provided")
}

func Test_NewSquadObj_wrong_input_no_bussiness_name(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := ""
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, squad)
	assert.Error(t, err)
	assert.EqualError(t, err, "please provide a bussiness name")
}

func Test_SquadBaseAcc_CreatePaymentObject(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	charge := false
	paymentChan := []string{"card", "bank", "ussd"}
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	payment := squad.CreatePaymentObject(charge, paymentChan)
	payImp := payment.(*paymentObjectImp)
	assert.Equal(t, url, payImp.CallBack)
	assert.Equal(t, apiKey, payImp.ApiKey)
	assert.Equal(t, live, payImp.live)
	assert.Equal(t, charge, payImp.PassCharge)
	assert.Equal(t, paymentChan, payImp.PaymentChans)
}
