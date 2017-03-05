package test

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Test テスティング
func Test(t *testing.T) {
	TestingT(t)
}

// TSuite テストスイーツ
type TSuite struct {
	TableNameList []string
}

// Accessor テストスイーツ
type Accessor interface {
	SetUpSuite(*C)
	SetUpTest(*C)
	TearDownTest(*C)
	TearDownSuite(*C)
	SetTableNameList([]string)
}

// SetUpSuite テストスイーツセットアップ
func (t *TSuite) SetUpSuite(c *C) {
	Setup()

	removeLogFile("error")
	removeLogFile("batch")
}

// SetUpTest テストセットアップ
func (t *TSuite) SetUpTest(c *C) {
	SetupFixture(t.TableNameList)
}

// TearDownTest テストダウン
func (t *TSuite) TearDownTest(c *C) {}

// TearDownSuite テストスイーツダウン
func (t *TSuite) TearDownSuite(c *C) {}

// SetTableNameList テーブル名を設定する
func (t *TSuite) SetTableNameList(s []string) {
	t.TableNameList = s
}
