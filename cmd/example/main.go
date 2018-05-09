package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/order"
	"github.com/satang-official/tdax-sdk-go/pkg/user"
)

func main() {
	// key, secret, userID := "", "", ""

	log.Infof("Key: %s", key)
	log.Infof("Secret: %s", secret)

	c := client.NewClient("https://api.tdax.com/", userID, key, secret)

	resp, err := user.Get(c)
	fmt.Printf("User: %v\n", resp)

	fmt.Println("-----------------------\n")

	// o, err := order.Create(c, order.NewLimitSell(
	// 	big.NewInt(65432100),
	// 	big.NewInt(1e6),
	// 	"BTC", "THB"))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Created Order: %v", o)

	fmt.Println("-----------------------\n")

	// resp, err = withdrawal.List(c)
	// resp, err = withdrawal.Create(c, withdrawal.NewWithdrawal(
	// 	"StellarAddress", "GDHISAMF43GJN3Q2F2WEOFASN4HHM5S435X4GQL62FBMZCY7IHWPEQ5X", "20000000", "100",
	// ))

	fmt.Println("-----------------------")
	oss, err := order.List(c, "BTC", "THB")
	if err != nil {
		panic(err)
	}

	for _, o := range oss.Items {
		fmt.Printf("O %+v\n", o)
		if o.Side == order.BuySide {
			fmt.Printf("Buy %+v\n", o)
		}
	}

	// order.Cancel(c, o.ID)

}
