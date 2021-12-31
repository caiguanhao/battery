//go:build !linux && !windows && !darwin

package battery

import "errors"

func getStatus() (*Status, error) {
	return nil, ErrNotSupported
}
