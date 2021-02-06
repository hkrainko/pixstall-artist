package model

import "strconv"

type OpenCommissionError int

func (e OpenCommissionError) Error() string {
	switch e {
	case OpenCommissionErrorNotFound:
		return "OpenCommissionErrorNotFound"
	case OpenCommissionErrorDuplicate:
		return "OpenCommissionErrorDuplicate"
	case OpenCommissionErrorUnknown:
		return "OpenCommissionErrorUnknown"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	OpenCommissionErrorNotFound OpenCommissionError = 10
	OpenCommissionErrorDuplicate OpenCommissionError = 11
	OpenCommissionErrorUnknown OpenCommissionError = 99
)