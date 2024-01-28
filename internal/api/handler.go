package api

import (
	"net/http"

	"github.com/Geovanny0401/bookmarks/internal/config"
	"github.com/Geovanny0401/bookmarks/internal/domain"
	"github.com/gin-gonic/gin"
)

type BookmarkController struct {
	repo   domain.BookmarkRepository
	logger *config.Logger
}

func NewBookmarkController(repo domain.BookmarkRepository, logger *config.Logger) BookmarkController {
	return BookmarkController{
		repo:   repo,
		logger: logger,
	}
}

func (p BookmarkController) GetAll(c *gin.Context) {
	p.logger.Info("Finding all bookmarks")
	ctx := c.Request.Context()
	bookmarks, err := p.repo.GetAll(ctx)
	if err != nil {
		if err != nil {
			p.logger.Errorf("Error :%v", err)
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch bookmarks",
		})
		return
	}
	c.JSON(http.StatusOK, bookmarks)
}
