package http

import (
	"pixstall-artist/domain/artist/model"
	model2 "pixstall-artist/domain/model"
)

func getArtistsSorter(str string) *model.ArtistSorter {
	if str == "" {
		return nil
	}
	sorter := model.ArtistSorter{}
	symbol := str[:1]
	if symbol == "-" {
		switch str[1:len(str)] {
		case "reg-time":
			v := model2.SortOrderDescending
			sorter.RegTime = &v
		default:
			break
		}
	} else {
		switch str {
		case "reg-time":
			v := model2.SortOrderAscending
			sorter.RegTime = &v
		default:
			break
		}
	}
	return &sorter
}