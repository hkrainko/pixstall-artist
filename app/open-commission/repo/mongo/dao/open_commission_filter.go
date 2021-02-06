package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pixstall-artist/domain/open-commission/model"
)

func NewFilterFromDomainOpenCommissionFilter(d model.OpenCommissionFilter) bson.D {

	filter := bson.D{}

	if d.ArtistID != nil {
		filter = append(filter, bson.E{Key: "artistId", Value: d.ArtistID})
	}
	if d.PriceForm != nil {
		filter = append(filter, bson.E{Key: "price", Value: bson.M{"$gte": d.PriceForm}})
	}
	if d.PriceTo != nil {
		filter = append(filter, bson.E{Key: "price", Value: bson.M{"lte": d.PriceTo}})
	}
	if d.DayNeedGreaterOrEqual != nil {
		filter = append(filter, bson.E{Key: "dayNeed", Value: bson.M{"$gte": d.DayNeedGreaterOrEqual}})
	}
	if d.DayNeedLessOrEqual != nil {
		filter = append(filter, bson.E{Key: "dayNeed", Value: bson.M{"lte": d.DayNeedLessOrEqual}})
	}

	return filter
}