package account

import (
	"fmt"
	"net/http"
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
	assert.Equal(t, id, r.Data.ID)
	assert.Equal(t, http.StatusOK, r.Code)
}

func TestDeleteAccount(t *testing.T) {
	c := NewClient()

	//TODO: first create account and then fetch same account and check important fields and id should be not nil
	//TODO: do I need to put http status code as field in our response struct?
	id := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	ver := 0
	r, err := c.Delete(id, ver)

	fmt.Printf("delete response = %#v \n", r)

	assert.Nil(t, err)
	// assert.NotNil(t, r)
	// assert.Equal(t, id, r.Data.ID)
	// assert.Equal(t, http.StatusOK, r.Code)
}
