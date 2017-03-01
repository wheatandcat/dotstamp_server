package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

// getAppPath アプリケーションパスを取得する
func getAppPath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))

	return apppath
}

// LogFile ログファイル
func LogFile(file string) (*os.File, error) {
	apppath := getAppPath()

	return os.OpenFile(apppath+"/logs/"+file+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}
