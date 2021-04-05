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
	if d.IsR18 != nil {
		setter = append(setter, bson.E{Key: "isR18", Value: d.IsR18})
	}
	if d.AllowBePrivate != nil {
		setter = append(setter, bson.E{Key: "allowBePrivate", Value: d.AllowBePrivate})
	}
	if d.AllowAnonymous != nil {
		setter = append(setter, bson.E{Key: "allowAnonymous", Value: d.AllowAnonymous})
	}
	if d.State != nil {
		setter = append(setter, bson.E{Key: "state", Value: d.State})
	}
	if d.LastUpdatedTime != nil {
		setter = append(setter, bson.E{Key: "lastUpdatedTime", Value: d.LastUpdatedTime})
	}

	return bson.D{{"$set", setter}}
}
