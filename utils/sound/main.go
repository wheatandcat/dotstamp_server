package sound

import (
	"dotstamp_server/utils"
	"os/exec"
	"strings"
)

// getRootPath パスを取得する
func getRootPath() (string, error) {
	p, err := utils.GetAppPath()
	if err != nil {
		return "", err
	}

	return p + "/", nil
}

// AddTmpSound 一時音声を追加する
func AddTmpSound(text string, file string, v string) error {
	return add(text, "tmp/sound/"+file, v)
}

// add 追加する
func add(text string, file string, v string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	dic := path + "tool/open-jtalk/dic/"
	voice := path + "tool/open-jtalk/voice/" + v
	output := path + "static/files/" + file + ".wav"

	text = strings.Replace(text, "\n", "。", -1)

	cmd := "echo '" + text + "' | open_jtalk -x " + dic + " -m " + voice + " -ow " + output

	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}

// Join 結合する
func Join(list []string, file string) error {
	cmd := "sox"
	path, err := getRootPath()
	if err != nil {
		return err
	}

	for _, v := range list {
		cmd += " " + path + "static/files/tmp/sound/" + v + ".wav"
	}

	cmd += " " + path + "static/files/tmp/sound/" + file + ".wav"

	if _, err := exec.Command("sh", "-c", cmd).Output(); err != nil {
		return err
	}

	return toMp3(file)
}

// toMp3 mp3に変換する
func toMp3(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	src := path + "static/files/tmp/sound/" + file + ".wav"
	dest := path + "static/files/sound/" + file + ".mp3"

	cmd := "lame -V2 " + src + " " + dest

	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}

// ToM4a m4aに変換する
func ToM4a(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	src := path + "static/files/sound/" + file + ".mp3"
	dest := path + "static/files/tmp/sound/" + file + ".m4a"

	cmd := "ffmpeg -y -i " + src + " -vn -ac 2 -vol 256 -ab 112k " + dest

	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}
