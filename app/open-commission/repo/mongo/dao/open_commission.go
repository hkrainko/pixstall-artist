package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-artist/domain/open-commission/model"
	"time"
)

type OpenCommission struct {
	ObjectID             primitive.ObjectID `bson:"_id,omitempty"`
	model.OpenCommission `bson:",inline"`
}

func NewFromDomainOpenCommissionCreator(artistID string, d model.OpenCommissionCreator, openCommID string) OpenCommission {
	return OpenCommission{
		OpenCommission: model.OpenCommission{
			ID:                     openCommID,
			ArtistID:                       artistID,
			Title:                          d.Title,
			Desc:                           d.Desc,
			DepositRule:                    d.DepositRule,
			Price:                          d.Price,
			DayNeed:                        d.DayNeed,
			TimesAllowedDraftToChange:      d.TimesAllowedDraftToChange,
			TimesAllowedCompletionToChange: d.TimesAllowedCompletionToChange,
			SampleImagePaths:               d.SampleImagePaths,
			IsR18:                          d.IsR18,
			AllowBePrivate:                 d.AllowBePrivate,
			AllowAnonymous:                 d.AllowAnonymous,
			State:                          model.OpenCommissionStateActive,
			CreateTime:                     time.Now(),
			LastUpdatedTime:                time.Now(),
		},
	}
}

func (o *OpenCommission) ToDomainOpenCommission() *model.OpenCommission {
	return &model.OpenCommission{
		ID:                             o.ID,
		ArtistID:                       o.ArtistID,
		Title:                          o.Title,
		Desc:                           o.Desc,
		DepositRule:                    o.DepositRule,
		Price:                          o.Price,
		DayNeed:                        o.DayNeed,
		TimesAllowedDraftToChange:      o.TimesAllowedDraftToChange,
		TimesAllowedCompletionToChange: o.TimesAllowedCompletionToChange,
		SampleImagePaths:               o.SampleImagePaths,
		IsR18:                          o.IsR18,
		AllowBePrivate:                 o.AllowBePrivate,
		AllowAnonymous:                 o.AllowAnonymous,
		State:                          o.State,
		CreateTime:                     o.CreateTime,
		LastUpdatedTime:                o.LastUpdatedTime,
	}
}
