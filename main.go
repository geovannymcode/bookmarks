package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Geovanny0401/bookmarks/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	logger := config.NewLogger(cfg)
	db := config.GetDb(cfg)

	r := gin.Default()
	r.GET("/api/bookmarks", getAllBookmarks(db, logger))
	log.Fatal(r.Run(fmt.Sprintf(":%d", cfg.ServerPort)))
}

func getAllBookmarks(db *pgx.Conn, logger *config.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		bookmarks, err := getAll(ctx, db)
		if err != nil {
			logger.Errorf("Error fetching bookmarks from db: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failing to fetch bookmarks",
			})
		}
		c.JSON(http.StatusOK, bookmarks)
	}
}

type Bookmark struct {
	ID        int
	Title     string
	Url       string
	CreatedAt time.Time
}

func getAll(ctx context.Context, db *pgx.Conn) ([]Bookmark, error) {
	query := `select id, title, url, created_at FROM bookmarks`
	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []Bookmark
	for rows.Next() {
		var bookmark = Bookmark{}
		err = rows.Scan(&bookmark.ID, &bookmark.Title, &bookmark.Url, &bookmark.CreatedAt)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bookmarks, nil
}
