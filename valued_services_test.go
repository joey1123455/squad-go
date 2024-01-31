package squad

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_squadBaseACC_VendAirtime_wrong_input(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewServicesClient()
	res, err := utilClient.VendAirtime("08111", 100)
	t.Log(err)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "please provide a valid phone number")

	res1, err1 := utilClient.VendAirtime("08118997115", 0)
	t.Log(err1)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "please provide an amount more then 0")
	// t.Log(res)
}

func Test_squadBaseACC_VendAirtime(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewServicesClient()
	res, err := utilClient.VendAirtime("08118997115", 100)
	t.Log(res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
	// t.Log(res)
}

func Test_squadBaseACC_VendDataBundles(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewServicesClient()
	res, err := utilClient.VendDataBundles("08118997115", "1001", 100)
	t.Log(res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
	// t.Log(res)
}

func Test_squadBaseACC_VendDataBundles_wrong_input(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewServicesClient()
	res, err := utilClient.VendDataBundles("08111", "1001", 100)
	t.Log(err)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "please provide a valid phone number")

	res1, err1 := utilClient.VendDataBundles("08118997115", "1001", 0)
	t.Log(err1)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "please provide an amount more then 0")
	// t.Log(res)
}

func Test_squadBaseACC_GetDataBundles(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewServicesClient()
	res, err := utilClient.GetDataBundles("mtn")
	t.Log(res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
	// t.Log(res)
}

func Test_squadBaseACC_TransactionHistory(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewServicesClient()
	res, err := utilClient.TransactionHistory("1", "10", "debit")
	t.Log(res)
	t.Log(err)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
	// t.Log(res)
}

func Test_squadBaseACC_TransactionHistory_wrong_input(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false

	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	assert.NotNil(t, squad)

	utilClient := squad.NewServicesClient()
	res, err := utilClient.TransactionHistory("1", "10", "hello")
	t.Log(err)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "action must be 'credit' or 'debit'")

	res1, err1 := utilClient.TransactionHistory("a", "10", "debit")
	t.Log(err1)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "page must be greater then 0")

	res2, err2 := utilClient.TransactionHistory("1", "yy", "debit")
	t.Log(err2)
	assert.Nil(t, res2)
	assert.Error(t, err2)
	assert.EqualError(t, err2, "perPage must be greater then 0")
}
