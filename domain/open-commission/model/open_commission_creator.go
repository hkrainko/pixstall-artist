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
	SampleImages                   []image.Image
	SampleImagePaths               []string
}
