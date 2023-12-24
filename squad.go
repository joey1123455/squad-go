/*
 * squad-go - A simple wrapper written over squad apis payment endpoints.
 * this file contains the object required by the squad endpoints and function to initiate it.
 * Author - Joseph Folayan
 * Github - joey1123455
 */

package squad

// A interface exposing a private squad objects and its public methods.
type PaymentObject interface {
	Initiate(amount float64, currency, ref string, customer map[string]string, metaData ...any) (apiRes, error)
}

/*
  - The original squad object to be implemented as an interface.
  - @ApiKey - a string representing a squad account api key.
  - @CallBack - a string representing a url, where transaction status webkooks will be sent to.
  - @PassCharge - a boolean field representing if the payment charges while be charged from the
    customer or merchant (false is default).
  - @PaymentChans - a string slice of payment channels accepted by the merchant.
  - @Live - a bool representing if the object is being used for tests or live transaction.
*/
type paymentObjectImp struct {
	ApiKey       string   `json:"api_key"`
	CallBack     string   `json:"callback_url"`
	PassCharge   bool     `json:"pass_charge"`
	PaymentChans []string `json:"payment_channels"`
	live         bool
}

/*
 * NewSquadObj - returns the pointer to a Payment Object that would be used to implement squad api functionalities.
 * @key string - A unique api key placed in the header of all squad requests.
 * @url string - The URL used to recieve transaction status webhooks.
 * @charge bool - a bool used to represent wether transaction fees are charged from the customer or merchant.
 * @live - a bool representing if the object is being used for tests or live transaction.
 * @chans - a string slice of payment channels accepted by the merchant.
 */
func NewSquadObj(key, url string, charge, live bool, chans []string) PaymentObject {
	return &paymentObjectImp{
		ApiKey:       key,
		CallBack:     url,
		PassCharge:   charge,
		PaymentChans: chans,
		live:         live,
	}
}
