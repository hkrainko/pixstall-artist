package model

type ValidatedCommission struct {
	IsValid bool    `json:"isValid"`
	Reason  *string `json:"reason,omitempty"`
}
