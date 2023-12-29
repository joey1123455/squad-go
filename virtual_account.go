/*
 * squad-go - A simple wrapper written over squad apis payment endpoints.
 * this file contains the methods used to interact with the virtual account endpoints.
 * Author - Joseph Folayan
 * Github - joey1123455
 */

package squad

import "github.com/joey1123455/squad-go/utils"

// an interface exposing virtual accounts of either customer or bussiness model
type VirtualAccount interface {
	Initiate() (map[string]any, error)
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
*/
type bussinessVA struct {
	customerID     string
	bussinessName  string
	mobileNo       string
	beneficiaryAcc string
	apiKey         string
	accountName    string
	bvn            string
	live           bool
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
 */
type customerVA struct {
	customerID     string
	firstName      string
	lastName       string
	mobileNo       string
	email          string
	bvn            string
	dob            string
	address        string
	gender         string
	beneficiaryAcc string
	apiKey         string
	accountName    string
	live           bool
}

/*
 * Initiate - makes a request to the create virtual accounts end point
 */
func (bv bussinessVA) Initiate() (map[string]any, error) {
	body := map[string]any{
		"bvn":                 bv.bvn,
		"business_name":       bv.bussinessName,
		"customer_identifier": bv.customerID,
		"mobile_num":          bv.mobileNo,
		"beneficiary_account": bv.beneficiaryAcc,
	}
	return utils.MakeRequest(body, utils.CompleteUrl("virtual-account/business", bv.live), bv.apiKey)
}

func (cv customerVA) Initiate() (map[string]any, error) {
	body := map[string]any{
		"bvn":                 cv.bvn,
		"first_name":          cv.firstName,
		"customer_identifier": cv.customerID,
		"last_name":           cv.lastName,
		"mobile_num":          cv.mobileNo,
		"email":               cv.email,
		"dob":                 cv.dob,
		"address":             cv.address,
		"gender":              cv.gender,
		"beneficiary_account": cv.beneficiaryAcc,
	}
	return utils.MakeRequest(body, utils.CompleteUrl("virtual-account", cv.live), cv.apiKey)
}
