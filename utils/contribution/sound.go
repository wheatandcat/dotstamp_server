package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/utils/sound"
	"regexp"
	"strconv"
)

const (
	// VoiceTypeMeiNormal 音声タイプ:mei_normal
	VoiceTypeMeiNormal = 1
)

// GetSoundByUserContributionID 投稿IDから音声を取得する
func GetSoundByUserContributionID(uID int) (models.UserContributionSound, error) {
	u := models.UserContributionSound{}

	r, _, err := u.GetByUserContributionID(uID)

	return r, err
}

// AddSound 音声を追加する
func AddSound(uID int, s int) error {
	u := models.UserContributionSound{
		UserContributionID: uID,
		SoundStatus:        s,
	}

	return u.Add()
}

// GetSoundDetailListByUserContributionID 投稿IDから音声リストを取得する
func GetSoundDetailListByUserContributionID(uID int) ([]models.UserContributionSoundDetail, error) {
	u := models.UserContributionSoundDetail{}

	r, _, err := u.GetListByUserContributionID(uID)

	return r, err
}

func getBodySoundFormat(str string) string {
	rep := regexp.MustCompile("[!-/:-@≠¥[-`{-~]")
	str = rep.ReplaceAllString(str, "")

	return str
}

// AddSoundDetail 音声詳細を追加する
func AddSoundDetail(uID int, b GetBody) error {
	s := ""

	if b.TalkType == models.TalkTypeText {
		s = getBodySoundFormat(b.Body)
	}

	u := models.UserContributionSoundDetail{
		UserContributionID: uID,
		Priority:           b.Priority,
		TalkType:           b.TalkType,
		Body:               b.Body,
		BodySound:          s,
		VoiceType:          b.Character.VoiceType,
	}

	return u.Add()
}

// AddSoundDetailList 音声詳細リストを追加する
func AddSoundDetailList(uID int, list []GetBody) error {
	for _, b := range list {
		if err := AddSoundDetail(uID, b); err != nil {
			return err
		}
	}

	return nil
}

func getVoiceTypeFile(voiceType int) string {
	switch voiceType {
	case VoiceTypeMeiNormal:
		return "mei_normal.htsvoice"
	default:
		return "mei_normal.htsvoice"
	}
}

func getFileName(u models.UserContributionSoundDetail) string {
	return strconv.Itoa(u.UserContributionID) + "_" + strconv.Itoa(u.Priority)
}

// MakeSoundFile 音声ファイルを作成する
func MakeSoundFile(uID int, list []models.UserContributionSoundDetail) error {
	fileList := []string{}

	for _, u := range list {
		if u.BodySound == "" {
			continue
		}

		file := getFileName(u)
		voice := getVoiceTypeFile(u.VoiceType)
		if err := sound.AddTmpSound(u.BodySound, file, voice); err != nil {
			return err
		}

		fileList = append(fileList, file)
	}

	return sound.Join(fileList, strconv.Itoa(uID))
}
