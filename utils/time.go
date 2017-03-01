package utils

import "time"

var now time.Time

// Now 現在日時を取得する
func Now() time.Time {
	if IsTest() {
		return now
	}

	return time.Now()
}

// SetNow 現在日時を設定する
func SetNow(n time.Time) {
	now = n
}
