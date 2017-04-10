package sound

import (
	"dotstamp_server/models/csv_models"
	"dotstamp_server/utils"
	"log"
	"os"
	"os/exec"
	"strconv"
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
func AddTmpSound(text string, file string, voiceType int) error {
	v := csvModels.VoiceType{}
	voice, err := v.GetStruct(voiceType)
	if err != nil {
		return err
	}

	if voice.VoiceSystemType == "" {
		return addOpenJtalk(text, "tmp/sound/"+file, "mei/mei_normal.htsvoice")
	}

	voiceSystemType, err := strconv.Atoi(voice.VoiceSystemType)
	if err != nil {
		return err
	}

	switch voiceSystemType {
	case csvModels.VoiceSystemTypeOpenjtalk:
		return addOpenJtalk(text, "tmp/sound/"+file, voice.VoiceType)
	case csvModels.VoiceSystemTypeAquesTalk:
		return addAquesTalk(text, "tmp/sound/"+file)
	default:
		return addOpenJtalk(text, "tmp/sound/"+file, "mei/mei_normal.htsvoice")
	}
}

// addOpenJtalk OpenJtalkを追加する
func addOpenJtalk(text string, file string, v string) error {
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

func toAqk2k(text string) (string, error) {
	path, err := getRootPath()
	if err != nil {
		return "", err
	}

	voice := path + "tool/aqk2k/Kanji2KoeCmd"
	dic := path + "tool/aqk2k/aq_dic"

	text = strings.Replace(text, "\n", "。", -1)
	text = strings.Replace(text, "'", "", -1)
	text = strings.Replace(text, "・", "、", -1)

	cmd := "echo '" + text + "' | " + voice + " " + dic

	r, err := exec.Command("sh", "-c", cmd).Output()

	return string(r), err
}

// addSoundless 無音を追加する
func addSoundless(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	output := path + "static/files/" + file + ".wav"
	soundless := path + "static/sound/soundless.wav"

	cmd := "sox " + output + " " + soundless + " " + output
	log.Println(cmd)
	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}

// addAquesTalk AquesTalkを追加する
func addAquesTalk(text string, file string) error {
	text, err := toAqk2k(text)
	if err != nil {
		return err
	}

	path, err := getRootPath()
	if err != nil {
		return err
	}

	voice := path + "tool/aques-talk/Talk"
	output := path + "static/files/" + file + ".wav"

	text = strings.Replace(text, "\n", "。", -1)
	text = strings.Replace(text, "'", "", -1)

	cmd := "echo '" + text + "' | " + voice + " > " + output

	_, err = exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return err
	}

	return addSoundless(file)
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

// RemoveDetailFile 詳細ファイルを削除する
func RemoveDetailFile(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	if err := os.Remove(path + "static/files/tmp/sound/" + file + ".wav"); err != nil {
		return err
	}

	return nil
}

// RemoveJoinFile 連結ファイルを削除する
func RemoveJoinFile(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	wav := path + "static/files/tmp/sound/" + file + ".wav"
	if err := os.Remove(wav); err != nil {
		return err
	}

	mp3 := path + "static/files/sound/" + file + ".mp3"
	if err := os.Remove(mp3); err != nil {
		return err
	}

	m4a := path + "static/files/tmp/sound/" + file + ".m4a"
	if err := os.Remove(m4a); err != nil {
		return err
	}

	return nil
}
