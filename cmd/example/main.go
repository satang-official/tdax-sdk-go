package main

import (
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/order"
	"github.com/satang-official/tdax-sdk-go/pkg/user"
	"github.com/satang-official/tdax-sdk-go/pkg/withdrawal"
)

func main() {
	key, secret, userID := "", "", ""

	log.Infof("Key: %s", key)
	log.Infof("Secret: %s", secret)

	c := client.NewClient("https://api.tdax.com/", userID, key, secret)

	// Get User wallet
	resp, err := user.Get(c)
	fmt.Printf("User Wallet: %v\n", resp)

	// Create New Order
	o, err := order.Create(c, order.NewLimitSell(
		big.NewInt(65432100),
		big.NewInt(1e6),
		"BTC", "THB"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Created: %v", o)

	fmt.Println("-----------------------\n")

	// List Withdrawals
	resp, err = withdrawal.List(c)
	fmt.Printf("Withdrawas: %v\n", resp)
	// Create New Withdrawal
	resp, err = withdrawal.Create(c, withdrawal.NewWithdrawal(
		"StellarAddress", "GDHISAMF43GJN3Q2F2WEOFASN4HHM5S435X4GQL62FBMZCY7IHWPEQ5X", "20000000", "100",
	))

	// List Open Orders
	list, err := order.List(c, "BTC", "THB")
	if err != nil {
		panic(err)
	}

	// Delete Open Orders
	for _, o := range list.Items {
		if o.Side == order.BuySide {
			fmt.Printf("Delete Buy Order %+v\n", o)
			order.Cancel(c, o.ID)
		}
	}

}
