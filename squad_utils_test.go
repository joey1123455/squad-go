package squad

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_squadBaseACC_CreateSubMerchant(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	customerData := map[string]any{
		"display_name":   "james dash",
		"account_name":   "Joseph Folayan",
		"bank_code":      "000015",
		"account_number": "2218347027",
		"bank":           "Zenith Bank Plc",
	}

	utilClient := squad.NewUtilClient()
	res, err := utilClient.CreateSubMerchant(customerData)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_squadBaseACC_WalletBalance(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewUtilClient()
	res, err := utilClient.WalletBalance()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_squadBaseACC_Refund(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	data := map[string]any{

		"gateway_transaction_ref": "wvszqsdrujscpuaofnea529117332_1_1",
		"refund_type":             "Full",
		"reason_for_refund":       "Any reason",
		"transaction_ref":         "vszqsdrujscpua",
	}

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewUtilClient()
	res, err := utilClient.Refund(data)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
