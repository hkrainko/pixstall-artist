package http

import (
	"net/http"
	dError "pixstall-artist/domain/error"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	if domainError, isError := err.(dError.DomainError); isError {
		switch domainError {
		case dError.UnAuthError:
			return http.StatusUnauthorized, ErrorResponse{
				Message: domainError.Error(),
			}
		case dError.NotFoundError:
			return http.StatusNotFound, ErrorResponse{
				Message: domainError.Error(),
			}
		case dError.BadRequestError:
			return http.StatusBadRequest, ErrorResponse{
				Message: domainError.Error(),
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