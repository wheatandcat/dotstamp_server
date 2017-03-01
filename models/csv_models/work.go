package csvModels

import (
	"time"
)

// Work 作品
type Work struct {
	ID         int
	CategoryID int
	Name       string
	AuthorID   int
	Country    string
	Released   time.Time
}

func (w Work) getFileName() string {
	return "work.csv"
}
