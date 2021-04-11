package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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

// type Relationships struct {
// 	AccountEvents AccountEvents `json:"account_events"`
// }

// type AccountEvents struct {
// 	Data []AccountEventData `json:"data"`
// }

// type AccountEventData struct {
// 	Type string `json:"type"`
// 	ID   string `json:"id"`
// }

func (c *Client) Fetch(id string) (*AccountRes, error) {
	rp := fmt.Sprintf("%s/organisation/accounts/%s", c.baseURL, id)
	fmt.Println("rp is = ", rp)

	//create http Get request
	req, err := http.NewRequest(http.MethodGet, rp, nil)
	if err != nil {
		return nil, err
	}

	//make http request
	o, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer o.Body.Close()

	//check request response
	if o.StatusCode != http.StatusOK {
		var errRes errResponse
		if err = json.NewDecoder(o.Body).Decode(&errRes); err == nil {
			//if http api error_message is decoded without any err return that
			return nil, errors.New(errRes.Message)
		}

		return nil, fmt.Errorf("unknown error, status code %d", o.StatusCode)
	}

	//decode api response into AccountRes

	res := AccountRes{}
	if err := json.NewDecoder(o.Body).Decode(&res); err != nil {
		return nil, err
	}
	res.Code = http.StatusOK
	return &res, nil
}

func (c *Client) Delete(id string, version int) (*AccountRes, error) {

	rp := fmt.Sprintf("%s/organisation/accounts/%s?version=%d", c.baseURL, id, version)
	fmt.Println("rp is = ", rp)

	//create http Get request
	req, err := http.NewRequest(http.MethodDelete, rp, nil)
	if err != nil {
		return nil, err
	}

	//make http request
	o, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer o.Body.Close()

	//check request response
	if o.StatusCode != http.StatusNoContent {
		var errRes errResponse
		if err = json.NewDecoder(o.Body).Decode(&errRes); err == nil {
			//if http api error_message is decoded without any err return that
			return nil, errors.New(errRes.Message)
		}

		return nil, fmt.Errorf("unknown error, status code %d", o.StatusCode)
	}

	res := AccountRes{}
	// if err := json.NewDecoder(o.Body).Decode(&res); err != nil {
	// 	return nil, err
	// }
	res.Code = http.StatusNoContent
	return &res, nil
}
