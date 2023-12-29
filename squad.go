package squad

import (
	"errors"
	"strings"

	"github.com/joey1123455/squad-go/utils"
)

// an interface exposing the methods of a squad object implementation
type SquadBaseAcc interface {
	CreatePaymentObject(charge bool, chans []string) PaymentObject
	NewBussinessVirtualAcc(id, name, no, acc, bvn string) (VirtualAccount, error)
	NewCustomerVirtualAcc(customerID, firstName, lastName, mobileNo, email, dob, address, gender, beneficiaryAcc, bvn string) (VirtualAccount, error)
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
 * NewSquadObj - returns the pointer to a Squad Object that would be used to create other squad objects.
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
	// TODO: implement validation for the payment channel slice
	return &paymentObjectImp{
		ApiKey:       sba.ApiKey,
		CallBack:     sba.CallBack,
		PassCharge:   charge,
		PaymentChans: chans,
		live:         sba.Live,
	}
}

/*
 * NewBussinessVirtualAcc - function that returns a bussines virtual account.
 * @id - a unique string representing the bussiness id.
 * @name - a string representings the bussiness name. all virtual accounts must carry a slug as a prefix
 * to the name. The slug must be a portion of the bussiness name or abbreviations of your business name used to open
 * the squad account as one word. Please note that slash (/) is not allowed and only hyphen (-) can be used.
 * @acc - string representing a users original account no.
 * @no - string representing a users phone number.
 * @bvn - string representing the bussiness bvn associated to the provided bvn
 */
func (sba squadBaseACC) NewBussinessVirtualAcc(id, name, no, acc, bvn string) (VirtualAccount, error) {
	// input validation
	switch {
	case id == "":
		return nil, errors.New("unique id must be passed")
	case name == "":
		return nil, errors.New("bussiness name must be passed")
	case !utils.IsValidNigerianBVN(bvn):
		return nil, errors.New("invalid bvn format should be 11 digits")
	case !utils.IsValidNigerianPhoneNumber(no):
		return nil, errors.New("invalid phone no format")
	case !utils.IsValidNigerianAccountNumber(acc):
		return nil, errors.New("invalid bank account number format")
	}

	return &bussinessVA{
		customerID:     id,
		bussinessName:  sba.parseVirtualAccName(name),
		mobileNo:       no,
		apiKey:         sba.ApiKey,
		bvn:            bvn,
		beneficiaryAcc: acc,
		accountName:    sba.AccountName,
		live:           sba.Live,
	}, nil
}

/*
 * NewCustomerVirtualAcc - function that returns a customer virtual account.
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
 * @bvn - string representing the bussiness bvn associated to the provided bvn
 */
func (sba squadBaseACC) NewCustomerVirtualAcc(customerID, firstName, lastName, mobileNo, email, dob, address, gender, beneficiaryAcc, bvn string) (VirtualAccount, error) {
	// input validation
	switch {
	case customerID == "":
		return nil, errors.New("unique id must be passed")
	case firstName == "":
		return nil, errors.New("first name must be passed")
	case lastName == "":
		return nil, errors.New("last name must be passed")
	case !utils.IsValidNigerianPhoneNumber(mobileNo):
		return nil, errors.New("invalid phone no format")
	case !utils.IsValidEmail(email):
		return nil, errors.New("invalid email address")
	case !utils.IsValidDateOfBirth(dob):
		return nil, errors.New("invalid date of birth")
	case address == "":
		return nil, errors.New("please provide customer address")
	case gender != "1" && gender != "2":
		return nil, errors.New("gender should be '1' for male or '2' for female")
	case !utils.IsValidNigerianAccountNumber(beneficiaryAcc):
		return nil, errors.New("invalid bank account number format")
	case !utils.IsValidNigerianBVN(bvn):
		return nil, errors.New("invalid bvn format should be 11 digits")
	}

	return &customerVA{
		customerID:     customerID,
		firstName:      sba.parseVirtualAccName(firstName),
		lastName:       lastName,
		mobileNo:       mobileNo,
		email:          email,
		bvn:            bvn,
		dob:            dob,
		address:        address,
		gender:         gender,
		beneficiaryAcc: beneficiaryAcc,
		apiKey:         sba.ApiKey,
		accountName:    sba.AccountName,
		live:           sba.Live,
	}, nil

}

/*
 * parsedVirtualAccName - combines the customer name with the bussines name for virtual account creation
 * @customerName - string representing customers name for a virtual account
 */
func (sba squadBaseACC) parseVirtualAccName(customerName string) string {
	parts := strings.Split(sba.AccountName, " ")
	return parts[0] + "-" + customerName
}
