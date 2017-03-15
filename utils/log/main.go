package logs

import (
	"dotstamp_server/utils"
	"io"
	"log"
	"os"
)

// LogFile ログファイルを開く
func LogFile(file string) (o *os.File, err error) {
	apppath, err := utils.GetAppPath()
	if err != nil {
		return o, err
	}
	log.Println(apppath)

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
	apppath, err := utils.GetAppPath()
	if err != nil {
		return err
	}

	return os.Remove(apppath + "/..//logs/" + file + ".log")
}
