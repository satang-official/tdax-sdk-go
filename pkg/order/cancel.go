package order

import (
	"fmt"

	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	resty "gopkg.in/resty.v1"
	"github.com/satang-official/tdax-sdk-go/pkg/utils"
)

func Cancel(c client.Client, orderID int64) error {
	sig := signature.Sign(c.APISecret(), nil)

	resp, err := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		Delete(fmt.Sprintf("%s/%s/%d", c.URL(), resourceURL, orderID))
	if err != nil {
		return err
	}
	err = utils.HandleResponse(resp)
	return err
}
