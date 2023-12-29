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
