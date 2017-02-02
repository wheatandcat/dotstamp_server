package models

import (
	"time"
)

// TmpWork 作品
type TmpWork struct {
	ID         int `beedb:"PK"`
	UserID     int `sql:"user_id"`
	CategoryID int `sql:"category_id"`
	Name       string
	AuthorID   int `sql:"author_id"`
	CountryID  int `sql:"country_id"`
	Released   time.Time
	DeleteFlag int `sql:"delete_flag"`
	Created    time.Time
	Updated    time.Time
}

// Add 追加する
func (t *TmpWork) Add() error {
	t.Updated = time.Now()
	t.Created = time.Now()

	return Save(t)
}

// GetFindAll 全て取得する
func (t *TmpWork) GetFindAll() (tmpWork []TmpWork) {
	GetFindAll(&tmpWork)

	return
}
