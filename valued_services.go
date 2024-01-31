package squad

import (
	"errors"
	"strconv"
	"strings"

	"github.com/joey1123455/squad-go/utils"
)

const (
	vendAirtimeEndpoint        string = "vending/purchase/airtime"
	vendDataBundlesEndpoint    string = "vending/purchase/data"
	getDataBundlesEndpoint     string = "vending/data-bundles"
	transactionHistoryEndpoint string = "vending/transactions"
)

type ServicesClient interface {
	VendAirtime(phoneNumber string, amount int) (map[string]any, error)
	VendDataBundles(phoneNumber, planCode string, amount int) (map[string]any, error)
	GetDataBundles(network string) (map[string]any, error)
	TransactionHistory(page, perPage, action string) (map[string]any, error)
}

/*
 * VendAirtime - This API allows you vend airtime. Minimum amount that can be vended is 50 naira.
 * @phoneNumber - String 11 digit phone number. Format: "08139011943"
 * @amount - Integer Amount is in Naira.
 */
func (s *squadBaseACC) VendAirtime(phoneNumber string, amount int) (map[string]any, error) {
	switch {
	case amount < 1:
		return nil, errors.New("please provide an amount more then 0")
	case !utils.IsValidNigerianPhoneNumber(phoneNumber):
		return nil, errors.New("please provide a valid phone number")
	}
	data := map[string]any{
		"amount":       amount,
		"phone_number": phoneNumber,
	}
	return utils.MakeRequest(data, utils.CompleteUrl(vendAirtimeEndpoint, s.Live), s.ApiKey, "POST")
}

/*
 * VendDataBundles - This API allows you vend data bundles.
 * @phoneNumber - String 11 digit phone number. Format: : "08139011943"
 * @planCode - String The plan code is gotten from the Get Plan Code endpoint and usually in the format: "1001"
 * @amount - Integer Amount is in Naira.
 */
func (s *squadBaseACC) VendDataBundles(phoneNumber, planCode string, amount int) (map[string]any, error) {
	switch {
	case amount < 1:
		return nil, errors.New("please provide an amount more then 0")
	case !utils.IsValidNigerianPhoneNumber(phoneNumber):
		return nil, errors.New("please provide a valid phone number")
	}
	data := map[string]any{
		"phone_number": phoneNumber,
		"amount":       amount,
		"plan_code":    planCode,
	}
	return utils.MakeRequest(data, utils.CompleteUrl(vendDataBundlesEndpoint, s.Live), s.ApiKey, "POST")
}

/*
 * GetDataBundles - This API returns all available data bundle plans for all telcos
 * @network - String Teleco ID: MTN, GLO, AIRTEL, 9MOBILE
 */
func (s *squadBaseACC) GetDataBundles(network string) (map[string]any, error) {
	return utils.MakeGetRequest(map[string]string{
		"network": strings.ToUpper(network),
	}, utils.CompleteUrl(getDataBundlesEndpoint, s.Live), s.ApiKey)
}

/*
 * TransactionHistory - This API returns all transactions done by a merchant.
 * @page - string The page of the transaction you want to view
 * @perPage - string Number of transaction you want to view per page
 * @action - string The type of transaction you want to see: "debit"
 */
func (s *squadBaseACC) TransactionHistory(page, perPage, action string) (map[string]any, error) {

	if action != "credit" && action != "debit" {
		return nil, errors.New("action must be 'credit' or 'debit'")
	}

	querry := map[string]string{"action": action}

	if page != "" {
		valForPage, err := strconv.Atoi(page)
		if valForPage < 1 || err != nil {
			return nil, errors.New("page must be greater then 0")
		} else {
			querry["page"] = page
		}

	}
	if perPage != "" {
		valForPerPage, err1 := strconv.Atoi(perPage)
		if valForPerPage < 1 || err1 != nil {
			return nil, errors.New("perPage must be greater then 0")
		} else {
			querry["perPage"] = perPage
		}
	}
	return utils.MakeGetRequest(querry, utils.CompleteUrl(transactionHistoryEndpoint, s.Live), s.ApiKey)
}
