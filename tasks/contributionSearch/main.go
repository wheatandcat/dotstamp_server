package contributionSearch

import (
	"os"

	"dotstamp_server/models"
	"dotstamp_server/tasks"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/log"
	"dotstamp_server/utils/tag"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "contributionSearch")
	}
}

func main() {
	logs.Batch("start", "contributionSearch")

	tx := models.Begin()

	if err = ResetSearch(); err != nil {
		models.Rollback(tx)
		tasks.Err(err, "contributionSearch")
		return
	}

	models.Commit(tx)

	logs.Batch("finish", "contributionSearch")
}

// ResetSearch 検索をリセットする
func ResetSearch() error {
	if err = contributions.TruncateSearch(); err != nil {
		return err
	}

	contributionIDList, err = contributions.GetViewStatusPublicIDList()
	if err != nil {
		return err
	}

	for id := range contributionIDList {
		u, err := contributions.GetByUserContributionID(id)
		if err != nil {
			return err
		}

		t, err := tags.GetTagNameJoin(id)
		if err != nil {
			return err
		}

		detail, err := contributions.GetDetailByUserContributionID(id)
		if err != nil {
			return err
		}

		searchWord := contributions.SearchWord{
			Title: u.Title,
			Body:  detail.Body,
			Tag:   t,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddOrSaveSearch(id, s); err != nil {
			return err
		}
	}

	return nil
}
