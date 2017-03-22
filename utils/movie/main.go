package movie

import (
	"dotstamp_server/utils"
	"os/exec"
)

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

	cmd := "ffmpeg -y -i " + movie + " -i " + sound + " -vcodec copy -acodec copy " + dist

	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}
