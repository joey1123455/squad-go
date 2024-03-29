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
	assert.Nil(t, err)
	virAcc, err := squad.NewBussinessVirtualAcc(id, vAName, no, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)

	virImp := virAcc.(*bussinessVA)
	assert.Equal(t, id, virImp.CustomerID)
	assert.Equal(t, vAName, virImp.BussinessName)
	assert.Equal(t, name, virImp.AccountName)
	assert.Equal(t, apiKey, virImp.ApiKey)
	assert.Equal(t, no, virImp.MobileNo)
	assert.Equal(t, accNo, virImp.BeneficiaryAcc)
	assert.Equal(t, bvn, virImp.Bvn)
	assert.Equal(t, live, virImp.Live)
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

func Test_NewCustomerVirtualAcc_success(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "folayan"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn bauchi"
	gender := "1"
	no := "08118995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, err)
	assert.NotNil(t, virAcc)

	virImp := virAcc.(*customerVA)
	assert.Equal(t, id, virImp.CustomerID)
	assert.Equal(t, first, virImp.FirstName)
	assert.Equal(t, last, virImp.LastName)
	assert.Equal(t, email, virImp.Email)
	assert.Equal(t, dob, virImp.Dob)
	assert.Equal(t, address, virImp.Address)
	assert.Equal(t, gender, virImp.Gender)
	assert.Equal(t, name, virImp.AccountName)
	assert.Equal(t, apiKey, virImp.ApiKey)
	assert.Equal(t, no, virImp.MobileNo)
	assert.Equal(t, accNo, virImp.BeneficiaryAcc)
	assert.Equal(t, bvn, virImp.Bvn)
	assert.Equal(t, live, virImp.Live)
}

func Test_NewCustomerVirtualAcc_wrong_input_first_name(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := ""
	last := "folayan"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn bauchi"
	gender := "1"
	no := "08118995454"
	id := "hex11rthyuirjahdu"
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "first name must be passed")
}

func Test_NewCustomerVirtualAcc_wrong_input_customer_id(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "folayan"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn bauchi"
	gender := "1"
	no := "08118995454"
	id := ""
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "unique id must be passed")
}

func Test_NewCustomerVirtualAcc_wrong_input_last_name(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := ""
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn bauchi"
	gender := "1"
	no := "08118995454"
	id := "xyztyr2"
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "last name must be passed")
}

func Test_NewCustomerVirtualAcc_wrong_input_mobile_no(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "james"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn bauchi"
	gender := "1"
	no := "0811899545"
	id := "xyztyr2"
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid phone no format")
}

func Test_NewCustomerVirtualAcc_wrong_input_email(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "james"
	email := "joeyfolayan5@.com"
	dob := "10/09/2000"
	address := "village inn bauchi"
	gender := "1"
	no := "08118995445"
	id := "xyztyr2"
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid email address")
}

func Test_NewCustomerVirtualAcc_wrong_input_address(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "james"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := ""
	gender := "1"
	no := "08118995445"
	id := "xyztyr2"
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "please provide customer address")
}

func Test_NewCustomerVirtualAcc_wrong_input_gender(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "james"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn"
	gender := "3"
	no := "08118995445"
	id := "xyztyr2"
	accNo := "1234567291"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "gender should be '1' for male or '2' for female")
}

func Test_NewCustomerVirtualAcc_wrong_input_acc_no(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "james"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn"
	gender := "2"
	no := "08118995445"
	id := "xyztyr2"
	accNo := "12345678891"
	bvn := "12345678911"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid bank account number format")
}

func Test_NewCustomerVirtualAcc_wrong_input_bvn(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	first := "joseph"
	last := "james"
	email := "joeyfolayan5@gmail.com"
	dob := "10/09/2000"
	address := "village inn"
	gender := "2"
	no := "08118995445"
	id := "xyztyr2"
	accNo := "1234567891"
	bvn := "123456789119"
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	virAcc, err := squad.NewCustomerVirtualAcc(id, first, last, no, email, dob, address, gender, accNo, bvn)
	assert.Nil(t, virAcc)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid bvn format should be 11 digits")
}

func Test_squadBaseACC_NewTransferClient(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	transferClient := squad.NewTransferClient()
	assert.NotNil(t, transferClient)
}

func Test_squadBaseACC_NewUtilClient(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	transferClient := squad.NewUtilClient()
	assert.NotNil(t, transferClient)
}

func Test_squadBaseACC_ServicesClient(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := "https://calback/correct.com"
	name := "test bussines"
	live := false
	squad, err := NewSquadObj(apiKey, url, name, live)
	assert.Nil(t, err)
	transferClient := squad.NewServicesClient()
	assert.NotNil(t, transferClient)
}
