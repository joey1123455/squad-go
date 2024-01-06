package squad

import (
	"errors"
	"strings"

	"github.com/joey1123455/squad-go/utils"
	"golang.org/x/exp/slices"
)

const (
	accountLookupEndpoint   string = "payout/account/lookup"
	transferFundsEndpoint   string = "payout/transfer"
	getAllTransfersEndpoint string = "payout/list"
)

/*
  - @AccountLookup - This API allows you lookup/confirm the account name of the recipient you intend to credit. This should be done before initiating the transfer.
  - @bankCode - String Unique NIP code that identifies a bank.

 * @accountNumber - String Account number you want to transfer to
*/
func (sba *squadBaseACC) AccountLookup(bankCode, accountNumber string) (map[string]any, error) {
	switch {
	case !utils.IsValidNigerianAccountNumber(accountNumber):
		return nil, errors.New("invalid account no")
	case bankCode == "":
		return nil, errors.New("please provide a bank code")
	}
	body := map[string]any{
		"bank_code":      bankCode,
		"account_number": accountNumber,
	}

	return utils.MakeRequest(body, utils.CompleteUrl(accountLookupEndpoint, sba.Live), sba.ApiKey, "POST")
}

/*
 * FundTransfer - This API allows you to transfer funds from your Squad Wallet to the account you have looked up. Please be informed that squad will not be held liable for mistake in transferring to a wrong account or an account that wasn't looked up.
 * @transactionData - map[string]any required data for the transaction {
 *		@transaction_reference - String Unique Transaction Reference used to initiate a transfer. Please ensure that you append your merchant ID to the transaction Reference you are creating. This is compulsory as it will throw an error if you don't append it.
 * 		@amount - String Amount to be transferred. Value is in Kobo.
 * 		@bank_code - String Unique NIP Code that identifies a bank.
 * 		@account_number - String 10-digit NUBAN account number to be transferred to. Must be an account that has been looked up and vetted to be transferred to.
 * 		@account_name - String The account name tied to the account number you are transferring to which you have looked up using our look up API.
 * 		@currency_id - String Takes only the value "NGN"
 *		@remark - String A unique remark that will be sent with the transfer.
 * }
 */
func (sba *squadBaseACC) FundTransfer(transactionData map[string]any) (map[string]any, error) {
	options := []string{
		"transaction_reference", "account_name", "account_number",
		"bank_code", "amount", "currency_id", "remark",
	}
	for k, v := range transactionData {
		switch {
		case !slices.Contains(options, k):
			delete(transactionData, k)
		case v == "":
			delete(transactionData, k)
		}
	}
	return utils.MakeRequest(transactionData, utils.CompleteUrl(transferFundsEndpoint, sba.Live), sba.ApiKey, "POST")
}

/*
 * GetAllTransfers - This API Allows you retrieve the details of all transfers you have done from your Squad Wallet using this transfer solution.
 * @page - String Number of Pages
 * @perPage - String Number of records per page
 * @dir - String Allows you sort the records in either ascending or descending order. It takes the value "ASC" or "DESC"
 */
func (sba *squadBaseACC) GetAllTransfers(page, perPage, dir string) (map[string]any, error) {
	querries := make(map[string]string)
	if dir != "" {
		dir = strings.ToUpper(dir)
	}
	switch {
	case page != "":
		querries["page"] = page
	case perPage != "":
		querries["perPage"] = perPage
	case dir == "ASC" || dir == "DESC":
		querries["dir"] = dir
	}
	return utils.MakeGetRequest(querries, utils.CompleteUrl(getAllTransfersEndpoint, sba.Live), sba.ApiKey)
}
