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

func Test_squadBaseACC_parsedVirtualAccName(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	customerName := "Joseph Folayan"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	squadImp := squad.(*squadBaseACC)
	res := squadImp.parseVirtualAccName(customerName)

	assert.Equal(t, "test-Joseph Folayan", res)

	name1 := "test"
	customerName1 := "James Muhammed"
	squad1, err := NewSquadObj(apiKey, url, name1, live)
	assert.Nil(t, err)
	squadImp1 := squad1.(*squadBaseACC)
	res1 := squadImp1.parseVirtualAccName(customerName1)
	assert.Equal(t, "test-James Muhammed", res1)
}

func Test_squadBaseACC_NewBussinessVirtualAcc_success(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	vAName := "joseph folayan"
	no := "08118995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567891"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	expectedName := squad.(*squadBaseACC).parseVirtualAccName(vAName)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)

	virImp := virAcc.(*bussinessVA)
	assert.Equal(t, id, virImp.customerID)
	assert.Equal(t, expectedName, virImp.bussinessName)
	assert.Equal(t, name, virImp.accountName)
	assert.Equal(t, apiKey, virImp.apiKey)
	assert.Equal(t, no, virImp.mobileNo)
	assert.Equal(t, accNo, virImp.beneficiaryAcc)
	assert.Equal(t, bvn, virImp.bvn)
	assert.Equal(t, live, virImp.live)
}

func Test_squadBaseACC_NewBussinessVirtualAcc_wrong_input_id(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	vAName := "joseph folayan"
	no := "08118995454"
	id := ""
	accNo := "1234567891"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "unique id must be passed")
}

func Test_squadBaseACC_NewBussinessVirtualAcc_wrong_input_name(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	vAName := ""
	no := "08118995454"
	id := "xyzr"
	accNo := "1234567891"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "bussiness name must be passed")
}

func Test_squadBaseACC_NewBussinessVirtualAcc_wrong_input_bvn(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	vAName := "james john"
	no := "08118995454"
	id := "xyzr"
	accNo := "1234567891"
	bvn := "1234567891"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid bvn format should be 11 digits")
}

func Test_squadBaseACC_NewBussinessVirtualAcc_wrong_input_phone_no(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	vAName := "james john"
	no := "081189954541234"
	id := "xyzr"
	accNo := "1234567891"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid phone no format")
}

func Test_squadBaseACC_NewBussinessVirtualAcc_wrong_input_acc_no(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	vAName := "james john"
	no := "09118995454"
	id := "xyzr"
	accNo := "1234567567891"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid bank account number format")
}
