package account

type AccountRes struct {
	Data Account `json:"data"` //could be interface{} so we can pass any kind of data
	Code int     `json:"code"`
	//TODO: Links ??
}

type Account struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
	// Relationships  Relationships `json:"relationships"`
}

type Attributes struct {
	Country       string `json:"country"`
	BaseCurrency  string `json:"base_currency"`
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
	Bic           string `json:"bic"`
	Iban          string `json:"iban"`
	Status        string `json:"status"`
}
