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
	case OpenCommissionErrorUnAuth:
		return "OpenCommissionErrorUnAuth"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	OpenCommissionErrorNotFound OpenCommissionError = 10
	OpenCommissionErrorDuplicate OpenCommissionError = 11
	OpenCommissionErrorUnAuth OpenCommissionError = 12
	OpenCommissionErrorUnknown OpenCommissionError = 99
)