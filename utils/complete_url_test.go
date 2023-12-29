package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_paymentObjectImp_completeUrl_false(t *testing.T) {
	url := CompleteUrl("transaction/initiate", false)
	notExpected := "https://api-d.squadco.com/transaction/initiate"
	expected := "https://sandbox-api-d.squadco.com/transaction/initiate"
	assert.Equal(t, expected, url)
	assert.NotEqual(t, notExpected, url)
}

// go test -run Test_paymentObjectImp_completeUrl_true
func Test_paymentObjectImp_completeUrl_true(t *testing.T) {

	url := CompleteUrl("transaction/initiate", true)
	expected := "https://api-d.squadco.com/transaction/initiate"
	notExpected := "https://sandbox-api-d.squadco.com/transaction/initiate"
	assert.Equal(t, expected, url)
	assert.NotEqual(t, notExpected, url)
}
