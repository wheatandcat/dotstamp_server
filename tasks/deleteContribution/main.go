package main

import (
	"os"

	"dotstamp_server/tasks"
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

	if err = SaveUserContributionSearchToFollowCount(); err != nil {
		tasks.Err(err, "contributionTotalFollows")
		return
	}

	logs.Batch("finish", "contributionTotalFollows")
}
