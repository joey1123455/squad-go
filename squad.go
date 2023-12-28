package squad

import (
	"errors"
	"strings"

	"github.com/joey1123455/squad-go/utils"
)

// an interface exposing the methods of a squadimplementation
type SquadBaseAcc interface {
	CreatePaymentObject(charge bool, chans []string) PaymentObject
}

/*
 * squadBaseACC - an object representing a squad base account.
 * @ApiKey - a string representing your squad api key.
 * @accountName - a string representing the registered squad account name.
 * @Live - a bool representing if its a test or live object.
 * @CallBack - a string representing a url, where transaction status webkooks will be sent to.
 */
type squadBaseACC struct {
	ApiKey      string
	AccountName string
	Live        bool
	CallBack    string
}

/*
 * NewPaymentObj - returns the pointer to a Payment Object that would be used to implement squad api functionalities.
 * @key string - A unique api key placed in the header of all squad requests.
 * @url string - The URL used to recieve transaction status webhooks.
 * @name - a string representing the registered squad account name.
 * @live - a bool representing if the object is being used for tests or live transaction.
 */
// TODO: create validation for input
func NewSquadObj(key, url, name string, live bool) (SquadBaseAcc, error) {

	// input validation
	switch {
	case name == "":
		return nil, errors.New("please provide a bussiness name")
	case !utils.IsValidURL(url):
		return nil, errors.New("invalid callback url, should have a valid protocol 'http://' or 'https://'")
	case key == "":
		return nil, errors.New("api key must be provided")
	case !live && !strings.HasPrefix(key, "sandbox_sk"):
		return nil, errors.New("api key for test account must start with 'sandbox_sk'")
	case live && !strings.HasPrefix(key, "sk"):
		return nil, errors.New("api key for account must start with 'sk'")
	}

	return &squadBaseACC{
		ApiKey:      key,
		CallBack:    url,
		AccountName: name,
		Live:        live,
	}, nil
}

/*
 * CreatePaymentObject - returns the pointer to a Payment Object that would be used to implement squad api functionalities.
 * @charge bool - a bool used to represent wether transaction fees are charged from the customer or merchant.
 * @chans - a string slice of payment channels accepted by the merchant.
 */
func (sba squadBaseACC) CreatePaymentObject(charge bool, chans []string) PaymentObject {
	return &paymentObjectImp{
		ApiKey:       sba.ApiKey,
		CallBack:     sba.CallBack,
		PassCharge:   charge,
		PaymentChans: chans,
		live:         sba.Live,
	}
}
