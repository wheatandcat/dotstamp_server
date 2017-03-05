package logs

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// getAppPath アプリケーションパスを取得する
func getAppPath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.."+string(filepath.Separator))))

	return apppath
}

// LogFile ログファイルを開く
func LogFile(file string) (*os.File, error) {
	apppath := getAppPath()

	return os.OpenFile(apppath+"/logs/"+file+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

// Err エラーを出力する
func Err(v interface{}, userID int) error {
	return output("error", v, userID)
}

// Batch 実行を出力する
func Batch(v interface{}, batchName string) error {
	return output("batch", v, batchName)
}

// Output 出力する
func output(file string, v interface{}, t interface{}) error {
	logfile, err := LogFile(file)
	if err != nil {
		return err
	}
	defer logfile.Close()

	log.SetOutput(io.MultiWriter(logfile))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(v, "[", t, "]")

	return nil
}

// RemoveLogFile ログファイル削除する
func RemoveLogFile(file string) error {
	apppath := getAppPath()

	return os.Remove(apppath + "/logs/" + file + ".log")
}
