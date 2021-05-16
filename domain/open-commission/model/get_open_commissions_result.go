package model

type GetOpenCommissionsResult struct {
	OpenCommissions []OpenCommission
	Offset          int
	Total           int
}
