package model

import "image"

type OpenCommissionCreator struct {
	Title                          string
	Desc                           string
	DepositRule                    *string
	Price                          Price
	DayNeed                        DayNeed
	TimesAllowedDraftToChange      *int
	TimesAllowedCompletionToChange *int
	IsR18                          bool
	AllowBePrivate                 bool
	AllowAnonymous                 bool
	SampleImages                   []image.Image
	SampleImagePaths               []string
}
