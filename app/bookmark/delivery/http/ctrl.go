package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	add_bookmark "pixstall-artist/app/bookmark/delivery/http/resp/add-bookmark"
	get_bookmarks "pixstall-artist/app/bookmark/delivery/http/resp/get-bookmarks"
	http2 "pixstall-artist/app/error/http"
	"pixstall-artist/domain/bookmark"
	error2 "pixstall-artist/domain/error"
	"strconv"
)

type BookmarkController struct {
	bookmarkUseCase bookmark.UseCase
}

func NewBookmarkController(bookmarkUseCase bookmark.UseCase) BookmarkController {
	return BookmarkController{
		bookmarkUseCase: bookmarkUseCase,
	}
}

func (b BookmarkController) AddBookmark(c *gin.Context) {
	tokenUserID := c.GetString("userId")
	artistID := c.Param("id")

	err := b.bookmarkUseCase.AddBookmark(c, tokenUserID, artistID)
	if err != nil {
		c.AbortWithStatusJSON(http2.NewErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, add_bookmark.Response{ArtistID: artistID})
}

func (b BookmarkController) GetBookmarks(c *gin.Context) {
	tokenUserID := c.GetString("userId")
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		c.JSON(http2.NewErrorResponse(error2.BadRequestError))
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http2.NewErrorResponse(error2.BadRequestError))
		return
	}
	result, err := b.bookmarkUseCase.GetBookmarksForUser(c, tokenUserID, count, offset)
	if err != nil {
		c.AbortWithStatusJSON(http2.NewErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, get_bookmarks.NewResponse(*result, count, offset))
}

func (b BookmarkController) DeleteBookmark(c *gin.Context) {
	tokenUserID := c.GetString("userId")
	artistID := c.Param("id")


}