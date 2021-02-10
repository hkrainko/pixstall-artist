package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-artist/domain/open-commission/model"
	"time"
)

type OpenCommission struct {
	ObjectID                       primitive.ObjectID       `bson:"_id,omitempty"`
	OpenCommID                     string                   `bson:"openCommId"`
	ArtistID                       string                   `bson:"artistId"`
	Title                          string                   `bson:"title"`
	Desc                           string                   `bson:"desc"`
	DepositRule                    *string                  `bson:"depositRule"`
	Price                          model.Price              `bson:"price,omitempty"`
	DayNeed                        model.DayNeed            `bson:"dayNeed"`
	TimesAllowedDraftToChange      *int                     `bson:"timesAllowedDraftToChange,omitempty"`
	TimesAllowedCompletionToChange *int                     `bson:"timesAllowedCompletionToChange,omitempty"`
	SampleImagePaths               []string                 `bson:"sampleImagePaths"`
	State                          model.OpenCommissionSate `bson:"state"`
	CreateTime                     time.Time                `bson:"createTime"`
	LastUpdatedTime                time.Time                `bson:"lastUpdatedTime"`
}

func NewFromDomainOpenCommissionCreator(artistID string, d model.OpenCommissionCreator, openCommID string) OpenCommission {
	return OpenCommission{
		OpenCommID:                     openCommID,
		ArtistID:                       artistID,
		Title:                          d.Title,
		Desc:                           d.Desc,
		DepositRule:                    d.DepositRule,
		Price:                          d.Price,
		DayNeed:                        d.DayNeed,
		TimesAllowedDraftToChange:      d.TimesAllowedDraftToChange,
		TimesAllowedCompletionToChange: d.TimesAllowedCompletionToChange,
		SampleImagePaths:               d.SampleImagePaths,
		State:                          model.OpenCommissionStateActive,
		CreateTime:                     time.Now(),
		LastUpdatedTime:                time.Now(),
	}
}

func (o *OpenCommission) ToDomainOpenCommission() *model.OpenCommission {
	return &model.OpenCommission{
		ID:                             o.OpenCommID,
		ArtistID:                       o.ArtistID,
		Title:                          o.Title,
		Desc:                           o.Desc,
		DepositRule:                    o.DepositRule,
		Price:                          o.Price,
		DayNeed:                        o.DayNeed,
		TimesAllowedDraftToChange:      o.TimesAllowedDraftToChange,
		TimesAllowedCompletionToChange: o.TimesAllowedCompletionToChange,
		SampleImagePaths:               o.SampleImagePaths,
		State:                          o.State,
		CreateTime:                     o.CreateTime,
		LastUpdatedTime:                o.LastUpdatedTime,
	}
}
