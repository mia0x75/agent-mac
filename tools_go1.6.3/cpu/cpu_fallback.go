// +build !darwin,!linux,!freebsd,!windows

package cpu

import (
	"time"

	"../internal/common"
)

func Times(percpu bool) ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func Info() ([]InfoStat, error) {
	return []InfoStat{}, common.ErrNotImplementedError
}

func Percent(interval time.Duration, percpu bool) ([]float64, error) {
	return []float64{}, common.ErrNotImplementedError
}
