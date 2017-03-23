package movie

import (
	"dotstamp_server/utils"
	"log"
	"os/exec"
)

// Upload アップロード
type Upload struct {
	UserContributionID string
	Title              string
	Description        string
	CategoryID         string
	VideoStatus        string
}

// getRootPath パスを取得する
func getRootPath() (string, error) {
	p, err := utils.GetAppPath()
	if err != nil {
		return "", err
	}

	return p + "/", nil
}

// Make 作成する
func Make(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	movie := path + "static/movie/input.mp4"
	sound := path + "static/files/tmp/sound/" + file + ".m4a"
	dist := path + "static/files/movie/" + file + ".mp4"

	cmd := "ffmpeg -y -i " + movie + " -i " + sound + " -map 0:0 -map 1:0 -movflags faststart -vcodec libx264 -acodec copy " + dist

	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}

// UploadYoutube Youtubeにアップロードする
func UploadYoutube(u Upload) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	title := "-title='" + u.Title + "'"
	description := "-description='" + u.Description + "'"
	categoryID := "-categoryId='" + u.CategoryID + "'"
	videoStatus := "-videoStatus='" + u.VideoStatus + "'"
	cachetoken := "-cachetoken=false"
	option := title + " " + description + " " + categoryID + " " + videoStatus + " " + cachetoken

	file := path + "static/files/movie/" + u.UserContributionID + ".mp4"

	cmd := path + "tasks/youtubeUpload/youtubeUpload " + option + " youtube " + file

	if utils.IsTest() {
		log.Println(cmd)
		return nil
	}

	return exec.Command("sh", "-c", cmd).Start()
}
