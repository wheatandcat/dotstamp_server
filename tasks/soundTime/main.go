package main

import (
	"github.com/wheatandcat/dotstamp_server/tasks"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/sound"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	contribution()
}

// contribution 投稿を確認する
func contribution() error {
	contributionIDList, err := contributions.GetViewStatusPublicIDList()
	if err != nil {
		tasks.Err(err, "moveTime")
	}

	for _, id := range contributionIDList {
		second, err := sound.GetLength(strconv.Itoa(id))
		if err != nil {
			return err
		}

		detail, err := contributions.GetDetailByUserContributionID(id)
		if err != nil {
			return err
		}

		s, err := contributions.GetSearchWordBody(detail.Body)
		if err != nil {
			return err
		}

		err = contributions.AddOrSaveSoundLength(id, int(math.Ceil(second)), utf8.RuneCountInString(s))
		if err != nil {
			return err
		}

	}

	return nil
}
