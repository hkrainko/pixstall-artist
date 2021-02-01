package get_artist

import (
	"net/http"
	"pixstall-artist/domain/artist/model"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	if artistError, isError := err.(model.ArtistError); isError {
		switch artistError {
		case model.ArtistErrorNotFound:
			return http.StatusNotFound, ErrorResponse{
				Message: artistError.Error(),
			}
		case model.ArtistErrorUnAuth:
			return http.StatusUnauthorized, ErrorResponse{
				Message: artistError.Error(),
			}
		default:
			return http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			}
		}
	} else {
		return http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		}
	}

}
