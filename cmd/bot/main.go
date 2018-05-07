package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"bitbucket.org/satangcorp/tdax-sdk/pkg/client"
	"bitbucket.org/satangcorp/tdax-sdk/pkg/user"
	"bitbucket.org/satangcorp/tdax-sdk/pkg/withdrawal"
)

func main() {

	log.Info("Start")
	log.Infof("Key: %s", key)
	log.Infof("Secret: %s", secret)

	c := client.NewClient("https://api.tdax.com/", userID, key, secret)

	resp, err := user.Get(c)
	fmt.Printf("\nMain Req: %v", resp.Request.QueryParam)
	fmt.Printf("\nMain Status Code: %v", resp.StatusCode())
	fmt.Printf("\nMain Body: %v", resp) // or resp.String() or string(resp.Body())

	fmt.Println("\n-----------------------")

	// resp, err = order.List(c, "BTC", "THB")

	// resp, err := order.Cancel(c, 768335)

	// resp, err := order.Create(c, order.NewLimitSell(
	// 	big.NewInt(32100000),
	// 	big.NewInt(1e6),
	// 	"BTC", "THB"))

	// resp, err := order.Create(c, order.NewLimitBuy(
	// 	big.NewInt(27566600),
	// 	big.NewInt(1e5),
	// 	"BTC", "THB"))

	// resp, err = withdrawal.List(c)

	resp, err = withdrawal.Create(c, withdrawal.NewWithdrawal(
		"StellarAddress", "GDHISAMF43GJN3Q2F2WEOFASN4HHM5S435X4GQL62FBMZCY7IHWPEQ5X", "20000000", "100",
	))

	if err != nil {
		panic(err)
	}
	fmt.Printf("\nMain Req: %v", resp.Request.QueryParam)
	fmt.Printf("\nMain Status Code: %v", resp.StatusCode())
	fmt.Printf("\nMain Body: %v", resp)

}

// resp, err := resty.R().Get("http://httpbin.org/get")

// // explore response object
// fmt.Printf("\nError: %v", err)
// fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
// // fmt.Printf("\nResponse Status: %v", resp.Status())
// // fmt.Printf("\nResponse Time: %v", resp.Time())
// // fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
// fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
