package account

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	Fetch
//err: account id should be uuid format
//err: record doesnt exist against account id
//data: response coming back

func TestFetchAccount(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	r, err := c.Fetch(id)

	fmt.Printf("r=%#v \n", r)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	// assert.Equal(t, id, r.Data.ID)
	// assert.Equal(t, http.StatusOK, r.Code)
}

func TestFetchAccountNotFound(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "b83dc772-6a9c-4375-b693-9e5ad8cd1e54"
	r, err := c.Fetch(id)

	fmt.Printf("r=%#v \n", r)

	assert.NotNil(t, err)
	assert.Nil(t, r)
	// assert.Equal(t, id, r.Data.ID)
	// assert.Equal(t, http.StatusOK, r.Code)
}

func TestFetchAccountInvalidID(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "b8"
	r, err := c.Fetch(id)

	fmt.Printf("r=%#v \n", r)

	assert.NotNil(t, err)
	assert.Nil(t, r)
	// assert.Equal(t, id, r.Data.ID)
	// assert.Equal(t, http.StatusOK, r.Code)
}

func TestDeleteAccount(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	ver := 0
	err := c.Delete(id, ver)

	assert.Nil(t, err)
}

func TestDeleteAccountInvalidVersion(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	ver := 122
	err := c.Delete(id, ver)

	assert.NotNil(t, err)
}
func TestDeleteAccountInvalidIdFormat(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73"
	ver := 0
	err := c.Delete(id, ver)

	assert.NotNil(t, err)
}
func TestDeleteAccountIDNotFound(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "b83dc772-6a9c-4375-b693-9e5ad8cd1e54"
	ver := 0
	err := c.Delete(id, ver)

	assert.NotNil(t, err)
}

// eb0bd6f5-c3f5-44b2-b677-acd23cdde73c

func TestCreateAccount(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?

	acc := Account{}

	acc.ID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	acc.Type = "accounts"
	acc.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"

	acc.Attributes = Attributes{
		Country:      "AU",
		BaseCurrency: "AUD",
		BankID:       "700300",
		BankIDCode:   "AUBSB",
		Bic:          "AUBKGB23",
	}
	accD := AccountRes{
		Data: acc,
	}
	r, err := c.Create(accD)

	// fmt.Printf("create response = %#v \n", r)
	// fmt.Printf("err=%#v \n", err.Error())

	assert.Nil(t, err)
	assert.NotNil(t, r)

	if err != nil {
		fmt.Printf("err=%#v \n", err.Error())
		t.Fail()
	}
	fmt.Printf("create response = %#v \n", r)
	// assert.Equal(t, acc.ID, r.Data.ID)
	// assert.Equal(t, http.StatusCreated, r.Code)
}
