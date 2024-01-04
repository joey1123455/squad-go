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

const (
	initiatePaymentEndPoint string = "transaction/initiate"
	chargeCardEndpoint      string = "transaction/charge_card"
	verifyPaymentEndpoint   string = "transaction/verify/"
)

// A interface exposing a private squad objects and its public methods.
type PaymentObject interface {
	Initiate(amount float64, currency, ref string, customer map[string]string, metaData any, reocure bool) (map[string]any, error)
	ChargeCard(transactionRef, tokenId string, amount float64) (map[string]any, error)
	VerifyTransaction(transactionRef string) (map[string]any, error)
}

/*
 * The original squad object to be implemented as an interface.
 * @ApiKey - a string representing a squad account api key.
 * @CallBack - a string representing a url, where transaction status webkooks will be sent to.
 * @PassCharge - a boolean field representing if the payment charges while be charged from the
 * customer or merchant (false is default).
 * @PaymentChans - a string slice of payment channels accepted by the merchant.
 * @Live - a bool representing if the object is being used for tests or live transaction.
 */
type paymentObjectImp struct {
	ApiKey       string   `json:"api_key"`
	CallBack     string   `json:"callback_url"`
	PassCharge   bool     `json:"pass_charge"`
	PaymentChans []string `json:"payment_channels"`
	live         bool
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
func (p *paymentObjectImp) Initiate(amount float64, currency, ref string, customer map[string]string, meta any, reoccuring bool) (map[string]any, error) {
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

	res, err := utils.MakeRequest(body, utils.CompleteUrl(initiatePaymentEndPoint, p.live), p.ApiKey, post)
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
func (*paymentObjectImp) convert(amount float64) int {
	return int(math.Round(amount * 100))
}

// TODO: refactore the complete url to be a function and not a reciever method

/*
 * ChargeCard - This allows you charge a card using the token generated during the initial transaction which was sent via webhook
 * @transactionRef - Unique case-sensitive transaction reference. If you do not pass this parameter, Squad will generate a unique reference for you.
 * @tokenId - String A unique tokenization code for each card transaction and it is returned via the webhook for first charge on the card.
 * @amount - Integer Amount to charge from card in the lowest currency value. kobo for NGN transactions or cent for USD transactions
 */
func (p *paymentObjectImp) ChargeCard(transactionRef, tokenId string, amount float64) (map[string]any, error) {
	switch {
	case tokenId == "":
		return nil, errors.New("please provide the token id returned via the webhook for first charge on the card")
	case amount < 1:
		return nil, errors.New("amount can not be less then 1")
	case transactionRef == "":
		return nil, errors.New("please provide a transaction refrence")
	}
	body := map[string]any{
		"transaction_ref": transactionRef,
		"token_id":        tokenId,
		"amount":          p.convert(amount),
	}
	return utils.MakeRequest(body, utils.CompleteUrl(chargeCardEndpoint, p.live), p.ApiKey, "POST")
}

/*
 * VerifyTransaction - This is an endpoint that allows you to query the status of a particular transaction using the unique transaction reference attached to the transaction.
 * @transactionRef - String Unique transaction reference that identifies each transaction
 */
func (p *paymentObjectImp) VerifyTransaction(transactionRef string) (map[string]any, error) {
	switch {
	case transactionRef == "":
		return nil, errors.New("please provide a transaction ref")
	}
	return utils.MakeGetRequest(nil, utils.CompleteUrl(verifyPaymentEndpoint+transactionRef, p.live), p.ApiKey)
}
