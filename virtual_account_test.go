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

func Test_QueryMerchantVirtualAccHistory_success(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := false

	res, err := QueryMerchantVirtualAccHistory(apiKey, live)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_QueryMerchantVirtualAccHistory_wrong_input_sandbox_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("TEST_WRONG_LIVE_API_KEY")
	live := false

	res, err := QueryMerchantVirtualAccHistory(apiKey, live)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for test account must start with 'sandbox_sk'")
}

func Test_QueryMerchantVirtualAccHistory_wrong_input_live_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := true

	res, err := QueryMerchantVirtualAccHistory(apiKey, live)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for account must start with 'sk'")
}

func Test_QueryMerchantHistoryFilters_success(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := false
	filter := map[string]string{
		"page":               "1",
		"perPage":            "10",
		"virtualAccount":     "",
		"customerIdentifier": "",
		"startDate":          "09-19-2022",
	}
	filter2 := map[string]string{
		"page":               "",
		"perPage":            "",
		"virtualAccount":     "",
		"customerIdentifier": "",
		"startDate":          "",
	}
	filter3 := map[string]string{
		"page":               "",
		"perPage":            "",
		"virtualAccount":     "",
		"customerIdentifier": "",
		"startDate":          "",
		"pain":               "kinglanding",
	}

	res, err := QueryMerchantHistoryFilters(apiKey, live, filter)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
	t.Log(err)

	res1, err1 := QueryMerchantHistoryFilters(apiKey, live, nil)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "if no filters are provided make use of the QueryMerchantHistory function instead")

	res2, err2 := QueryMerchantHistoryFilters(apiKey, live, filter2)
	assert.Nil(t, res2)
	assert.Error(t, err2)
	assert.EqualError(t, err2, "if no filters are provided make use of the QueryMerchantHistory function instead")

	res3, err3 := QueryMerchantHistoryFilters(apiKey, live, filter3)
	assert.Nil(t, res3)
	assert.Error(t, err3)
	assert.EqualError(t, err3, "if no filters are provided make use of the QueryMerchantHistory function instead")
}

func Test_QueryMerchantHistoryFilters_wrong_input_api_key(t *testing.T) {
	filter := map[string]string{
		"page":    "1",
		"perPage": "10",
	}
	// live key to sand box endpoint
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := true

	res, err := QueryMerchantHistoryFilters(apiKey, live, filter)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for account must start with 'sk'")

	// sand box key to live endpoint
	apiKey1 := os.Getenv("TEST_WRONG_LIVE_API_KEY")
	live1 := false

	res1, err1 := QueryMerchantHistoryFilters(apiKey1, live1, filter)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "api key for test account must start with 'sandbox_sk'")

}
