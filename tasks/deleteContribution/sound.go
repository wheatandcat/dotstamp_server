package main

import (
	"os"

	"dotstamp_server/tasks"
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

func deleteSound() {

}
