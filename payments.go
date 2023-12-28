/*
 * squad-go - A simple wrapper written over squad apis payment endpoints.
 * this file contains the methods used to interact with the payment endpoints.
 * Author - Joseph Folayan
 * Github - joey1123455
 */

package squad

import (
	"errors"
	"math"
	"strings"

	"github.com/joey1123455/squad-go/utils"
)

// A interface exposing a private squad objects and its public methods.
type PaymentObject interface {
	Initiate(amount float64, currency, ref string, customer map[string]string, metaData any, reocure bool) (map[string]any, error)
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
 * NewPaymentObj - returns the pointer to a Payment Object that would be used to implement squad api functionalities.
 * @key string - A unique api key placed in the header of all squad requests.
 * @url string - The URL used to recieve transaction status webhooks.
 * @charge bool - a bool used to represent wether transaction fees are charged from the customer or merchant.
 * @live - a bool representing if the object is being used for tests or live transaction.
 * @chans - a string slice of payment channels accepted by the merchant.
 */
// TODO: create validation for input
func NewPaymentObj(key, url string, charge, live bool, chans []string) PaymentObject {
	return &paymentObjectImp{
		ApiKey:       key,
		CallBack:     url,
		PassCharge:   charge,
		PaymentChans: chans,
		live:         live,
	}
}

/*
 * Initiate - a method used to initiate the payment.
 * @amount - the float value to be charged.
 * @currency - the currency to be charged.
 * @ref - unique string used to refrence transactions. If == "" squad will provide
 * @customer - map containing customers name and email, (email must be passed).
 * @meta - an object of additional data used to track payment may be a string.
 * @reoccuring - a bool value to set if a card should always be charged i.e a subscription plan.
 */
func (p paymentObjectImp) Initiate(amount float64, currency, ref string, customer map[string]string, meta any, reoccuring bool) (map[string]any, error) {
	// todo: write a private method to verify input
	if customer == nil {
		return nil, errors.New("customer map must be passed")
	}
	email, emailOk := customer["email"]
	currency = strings.ToUpper(currency)
	switch {
	case currency != "NGN" && currency != "USD":
		return nil, errors.New("currency should be NGN or USD")
	case amount == 0:
		return nil, errors.New("amount is not provided")
	case customer == nil:
		return nil, errors.New("customer map must be provided")
	case !emailOk:
		return nil, errors.New("customer email must be provided")
	}

	body := map[string]any{
		"amount":           p.convert(amount),
		"email":            email,
		"initiate_type":    "inline",
		"currency":         currency,
		"is_recurring":     reoccuring,
		"callback_url":     p.CallBack,
		"pass_charge":      p.PassCharge,
		"payment_channels": p.PaymentChans,
	}
	if ref != "" {
		body["transaction_ref"] = ref
	}
	if customer["name"] != "" {
		body["customer_name"] = customer["name"]
	}
	if meta != nil {
		body["metadata"] = meta
	}

	res, err := utils.MakeRequest(body, p.completeUrl("transaction/initiate"), p.ApiKey)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/*
 * convert - convert from float to decimal
 * @amount - original value
 * returns - int value
 */
func (paymentObjectImp) convert(amount float64) int {
	return int(math.Round(amount * 100))
}

// TODO: refactore the complete url to be a function and not a reciever method

/*
 * completeUrl - returns the proper url for live and test objects
 * @endPoint - the endpoint to add to the base url
 * returns - the completed url
 */
func (p paymentObjectImp) completeUrl(endPoint string) string {
	const (
		testBase = "https://sandbox-api-d.squadco.com/"
		liveBase = "https://api-d.squadco.com/"
	)

	if p.live {
		return liveBase + endPoint
	}
	return testBase + endPoint
}
