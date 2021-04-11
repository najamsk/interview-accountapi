package main

import (
	"fmt"

	"github.com/najamsk/interview-accountapi/pkg/account"
)

func main() {
	fmt.Println("start account test client")
	c := account.NewClient()

	//ad27e265-9605-4b4b-a0e5-3003ea9cc4dc
	//ad27e265-9605-4b4b-a0e5-3003ea9cc4dc
	o, err := c.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	if err != nil {
		fmt.Println("client eror = ", err)
	}
	fmt.Printf("client.fetch = %#v \n", o)
}
