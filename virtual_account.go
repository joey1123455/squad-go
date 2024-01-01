/*
 * squad-go - A simple wrapper written over squad apis payment endpoints.
 * this file contains the methods used to interact with the virtual account endpoints.
 * Author - Joseph Folayan
 * Github - joey1123455
 */

package squad

import (
	"errors"
	"fmt"
	"strings"

	"github.com/joey1123455/squad-go/utils"
	"golang.org/x/exp/slices"
)

const (
	// methods
	post  string = "POST"
	patch string = "PATCH"

	// endpoints
	getWebhookEndpoint                      string = "virtual-account/webhook/logs"
	createBusinessVAEndpoint                string = "virtual-account/business"
	createCustomerVAEndpoint                string = "virtual-account"
	queryVirtualAccHistoryEndpoint          string = "virtual-account/customer/transactions/"
	queryMerchantHistoryEndpoint            string = "virtual-account/merchant/transactions"
	queryMerchantHistoryWithFiltersEndpoint string = "virtual-account/merchant/transactions/all"
	getMerchantVirtualAccEndpoint           string = "virtual-account/merchant/accounts"
	simultePaymentEndpoint                  string = "virtual-account/simulate/payment"
	virtualAccDetailsEndpoint               string = "virtual-account/customer/"
	virtualAccDetailsUsingIdEndpoint        string = "virtual-account/"
	updateAccountEndpoint                   string = "virtual-account/update/beneficiary/account"
	simPaymentEndpoint                      string = "virtual-account/simulate/payment"
)

// an interface exposing virtual accounts of either customer or bussiness model
type VirtualAccount interface {
	Initiate() (map[string]any, error)
	QueryVirtualAccHistory() (map[string]any, error)
	AccountDetails() (map[string]any, error)
	AccountDetailsUsingId() (map[string]any, error)
	UpdateAccount(beneficiaryAccount string) (map[string]any, error)
	SimulatePayment(amount int) (map[string]any, error)
}

/*
 * bussinessVA - a struct representing a virtual account for a bussiness model.
 * @customerID - a unique string representing the bussiness id.
 * @bussinessName - a string representings the bussiness name. all virtual accounts must carry a slug as a prefix
 * to the name. The slug must be a portion of the bussiness name or abbreviations of your business name used to open
 * the squad account as one word. Please note that slash (/) is not allowed and only hyphen (-) can be used.
 * @mobileNo - customers mobile no
 * @beneficiaryAcc - customers account no
 * @apiKey - string representing api key
 * @accountName - string representing the original squad account name
 * @bvn - string representing the bussiness bvn associated to the provided bvn
 * @live - a bool representing if the object is being used for tests or live transaction.
 * @virtualAccNo - the virtual account no provided by squad
 */
type bussinessVA struct {
	CustomerID     string
	BussinessName  string
	MobileNo       string
	BeneficiaryAcc string
	ApiKey         string
	AccountName    string
	Bvn            string
	Live           bool
	VirtualAccNo   string
}

/*
 * customerVA - a struct representing a virtual account for a customer model.
 * @customerID - a unique string representing the customer id.
 * @firstName - string representing customers first name.
 * @lastName - string representing customers last name.
 * @mobileNo - customers mobile no.
 * @email - customers email address.
 * @bvn - customers bvn.
 * @dob - customers date of birth.
 * @address - customers address.
 * @gender - customers gender 1 for male, 2 for female pass as a string.
 * @beneficiaryAcc - customers account no
 * @apiKey - string representing api key
 * @accountName - string representing the original squad account name
 * @live - a bool representing if the object is being used for tests or live transaction.
 * @virtualAccNo - the virtual account no provided by squad
 */
type customerVA struct {
	CustomerID     string
	FirstName      string
	LastName       string
	MobileNo       string
	Email          string
	Bvn            string
	Dob            string
	Address        string
	Gender         string
	BeneficiaryAcc string
	ApiKey         string
	AccountName    string
	Live           bool
	VirtualAccNo   string
}

/*
 * TestPaymentToVA - used to simulate a payment to a virtual account for testing
 * @apiKey - string representing api key
 * @acc - account number.
 * @amount - simulated payment amount.
 * @live - a bool representing if the object is being used for tests or live transaction.
 */
func TestPaymentToVA(apiKey, acc string, amount float64, live bool) (map[string]any, error) {
	switch {
	case !utils.IsValidNigerianAccountNumber(acc):
		return nil, errors.New("invalid account no")
	case amount < 1:
		return nil, errors.New("amount must be greater then 0")
	case !live && !strings.HasPrefix(apiKey, "sandbox_sk"):
		return nil, errors.New("api key for test account must start with 'sandbox_sk'")
	case live && !strings.HasPrefix(apiKey, "sk"):
		return nil, errors.New("api key for account must start with 'sk'")
	}
	body := map[string]any{
		"virtual_account_number": acc,
		"amount":                 fmt.Sprint(amount),
	}
	return utils.MakeRequest(body, utils.CompleteUrl(simultePaymentEndpoint, live), apiKey, post)
}

/*
 * MissedWebHookNotifications - used to get a list of missed transaction notifications to virtual accounts.
 * @apiKey - string representing api key
 * @live - a bool representing if the object is being used for tests or live transaction.
 * @page - the page to return.
 * @perPage - amount to return per page.
 */
func MissedWebHookNotifications(apiKey string, live bool, page, perPage int) (map[string]any, error) {
	// input validation
	switch {
	case page < 1:
		return nil, errors.New("page must have a value greater then 0")
	case perPage < 1:
		return nil, errors.New("perPage must have a value greater then 0")
	case !live && !strings.HasPrefix(apiKey, "sandbox_sk"):
		return nil, errors.New("api key for test account must start with 'sandbox_sk'")
	case live && !strings.HasPrefix(apiKey, "sk"):
		return nil, errors.New("api key for account must start with 'sk'")
	}

	return utils.MakeGetRequest(map[string]string{"page": fmt.Sprint(page), "perPage": fmt.Sprint(perPage)}, utils.CompleteUrl(getWebhookEndpoint, live), apiKey)
}

// TODO: implement delete missed web hook last nah zxenox go run this one.
/*
 * DeleteMissedWebHook - handles deletion of missed web hooks
 */

/*
 * QueryMerchantHistory - used to get all virtual account transactions registered under an api key
 * @apiKey - merchants api key
 * @live - bool to represent if a live or sandbox account
 */
func QueryMerchantVirtualAccHistory(apiKey string, live bool) (map[string]any, error) {
	// input validation
	switch {
	case !live && !strings.HasPrefix(apiKey, "sandbox_sk"):
		return nil, errors.New("api key for test account must start with 'sandbox_sk'")
	case live && !strings.HasPrefix(apiKey, "sk"):
		return nil, errors.New("api key for account must start with 'sk'")
	}
	return utils.MakeGetRequest(nil, utils.CompleteUrl(queryMerchantHistoryEndpoint, live), apiKey)
}

/*
 * QueryMerchantHistoryFilters - used to get all virtual account transactions registered under an api key with the use of filters
 * @apiKey - merchants api key
 * @live - bool to represent if a live or sandbox account
 * @filters - the filters to use {
 * 		@page - the page to return.
 * 		@perPage - amount to return per page.
 * 		@virtualAccount - used to return on values from this account no
 * 		@customerIdentifier - used to return values from this customer id
 * 		@startDate - used to specify which date transactions should start from
 * 		@endDate - used to specify which date transactions should stop
 * 		@transactionReference - used to return a transaction ref
 * 		@session_id - used to return all NIP transactions
 * 		@dir - takes to values return in "ASC" or "DESC"
 *	}
 */
func QueryMerchantHistoryFilters(apiKey string, live bool, filters map[string]string) (map[string]any, error) {
	options := []string{
		"page", "perPage", "dir", "virtualAccount",
		"customerIdentifier", "startDate", "endDate",
		"transactionReference", "session_id",
	}
	for k, v := range filters {
		switch {
		case !slices.Contains(options, k):
			delete(filters, k)
		case v == "":
			delete(filters, k)
		case k == "dir":
			filters[k] = strings.ToUpper(v)
		}
	}

	//input validation
	switch {
	case !live && !strings.HasPrefix(apiKey, "sandbox_sk"):
		return nil, errors.New("api key for test account must start with 'sandbox_sk'")
	case live && !strings.HasPrefix(apiKey, "sk"):
		return nil, errors.New("api key for account must start with 'sk'")
	case len(filters) == 0:
		return nil, errors.New("if no filters are provided make use of the QueryMerchantHistory function instead")
	}
	return utils.MakeGetRequest(filters, utils.CompleteUrl(queryMerchantHistoryWithFiltersEndpoint, live), apiKey)
}

/*
 * GetMerchantVirtualAcc - returns a list of all merchants virtual accounts
 * @apiKey - merchants api key
 * @live - bool to represent if a live or sandbox account
 */
func GetMerchantVirtualAcc(apiKey string, live bool) (map[string]any, error) {
	switch {
	case !live && !strings.HasPrefix(apiKey, "sandbox_sk"):
		return nil, errors.New("api key for test account must start with 'sandbox_sk'")
	case live && !strings.HasPrefix(apiKey, "sk"):
		return nil, errors.New("api key for account must start with 'sk'")
	}
	return utils.MakeGetRequest(nil, utils.CompleteUrl(getMerchantVirtualAccEndpoint, live), apiKey)
}

// *bussinessVA reciever methods

/*
 * Initiate - makes a request to the create business virtual accounts end point
 */
func (bv *bussinessVA) Initiate() (res map[string]any, err error) {
	body := map[string]any{
		"bvn":                 bv.Bvn,
		"business_name":       bv.BussinessName,
		"customer_identifier": bv.CustomerID,
		"mobile_num":          bv.MobileNo,
		"beneficiary_account": bv.BeneficiaryAcc,
	}
	res, err = utils.MakeRequest(body, utils.CompleteUrl(createBusinessVAEndpoint, bv.Live), bv.ApiKey, post)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	data, ok := res["data"].(map[string]any)
	if !ok {
		return
	}
	virtualAccountNumber, ok := data["virtual_account_number"].(string)
	if !ok {
		return
	}
	bv.VirtualAccNo = virtualAccountNumber

	return
}

/*
 * QueryVirtualAccHistory - gets a list of all transactions by a virtual account
 */
func (bv *bussinessVA) QueryVirtualAccHistory() (map[string]any, error) {
	return utils.MakeGetRequest(nil, utils.CompleteUrl(queryVirtualAccHistoryEndpoint+bv.CustomerID, bv.Live), bv.ApiKey)
}

/*
 * AccountDetails - returns virtual account details
 */
func (bv *bussinessVA) AccountDetails() (map[string]any, error) {
	return utils.MakeGetRequest(nil, utils.CompleteUrl(virtualAccDetailsEndpoint+bv.VirtualAccNo, bv.Live), bv.ApiKey)
}

/*
 * AccountDetailsUsingId - returns virtual account details
 */
func (bv *bussinessVA) AccountDetailsUsingId() (map[string]any, error) {
	return utils.MakeGetRequest(nil, utils.CompleteUrl(virtualAccDetailsUsingIdEndpoint+bv.CustomerID, bv.Live), bv.ApiKey)
}

/*
 * UpdateAccount - updates the original account for a virtual account
 * @beneficiaryAccount - 10 digit valid NUBAN account number
 */
func (bv *bussinessVA) UpdateAccount(beneficiaryAccount string) (map[string]any, error) {
	switch {
	case !utils.IsValidNigerianAccountNumber(beneficiaryAccount):
		return nil, errors.New("invalid account no")
	}
	body := map[string]any{
		"beneficiary_account":    beneficiaryAccount,
		"virtual_account_number": bv.VirtualAccNo,
	}
	return utils.MakeRequest(body, utils.CompleteUrl(updateAccountEndpoint, bv.Live), bv.ApiKey, patch)
}

/*
 * SimulatePayment - simulates a payment to a virtual account no
 * @amount - the amount to be paid
 */
func (bv *bussinessVA) SimulatePayment(amount int) (map[string]any, error) {
	body := map[string]any{
		"virtual_account_number": bv.VirtualAccNo,
		"amount":                 fmt.Sprint(amount),
	}
	return utils.MakeRequest(body, utils.CompleteUrl(simPaymentEndpoint, bv.Live), bv.ApiKey, post)
}

// *customerVA reciever methods

/*
 * Initiate - makes a request to the create customer virtual accounts end point
 */
func (cv *customerVA) Initiate() (res map[string]any, err error) {
	body := map[string]any{
		"bvn":                 cv.Bvn,
		"first_name":          cv.FirstName,
		"customer_identifier": cv.CustomerID,
		"last_name":           cv.LastName,
		"mobile_num":          cv.MobileNo,
		"email":               cv.Email,
		"dob":                 cv.Dob,
		"address":             cv.Address,
		"gender":              cv.Gender,
		"beneficiary_account": cv.BeneficiaryAcc,
	}
	res, err = utils.MakeRequest(body, utils.CompleteUrl(createCustomerVAEndpoint, cv.Live), cv.ApiKey, post)
	if err != nil {
		return
	}
	data, ok := res["data"].(map[string]any)
	if !ok {
		return
	}
	virtualAccountNumber, ok := data["virtual_account_number"].(string)
	if !ok {
		return
	}
	cv.VirtualAccNo = virtualAccountNumber

	return
}

/*
 * QueryVirtualAccHistory - gets a list of all transactions by a virtual account
 */
func (cv *customerVA) QueryVirtualAccHistory() (map[string]any, error) {

	return utils.MakeGetRequest(nil, utils.CompleteUrl(queryVirtualAccHistoryEndpoint+cv.CustomerID, cv.Live), cv.ApiKey)
}

/*
 * AccountDetails - returns virtual account details
 */
func (cv *customerVA) AccountDetails() (map[string]any, error) {
	return utils.MakeGetRequest(nil, utils.CompleteUrl(virtualAccDetailsEndpoint+cv.VirtualAccNo, cv.Live), cv.ApiKey)
}

/*
 * AccountDetailsUsingId - returns virtual account details
 */
func (cv *customerVA) AccountDetailsUsingId() (map[string]any, error) {
	return utils.MakeGetRequest(nil, utils.CompleteUrl(virtualAccDetailsUsingIdEndpoint+cv.CustomerID, cv.Live), cv.ApiKey)
}

/*
 * UpdateAccount - updates the original account for a virtual account
 * @beneficiaryAccount - 10 digit valid NUBAN account number
 */
func (cv *customerVA) UpdateAccount(beneficiaryAccount string) (map[string]any, error) {
	switch {
	case !utils.IsValidNigerianAccountNumber(beneficiaryAccount):
		return nil, errors.New("invalid account no")
	}
	body := map[string]any{
		"beneficiary_account":    beneficiaryAccount,
		"virtual_account_number": cv.VirtualAccNo,
	}
	return utils.MakeRequest(body, utils.CompleteUrl(updateAccountEndpoint, cv.Live), cv.ApiKey, patch)
}

/*
 * SimulatePayment - simulates a payment to a virtual account no
 * @amount - the amount to be paid
 */
func (cv *customerVA) SimulatePayment(amount int) (map[string]any, error) {

	body := map[string]any{
		"virtual_account_number": cv.VirtualAccNo,
		"amount":                 fmt.Sprint(amount),
	}
	return utils.MakeRequest(body, utils.CompleteUrl(simPaymentEndpoint, cv.Live), cv.ApiKey, post)
}
