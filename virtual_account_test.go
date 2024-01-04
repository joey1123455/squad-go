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

func Test_DeleteMissedWebHookNotifications_success(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := false

	res, err := DeleteMissedWebHookNotifications(apiKey, "ccghetyeh22", live)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_DeleteMissedWebHookNotifications_wrong_input_live_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := true

	res, err := DeleteMissedWebHookNotifications(apiKey, "ccghetyeh22", live)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for account must start with 'sk'")
}

func Test_DeleteMissedWebHookNotifications_wrong_input_sandbox_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("TEST_WRONG_LIVE_API_KEY")
	live := false

	res, err := DeleteMissedWebHookNotifications(apiKey, "ccghetyeh22", live)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for test account must start with 'sandbox_sk'")
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

func Test_GetMerchantVirtualAcc(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := false

	res, err := GetMerchantVirtualAcc(apiKey, live)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
	t.Log(res)
}

func Test_GetMerchantVirtualAcc_wrong_input_sandbox_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("TEST_WRONG_LIVE_API_KEY")
	live := false

	res, err := GetMerchantVirtualAcc(apiKey, live)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for test account must start with 'sandbox_sk'")
}

func Test_GetMerchantVirtualAcc_wrong_input_live_api_key(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := true

	res, err := GetMerchantVirtualAcc(apiKey, live)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for account must start with 'sk'")
}

func Test_TestPaymentToVA(t *testing.T) {
	var amount float64 = 100
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	live := false
	acc := "3456987768"

	res, err := TestPaymentToVA(apiKey, acc, amount, live)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_TestPaymentToVA_wrong_input(t *testing.T) {
	var amount float64 = 100
	var amount1 float64 = 0

	_ = godotenv.Load()
	acc := "3456987768"
	acc1 := "34561987768"
	apiKey := os.Getenv("API_KEY")
	apiKey1 := os.Getenv("TEST_WRONG_LIVE_API_KEY")
	live := true
	live1 := false

	// live api key
	res, err := TestPaymentToVA(apiKey, acc, amount, live)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "api key for account must start with 'sk'")

	// sand box key
	res1, err1 := TestPaymentToVA(apiKey1, acc, amount, live1)
	assert.Nil(t, res1)
	assert.Error(t, err1)
	assert.EqualError(t, err1, "api key for test account must start with 'sandbox_sk'")

	// wrong amount
	res2, err2 := TestPaymentToVA(apiKey, acc, amount1, live1)
	assert.Nil(t, res2)
	assert.Error(t, err2)
	assert.EqualError(t, err2, "amount must be greater then 0")

	// wrong acc
	res3, err3 := TestPaymentToVA(apiKey, acc1, amount, live1)
	assert.Nil(t, res3)
	assert.Error(t, err3)
	assert.EqualError(t, err3, "invalid account no")
}

func Test_customerVA_AccountDetails(t *testing.T) {
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
	res, err = virAcc.AccountDetails()
	assert.Nil(t, err)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_bussinessVA_AccountDetails(t *testing.T) {
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
	res, err = virAcc.AccountDetails()
	assert.Nil(t, err)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_bussinessVA_AccountDetailsUsingId(t *testing.T) {
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
	res, err = virAcc.AccountDetailsUsingId()
	assert.Nil(t, err)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_customerVA_AccountDetailsUsingId(t *testing.T) {
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
	res, err = virAcc.AccountDetailsUsingId()
	assert.Nil(t, err)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_customerVA_UpdateAccount(t *testing.T) {
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
	accNo1 := "1234587891"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.UpdateAccount(accNo1)
	assert.Nil(t, err)
	t.Log(res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_customerVA_UpdateAccount_wrong_input_(t *testing.T) {
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
	accNo1 := "123458789451"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.UpdateAccount(accNo1)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid account no")
}

func Test_bussinessVA_UpdateAccount(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	vAName := "james cord"
	no := "08018995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	accNo1 := "1254567891"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.UpdateAccount(accNo1)
	assert.Nil(t, err)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_bussinessVA_UpdateAccount_wrong_input_acc_no(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	vAName := "james cord"
	no := "08018994541"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	accNo1 := "125467891"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	t.Log(t, err)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.UpdateAccount(accNo1)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid account no")
}

func Test_bussinessVA_SimulatePayment(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	vAName := "james cord"
	no := "08018995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	amount := 100
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.SimulatePayment(amount)
	assert.Nil(t, err)
	t.Log(res)
	assert.NotNil(t, res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_customerVA_SimulatePayment(t *testing.T) {
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
	amount := 100
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.SimulatePayment(amount)
	assert.Nil(t, err)
	t.Log(res)
	assert.Equal(t, true, res["success"])
	assert.Equal(t, float64(200), res["status"])
}

func Test_customerVA_UpdateBvn(t *testing.T) {
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
	newBvn := os.Getenv("BVN")
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	t.Log(t, err)
	res, err = virAcc.UpdateBvn(newBvn)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_customerVA_UpdateBvn_wrong_input_bvn(t *testing.T) {
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
	newBvn := "235784841607"
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	t.Log(t, err)
	res, err = virAcc.UpdateBvn(newBvn)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid bvn")
}

func Test_bussinessVA_UpdateBvn(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "prince electronics"
	live := false
	vAName := "james cord"
	no := "08018995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	newBvn := os.Getenv("BVN")
	bvn := os.Getenv("BVN")
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.UpdateBvn(newBvn)
	assert.Nil(t, err)
	t.Log(res)
	assert.NotNil(t, res)
}

func Test_bussinessVA_UpdateBvn_wrong_input_bvn(t *testing.T) {
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
	newBvn := "235784841607"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)
	res, err := virAcc.Initiate()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	res, err = virAcc.UpdateBvn(newBvn)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid bvn")
}
