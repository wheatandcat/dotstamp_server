package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"dotstamp_server/utils/sound"
	"errors"
	"regexp"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	// VoiceTypeMeiNormal 音声タイプ:mei_normal
	VoiceTypeMeiNormal = 1
	// VoiceTypeMeiAngry 音声タイプ:mei_angry
	VoiceTypeMeiAngry = 2
	// VoiceTypeMeiBashful 音声タイプ:mei_bashful
	VoiceTypeMeiBashful = 3
	// VoiceTypeMeiHappy 音声タイプ:mei_happy
	VoiceTypeMeiHappy = 4
	// VoiceTypeMeiSad 音声タイプ:mei_sad
	VoiceTypeMeiSad = 5
	// VoiceTypeM100 音声タイプ:m100
	VoiceTypeM100 = 6
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

	if b.TalkType <= models.TalkTypeText {
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

func getByID(id uint) (models.UserContributionSoundDetail, error) {
	u := models.UserContributionSoundDetail{}
	r, _, err := u.GetByID(id)

	if r.ID == uint(0) {
		return r, errors.New("not found ID")
	}

	return r, err
}

// SaveSoundDetailToBodySound 音声本文を保存する
func SaveSoundDetailToBodySound(id uint, body string, userID int) error {
	u, err := getByID(id)
	if err != nil {
		return err
	}

	c, err := GetByUserContributionID(u.UserContributionID)
	if err != nil {
		return err
	}

	if userID != c.UserID {
		return errors.New("difference UserID")
	}

	u.BodySound = body

	return u.Save()
}

// SaveSoundDetailTVoiceType ボイスタイプを保存する
func SaveSoundDetailTVoiceType(id uint, v int, userID int) error {
	u, err := getByID(id)
	if err != nil {
		return err
	}

	c, err := GetByUserContributionID(u.UserContributionID)
	if err != nil {
		return err
	}

	if userID != c.UserID {
		return errors.New("difference UserID")
	}

	u.VoiceType = v

	return u.Save()
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
		return "mei/mei_normal.htsvoice"
	case VoiceTypeMeiAngry:
		return "mei/mei_angry.htsvoice"
	case VoiceTypeMeiBashful:
		return "mei/mei_bashful.htsvoice"
	case VoiceTypeMeiHappy:
		return "mei/mei_happy.htsvoice"
	case VoiceTypeMeiSad:
		return "mei/mei_sad.htsvoice"
	case VoiceTypeM100:
		return "m100/nitech_jp_atr503_m001.htsvoice"
	default:
		return "mei/mei_normal.htsvoice"
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

// ExistsSound 音声ファイルの存在判定する
func ExistsSound(uID int) bool {
	dir := beego.AppConfig.String("soundDir")
	root, _ := utils.GetAppPath()

	return utils.ExistsFile(root + "/" + dir + strconv.Itoa(uID) + ".mp3")
}
