package model

type CommissionError int

func (e CommissionError) Error() string {
	switch e {
	case CommissionErrorNotFound:
		return "CommissionErrorNotFound"
	case CommissionErrorUnAuth:
		return "CommissionErrorUnAuth"
	default:
		return "CommissionErrorUnknown"
	}
}

const (
	CommissionErrorNotFound CommissionError = 10
	CommissionErrorUnAuth   CommissionError = 11
	CommissionErrorUnknown  CommissionError = 99
)
