package dao

import "pixstall-artist/domain/open-commission/model"

type GetOpenCommissionsResult struct {
	OpenCommissions []OpenCommission `bson:"openCommissions"`
	Total           int              `bson:"total"`
}

func (g GetOpenCommissionsResult) ToDomainGetOpenCommissionsResult(offset int) *model.GetOpenCommissionsResult {

	var dOpenComms []model.OpenCommission
	for _, oc := range g.OpenCommissions {
		dOpenComms = append(dOpenComms, *oc.ToDomainOpenCommission())
	}

	return &model.GetOpenCommissionsResult{
		OpenCommissions: dOpenComms,
		Offset:          offset,
		Total:           g.Total,
	}
}