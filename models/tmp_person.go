package models

import (
	"time"
)

// TmpPerson 人物
type TmpPerson struct {
	ID      int `beedb:"PK"`
	UserID  int `sql:"user_id"`
	Name    string
	Created time.Time
	Updated time.Time
}

// Add 追加する
func (t *TmpPerson) Add() error {
	t.Updated = time.Now()
	t.Created = time.Now()

	return Save(t)
}

// GetIDAndAdd 追加してIDを取得する
func (t *TmpPerson) GetIDAndAdd() (int, error) {
	t.Updated = time.Now()
	t.Created = time.Now()

	err := Save(&t)
	if err != nil {
		return 0, err
	}

	return t.ID, nil
}

// GetFindAll 全て取得する
func (t *TmpPerson) GetFindAll() (tmpPerson []TmpPerson) {
	GetFindAll(&tmpPerson)

	return
}
