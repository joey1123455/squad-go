package squad

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_bussinessVA_Initiate(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	vAName := "james cord"
	no := "08018995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_bussinessVA_QueryVirtualAccHistory(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	vAName := "james cord"
	no := "08018995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.QueryVirtualAccHistory()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_customerVA_Initiate(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	first := "ibrahim"
	last := "james"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village ATBU"
	gender := "1"
	no := "08018995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_customerVA_QueryVirtualAccHistory(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	first := "ibrahim"
	last := "james"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village ATBU"
	gender := "1"
	no := "08018995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.QueryVirtualAccHistory()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_MissedWebHookNotifications_success(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := false
	page := 1
	perPage := 10

	res, err := MissedWebHookNotifications(apiKey, live, page, perPage)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_MissedWebHookNotifications_wrong_input_live_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := true
	page := 1
	perPage := 10

	res, err := MissedWebHookNotifications(apiKey, live, page, perPage)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for account must start with 'sk'")
}

func Test_MissedWebHookNotifications_wrong_input_sandbox_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("TEST_WRONG_LIVE_API_KEY")
	live := false
	page := 1
	perPage := 10

	res, err := MissedWebHookNotifications(apiKey, live, page, perPage)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for test account must start with 'sandbox_sk'")
}

func Test_MissedWebHookNotifications_wrong_input_page_and_perPage(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := false
	page := 0
	perPage := 10
	page1 := 1
	perPage1 := 0

	res, err := MissedWebHookNotifications(apiKey, live, page, perPage)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "page must have a value greater then 0")

	res1, err1 := MissedWebHookNotifications(apiKey, live, page1, perPage1)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "perPage must have a value greater then 0")
}
