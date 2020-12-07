package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-artist/app/artist/delivery/model/get_artist"
	domainArtist "pixstall-artist/app/domain/artist"
	domain "pixstall-artist/app/domain/artist/model"
)

type ArtistController struct {
	artistUseCase domainArtist.UseCase
}

func NewArtistController(useCase domainArtist.UseCase) ArtistController {
	return ArtistController{
		artistUseCase: useCase,
	}
}

func (a ArtistController) GetArtist(c *gin.Context) {
	artistID := c.Query("artistId")
	artist, err := a.artistUseCase.GetArtist(c, artistID)
	if err != nil {
		return
	}

	c.PureJSON(200, get_artist.NewResponse(*artist))
}

func (a ArtistController) UpdateArtist(c *gin.Context) {
	artistID := c.Query("artistId")
	UserName := c.Query("userName")
	email := c.Query("email")
	birthday := c.Query("birthday")
	gender := c.Query("gender")

	updater := domain.ArtistUpdater{
		ArtistID:         artistID,
		UserName:         &UserName,
		Email:            &email,
		Birthday:         &birthday,
		Gender:           &gender,
		PhotoURL:         nil,
		State:            nil,
		FansIDs:          nil,
		LikeIDs:          nil,
		RegistrationTime: nil,
		ArtistIntro:      nil,
		ArtistDetails:    nil,
		OpenCommissions:  nil,
		Artworks:         nil,
	}

	err := a.artistUseCase.UpdateArtist(c, artistID, updater)
	if err != nil {
		return
	}

	c.PureJSON(200, nil)
}