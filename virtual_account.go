/*
 * squad-go - A simple wrapper written over squad apis payment endpoints.
 * this file contains the methods used to interact with the virtual account endpoints.
 * Author - Joseph Folayan
 * Github - joey1123455
 */

package squad

// an interface exposing virtual accounts of either type custmer or bussiness model
type VirtualAccount interface{}

/*
  - bussinessVA - a struct representing a virtual account for a bussiness model.
  - @customerID - a unique string representing the bussiness id.
  - @bussinessName - a string representings the bussiness name. all virtual accounts must carry a slug as a prefix
    to the name. The slug must be a portion of the bussiness name or abbreviations of your business name used to open
    the squad account as one word. Please note that slash (/) is not allowed and only hyphen (-) can be used.
  - @mobileNo - customers mobile no
  - @beneficiaryAcc - customers account no
*/
type bussinessVA struct {
	customerID     string
	bussinessName  string
	mobileNo       string
	beneficiaryAcc string
}

/*
 * NewBussinessVirtualAcc - function that returns a bussines virtual account.
 * @id - a unique string representing the bussiness id.
 * @squadACCName - the name associated with the squad account api key.
 * @name - a string representings the bussiness name. all virtual accounts must carry a slug as a prefix
    to the name. The slug must be a portion of the bussiness name or abbreviations of your business name used to open
    the squad account as one word. Please note that slash (/) is not allowed and only hyphen (-) can be used.
 * @acc - users original account no.
*/
// func NewBussinessVirtualAcc(apiKey string, id, squadACCName, name, no, acc string) (VirtualAccount, error) {

// 	return &bussinessVA{
// 		customerID:     id,
// 		BussinessName:  name,
// 		MobileNo:       no,
// 		BeneficiaryAcc: acc,
// 	}, nil
// }
