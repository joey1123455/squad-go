package squad

import (
	"github.com/joey1123455/squad-go/utils"
	"golang.org/x/exp/slices"
)

const (
	makeSubMerchantEndpoint string = "merchant/create-sub-users"
	walletBalanceEndpoint   string = "merchant/balance"
	refundEndpoint          string = "transaction/refund"
)

type SquadUtilClient interface {
	CreateSubMerchant(customerData map[string]any) (map[string]any, error)
	WalletBalance() (map[string]any, error)
	Refund(data map[string]any) (map[string]any, error)
}

/*
 * CreateSubMerchant - This API is used to create a sub-merchant, the sub-merchant will have its own ID and will automatically have its own view on the dashboard.
 * @customerData - map[string]string sub merchants data required for creation {
 * 		@display_name - String Name of sub-merchant
 * 		@account_name - String Sub-merchant's settlement bank account name
 * 		@account_number -String Sub-merchant's settlement account number
 * 		@bank_code - String Sub-merchant's settlement bank code. e.g 058
 * 		@bank - String Name of sub-merchant's settlement bank e.g GTBank
 * }
 */
func (s *squadBaseACC) CreateSubMerchant(customerData map[string]any) (map[string]any, error) {
	options := []string{
		"display_name", "account_name", "account_number",
		"bank_code", "bank",
	}
	for k, v := range customerData {
		switch {
		case !slices.Contains(options, k):
			delete(customerData, k)
		case v == "":
			delete(customerData, k)
		}
	}
	return utils.MakeRequest(customerData, utils.CompleteUrl(makeSubMerchantEndpoint, s.Live), s.ApiKey, "POST")
}

/*
 * WalletBalance - This endpoint allows you get your Squad Wallet Balance. Please be informed that the wallet balance is in KOBO. (Please note that you can't get wallet balance for Dollar transactions)
 */
func (sba *squadBaseACC) WalletBalance() (map[string]any, error) {
	return utils.MakeGetRequest(map[string]string{"currency_id": "NGN"}, utils.CompleteUrl(walletBalanceEndpoint, sba.Live), sba.ApiKey)
}

/*
 * Refund - This API is used to initiate refund process on a successful transaction.
 * @data - map[string]any the required data for payment refund {
 *		@gateway_transaction_ref - String Unique reference that uniquely identifies the medium of payment and can be obtained from the webhook notification sent to you.
 *		@transaction_ref - String unique reference that identifies a transaction. Can be obtained from the dashboard or the webhook notification sent to you
 *		@refund_type - String The value of this parameter is either "Full" or "Partial"
 *		@reason_for_refund - String Reason for initiating the refund
 *		@refund_amount - String Refund amount is in kobo or cent. This is only required for "Partial" refunds
 * }
 */
func (sba *squadBaseACC) Refund(data map[string]any) (map[string]any, error) {
	options := []string{
		"gateway_transaction_ref", "transaction_ref", "refund_type",
		"reason_for_refund", "refund_amount",
	}
	for k, v := range data {
		switch {
		case !slices.Contains(options, k):
			delete(data, k)
		case v == "":
			delete(data, k)
		}
	}
	return utils.MakeRequest(data, utils.CompleteUrl(refundEndpoint, sba.Live), sba.ApiKey, "POST")
}
