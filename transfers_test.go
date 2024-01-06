package squad

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_squadBaseACC_AccountLookup(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	bank := "000015"
	account := "2218347027"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)
	transferCl := squad.NewTransferClient()

	res, err := transferCl.AccountLookup(bank, account)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_squadBaseACC_Wrong_input(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	bank := "000015"
	account := "2218347027"
	account1 := "221834702"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	transferCl := squad.NewTransferClient()
	res, err := transferCl.AccountLookup(bank, account1)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid account no")

	res1, err1 := transferCl.AccountLookup("", account)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "please provide a bank code")
}

func Test_squadBaseACC_FundTransfer(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	data := map[string]any{

		"remark":                "for test transfer to my customer",
		"bank_code":             "000013",
		"currency_id":           "NGN",
		"amount":                "100",
		"account_number":        "0123456789",
		"transaction_reference": "SBABCKDY_12345",
		"account_name":          "BOLUS PAUL",
	}
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	transferCl := squad.NewTransferClient()
	res, err := transferCl.FundTransfer(data)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_squadBaseACC_GetAllTransfers(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	transferCl := squad.NewTransferClient()
	res, err := transferCl.GetAllTransfers("1", "10", "asc")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_squadBaseACC_RequeryTransfer(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	transferCl := squad.NewTransferClient()
	res, err := transferCl.RequeryTransfer("1tyrndiirlsos;ie")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
