package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

// Fetch takes account id and resturns nilable accont or error
func (c *Client) Fetch(id string) (*Account, error) {
	rp := fmt.Sprintf("%s/organisation/accounts/%s", c.baseURL, id)
	fmt.Println("rp is = ", rp)

	//create http Get request
	req, err := http.NewRequest(http.MethodGet, rp, nil)
	if err != nil {
		return nil, err
	}
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")

	//make request and get parsed response
	res := Account{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	// res.Code = http.StatusOK
	return &res, nil
}

//Delete takes id and version of account and delete accounts
func (c *Client) Delete(id string, version int) error {

	rp := fmt.Sprintf("%s/organisation/accounts/%s?version=%d", c.baseURL, id, version)
	fmt.Println("rp is = ", rp)

	req, err := http.NewRequest(http.MethodDelete, rp, nil)
	if err != nil {
		return err
	}

	res := Account{}
	if err := c.sendRequest(req, &res); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}

// Create takes accountMessage and return account or error
func (c *Client) Create(acc accountMessage) (*Account, error) {

	rp := fmt.Sprintf("%s/organisation/accounts", c.baseURL)
	fmt.Println("rp is = ", rp)

	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(acc)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, rp, buf)
	if err != nil {
		return nil, err
	}

	res := Account{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
