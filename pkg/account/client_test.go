package account

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchAccount(t *testing.T) {
	//Arrange
	c := NewClient()

	a, err := createAccountHelper()
	if err != nil {
		t.Errorf("can't create account first time, error =%#v \n", err)
	}

	//Act
	r, err := c.Fetch(a.ID)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, r.ID, a.ID)

	//Clean up
	c.Delete(a.ID, 0)
}

func TestFetchAccountNotFound(t *testing.T) {
	c := NewClient()

	id := "b83dc772-6a9c-4375-b693-9e5ad8cd1e54"
	r, err := c.Fetch(id)

	fmt.Printf("r=%#v \n", r)

	assert.NotNil(t, err)
	assert.Nil(t, r)
}

func TestFetchAccountInvalidID(t *testing.T) {
	c := NewClient()

	id := "b8"
	r, err := c.Fetch(id)

	fmt.Printf("r=%#v \n", r)

	assert.NotNil(t, err)
	assert.Nil(t, r)
}

func TestDeleteAccount(t *testing.T) {
	//Arrange
	c := NewClient()
	//create account first time
	a, err := createAccountHelper()
	if err != nil {
		t.Errorf("can't create account first time, error =%#v \n", err)
	}

	//Act
	ver := 0
	err = c.Delete(a.ID, ver)

	//Assert
	assert.Nil(t, err)
}

func TestDeleteAccountInvalidIdFormat(t *testing.T) {
	c := NewClient()
	id := "adfa"
	ver := 0
	err := c.Delete(id, ver)
	assert.NotNil(t, err)
}

func TestCreateAccount(t *testing.T) {
	//Arrange
	c := NewClient()
	c.Delete("6ba7b810-9dad-11d1-80b4-00c04fd430c8", 0)

	acc := Account{}
	acc.ID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	acc.Type = "accounts"
	acc.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"

	acc.Attributes = attributes{
		Country:      "AU",
		BaseCurrency: "AUD",
		BankID:       "700300",
		BankIDCode:   "AUBSB",
		Bic:          "AUBKGB23",
	}
	accD := accountMessage{
		Data: acc,
	}

	// Act
	r, err := c.Create(accD)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	if err != nil {
		fmt.Printf("err=%#v \n", err.Error())
		t.Fail()
	}

	//Cleanup
	c.Delete("6ba7b810-9dad-11d1-80b4-00c04fd430c8", 0)
}

func TestCreateAccountSameID(t *testing.T) {
	//Arrange
	c := NewClient()
	c.Delete("6ba7b810-9dad-11d1-80b4-00c04fd430c8", 0)

	//first delete so that we know account doesnt exist
	id := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	ver := 0
	err := c.Delete(id, ver)

	if err != nil {
		t.Errorf("can't delete account first time, error =%#v \n", err)
	}

	//create account first time
	acc := Account{}
	acc.ID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	acc.Type = "accounts"
	acc.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	acc.Attributes = attributes{
		Country:      "AU",
		BaseCurrency: "AUD",
		BankID:       "700300",
		BankIDCode:   "AUBSB",
		Bic:          "AUBKGB23",
	}
	accD := accountMessage{
		Data: acc,
	}

	r, err := c.Create(accD)
	if err != nil {
		t.Errorf("can't create account first time, error =%#v \n", err)
	}
	//Act
	//create 2nd time
	r, err = c.Create(accD)

	//Assert
	assert.NotNil(t, err)
	assert.Nil(t, r)
	c.Delete("6ba7b810-9dad-11d1-80b4-00c04fd430c8", 0)
}

func TestDeleteLastHelper(t *testing.T) {
	c := NewClient()
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	ver := 0
	err := c.Delete(id, ver)
	assert.Nil(t, err)
}

// func TestCreateAccountHelper(t *testing.T) {
// 	c := NewClient()
// 	acc := Account{}

// 	acc.ID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
// 	acc.Type = "accounts"
// 	acc.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"

// 	acc.Attributes = attributes{
// 		Country:      "AU",
// 		BaseCurrency: "AUD",
// 		BankID:       "700300",
// 		BankIDCode:   "AUBSB",
// 		Bic:          "AUBKGB23",
// 	}
// 	accD := accountMessage{
// 		Data: acc,
// 	}
// 	_, e := c.Create(accD)
// 	assert.Nil(t, e)
// }

func createAccountHelper() (*Account, error) {
	c := NewClient()
	acc := Account{}

	acc.ID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	acc.Type = "accounts"
	acc.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"

	acc.Attributes = attributes{
		Country:      "AU",
		BaseCurrency: "AUD",
		BankID:       "700300",
		BankIDCode:   "AUBSB",
		Bic:          "AUBKGB23",
	}
	accD := accountMessage{
		Data: acc,
	}
	a, e := c.Create(accD)
	return a, e
}
