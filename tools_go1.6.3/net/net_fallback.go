// +build !darwin,!linux,!freebsd,!windows

package net

import "../common"

func IOCounters(pernic bool) ([]IOCountersStat, error) {
	return []IOCountersStat{}, common.ErrNotImplementedError
}

func FilterCounters() ([]FilterStat, error) {
	return []FilterStat{}, common.ErrNotImplementedError
}

func ProtoCounters(protocols []string) ([]ProtoCountersStat, error) {
	return []ProtoCountersStat{}, common.ErrNotImplementedError
}

func Connections(kind string) ([]ConnectionStat, error) {
	return []ConnectionStat{}, common.ErrNotImplementedError
}