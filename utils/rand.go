package utils

import (
	"math/rand"
	"time"

	"github.com/astaxie/beego"
)

var testRandNum int
var testRandString string

// SetTestRandNum テストランダム数を設定する
func SetTestRandNum(n int) {
	if !IsTest() {
		return
	}

	testRandNum = n
}

// GetRandNum ランダム数を取得する
func GetRandNum(n int) int {
	if IsTest() {
		return testRandNum
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(n)
}

// SetTestRandString テストランダム文字列を設定する
func SetTestRandString(s string) {
	if !IsTest() {
		return
	}

	testRandString = s
}

// GetRandString ランダムな文字列を取得する
func GetRandString(n int) string {
	if IsTest() {
		return testRandString
	}

	rs2Letters := beego.AppConfig.String("randKey")

	b := make([]byte, n)
	for i := range b {
		b[i] = rs2Letters[rand.Intn(len(rs2Letters))]
	}

	return string(b)
}
