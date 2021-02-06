package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pixstall-artist/domain/open-commission/model"
)

func NewUpdaterFromOpenCommissionUpdater(d model.OpenCommissionUpdater) bson.D {
	setter := bson.D{}

	if d.Title != nil {
		setter = append(setter, bson.E{Key: "title", Value: d.Title})
	}
	if d.Desc != nil {
		setter = append(setter, bson.E{Key: "desc", Value: d.Desc})
	}
	if d.DepositRule != nil {
		setter = append(setter, bson.E{Key: "depositRule", Value: d.DepositRule})
	}
	if d.Price != nil {
		setter = append(setter, bson.E{Key: "price", Value: d.Price})
	}
	if d.DayNeed != nil {
		setter = append(setter, bson.E{Key: "dayNeed", Value: d.DayNeed})
	}
	if d.TimesAllowedDraftToChange != nil {
		setter = append(setter, bson.E{Key: "timesAllowedDraftToChange", Value: d.TimesAllowedDraftToChange})
	}
	if d.TimesAllowedCompletionToChange != nil {
		setter = append(setter, bson.E{Key: "TimesAllowedCompletionToChange", Value: d.TimesAllowedCompletionToChange})
	}
	if d.SampleImagePaths != nil {
		setter = append(setter, bson.E{Key: "sampleImagePaths", Value: d.SampleImagePaths})
	}
	if d.State != nil {
		setter = append(setter, bson.E{Key: "state", Value: d.State})
	}

	return bson.D{{"$set", setter}}
}
