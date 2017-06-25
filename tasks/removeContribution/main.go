package main

import (
	"os"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/tasks"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/log"
	"github.com/wheatandcat/dotstamp_server/utils/movie"
	"github.com/wheatandcat/dotstamp_server/utils/sound"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

// RemoveDays 削除日数
const RemoveDays = 3

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "contributionTotalFollows")
	}
}

func main() {
	logs.Batch("start", "removeContribution")

	if err := RemoveSoundDetail(); err != nil {
		tasks.Err(err, "removeContribution")
	}

	if err := RemoveJoinFile(); err != nil {
		tasks.Err(err, "removeContribution")
	}

	logs.Batch("finish", "removeContribution")
}

// RemoveSoundDetail 音声詳細を削除する
func RemoveSoundDetail() error {
	list, err := contributions.GetSoundDetailListByMakeStatusMade()
	if err != nil {
		return err
	}

	list = contributions.GetSoudDetailListBySpecifiedDays(list, RemoveDays)

	for _, v := range list {
		file := strconv.Itoa(v.UserContributionID) + "_" + strconv.Itoa(v.Priority)
		sound.RemoveDetailFile(file)

		v.MakeStatus = models.MakeStatusUncreated
		v.Save()
	}

	return nil
}

// RemoveJoinFile 連結ファイルを削除する
func RemoveJoinFile() error {
	list, err := contributions.GetMovieListByMovieStatusPublic()
	if err != nil {
		return err
	}

	list = contributions.GetMovieListBySpecifiedDays(list, RemoveDays)

	for _, v := range list {
		if !contributions.ExistsMovie(v.UserContributionID) {
			continue
		}

		sound.RemoveJoinFile(strconv.Itoa(v.UserContributionID))
		movie.RemoveFile(strconv.Itoa(v.UserContributionID))
	}

	return nil
}
