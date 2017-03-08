package main

import (
	"os"

	"dotstamp_server/models"
	"dotstamp_server/tasks"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/follow"
	"dotstamp_server/utils/log"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "contributionTotalFollows")
	}
}

func main() {
	logs.Batch("start", "contributionTotalFollows")

	tx := models.Begin()

	if err = AddContributionTotalFollows(); err != nil {
		models.Rollback(tx)
		tasks.Err(err, "contributionTotalFollows")
		return
	}

	if err = SaveUserContributionSearchToFollowCount(); err != nil {
		models.Rollback(tx)
		tasks.Err(err, "contributionTotalFollows")
		return
	}

	models.Commit(tx)

	logs.Batch("finish", "contributionTotalFollows")
}

// AddContributionTotalFollows フォロー数を追加する
func AddContributionTotalFollows() error {
	contributionIDList, err = contributions.GetViewStatusPublicIDList()
	if err != nil {
		return err
	}

	followList, err := follows.GetListByUserContributionIDList(contributionIDList)
	if err != nil {
		return err
	}

	followMap = follows.GetFollowCountMap(followList)

	follows.TruncateTotal()

	if err = follows.AddTotalMap(followMap); err != nil {
		return err
	}

	return nil
}

// SaveUserContributionSearchToFollowCount 検索のフォロー数を更新する
func SaveUserContributionSearchToFollowCount() error {
	search, err := contributions.GetSearchListByUserContributionIDList(contributionIDList)
	if err != nil {
		return err
	}

	return contributions.SaveToFollowCount(search, followMap)
}
