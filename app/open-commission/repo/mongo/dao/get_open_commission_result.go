package dao

import "pixstall-artist/domain/open-commission/model"

type GetOpenCommissionResult struct {
	OpenCommissions []OpenCommission `bson:"openCommissions"`
	Total           int              `bson:"total"`
}

func (g GetOpenCommissionResult) ToDomainGetOpenCommissionResult(offSet int) *model.GetOpenCommissionResult {

	var dOpenComms []model.OpenCommission
	for _, oc := range g.OpenCommissions {
		dOpenComms = append(dOpenComms, *oc.ToDomainOpenCommission())
	}

	return &model.GetOpenCommissionResult{
		OpenCommissions: dOpenComms,
		Offset:          offSet,
		Total:           g.Total,
	}
}