package order

import (
	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"gopkg.in/resty.v1"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	"strconv"
	"fmt"
)

type Status string


type OrderResp struct {

}

func Get(c client.Client, id int64) error {
	sig := signature.Sign(c.APISecret(), nil)

	req := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig)
		//SetResult(&orders)

	resp, err := req.Get(c.URL() + resourceURL + "/" + strconv.FormatInt(id, 10))
	fmt.Printf("%+s", resp.String())
	return err

}
