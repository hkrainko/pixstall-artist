package model

type GetOpenCommissionResult struct {
	OpenCommissions []OpenCommission
	Offset          int
	Total           int
}
