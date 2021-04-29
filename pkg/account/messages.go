/* 
Package account implements a rest api client
*/
package account

type Account struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Version        int        `json:"version"`
	Attributes     attributes `json:"attributes"`
	// Relationships  Relationships `json:"relationships"`
}

type attributes struct {
	Country       string `json:"country"`
	BaseCurrency  string `json:"base_currency"`
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
	Bic           string `json:"bic"`
	Iban          string `json:"iban"`
	Status        string `json:"status"`
}
