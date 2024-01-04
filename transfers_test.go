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

	res, err := squad.AccountLookup(bank, account)
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

	res, err := squad.AccountLookup(bank, account1)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid account no")

	res1, err1 := squad.AccountLookup("", account)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "please provide a bank code")
}
