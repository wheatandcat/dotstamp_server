package sound

import (
	"dotstamp_server/utils"
	"os/exec"
)

// getRootPath パスを取得する
func getRootPath() string {
	return utils.GetAppPath() + "/../"
}

// AddTmpSound 一時音声を追加する
func AddTmpSound(text string, file string, v string) error {
	return add(text, "tmp/sound/"+file, v)
}

// add 追加する
func add(text string, file string, v string) error {
	path := getRootPath()

	dic := path + "tool/open-jtalk/dic/"
	voice := path + "tool/open-jtalk/voice/mei/" + v
	output := path + "static/files/" + file + ".wav"

	cmd := "echo " + text + " | open_jtalk -x " + dic + " -m " + voice + " -ow " + output

	_, err := exec.Command("sh", "-c", cmd).Output()

	return err
}

// Join 結合する
func Join(list []string, file string) error {
	cmd := "sox"
	path := getRootPath()

	for _, v := range list {
		cmd += " " + path + "static/files/tmp/sound/" + v + ".wav"
	}

	cmd += " " + path + "static/files/tmp/sound/" + file + ".wav"

	if _, err := exec.Command("sh", "-c", cmd).Output(); err != nil {
		return err
	}

	return toMp3(file)
}

// toMp3 mp3を変換する
func toMp3(file string) error {
	path := getRootPath()

	src := path + "static/files/tmp/sound/" + file + ".wav"
	dest := path + "static/files/sound/" + file + ".mp3"

	cmd := "lame -V2 " + src + " " + dest

	_, err := exec.Command("sh", "-c", cmd).Output()

	return err
}
