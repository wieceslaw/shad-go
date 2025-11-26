//go:build !solution

package retryupdate

import (
	"errors"

	"github.com/gofrs/uuid"
	"gitlab.com/slon/shad-go/retryupdate/kvapi"
)

func UpdateValue(c kvapi.Client, key string, updateFn func(oldValue *string) (newValue string, err error)) error {
	var val *string
	var oldVersion uuid.UUID

	for {
	GetLoop:
		for {
			resp, err := c.Get(&kvapi.GetRequest{
				Key: key,
			})

			if err != nil {
				var authErr *kvapi.AuthError
				switch {
				case errors.Is(err, kvapi.ErrKeyNotFound):
					break GetLoop
				case errors.As(err, &authErr):
					return err
				default:
					continue GetLoop
				}
			}

			val = &resp.Value
			oldVersion = resp.Version
			break GetLoop
		}

	UpdateLoop:
		for {
			newValue, err := updateFn(val)
			if err != nil {
				return err
			}
			newVersion := uuid.Must(uuid.NewV4())

		RetryLoop:
			for {
				_, err = c.Set(&kvapi.SetRequest{
					Key:        key,
					Value:      newValue,
					OldVersion: oldVersion,
					NewVersion: newVersion,
				})

				if err != nil {
					var authErr *kvapi.AuthError
					var conflictErr *kvapi.ConflictError
					switch {
					case errors.As(err, &authErr):
						return err
					case errors.Is(err, kvapi.ErrKeyNotFound):
						val = nil
						oldVersion = uuid.UUID{}
						continue UpdateLoop
					case errors.As(err, &conflictErr):
						if newVersion == conflictErr.ExpectedVersion {
							return nil
						}
						break UpdateLoop
					default:
						continue RetryLoop
					}
				}
				return nil
			}
		}
	}
}
