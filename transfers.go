package squad

import (
	"errors"

	"github.com/joey1123455/squad-go/utils"
)

const (
	accountLookupEndpoint string = "payout/account/lookup"
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
