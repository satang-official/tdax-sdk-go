package utils

import (
	"gopkg.in/resty.v1"
	"errors"
	"fmt"
)

func HandleResponse(resp *resty.Response) error {
	statusCode := resp.RawResponse.StatusCode
	statusText := resp.RawResponse.Status
	if statusCode >= 400 {
		errorText := fmt.Sprintf("%d : %s", statusCode, statusText)
		return errors.New(errorText)
	} else {
		return nil
	}
}