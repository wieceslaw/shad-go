//go:build !solution

package retryupdate

import (
	"errors"

	"gitlab.com/slon/shad-go/retryupdate/kvapi"
)

func UpdateValue(c kvapi.Client, key string, updateFn func(oldValue *string) (newValue string, err error)) error {
	var val *string

	resp, err := c.Get(&kvapi.GetRequest{
		Key: key,
	})
	var apiErr kvapi.APIError
	if errors.As(err, &apiErr) {
		var authErr kvapi.AuthError
		switch err := apiErr.Unwrap(); {
		case err == kvapi.ErrKeyNotFound:
			{
				val = nil
			}
		case errors.As(err, &authErr):
			{
				return &authErr
			}
		}
	}

	newValue, err := updateFn(val)
	if err != nil {
		return nil
	}

	return err
}
