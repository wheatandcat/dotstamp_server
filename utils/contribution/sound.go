package contributions

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/models/csv_models"
	"github.com/wheatandcat/dotstamp_server/utils"
	"github.com/wheatandcat/dotstamp_server/utils/sound"

	"github.com/astaxie/beego"
)

// GetSoundByUserContributionID 投稿IDから音声を取得する
func GetSoundByUserContributionID(uID int) (models.UserContributionSound, error) {
	u := models.UserContributionSound{}

	r, _, err := u.GetByUserContributionID(uID)

	return r, err
}

// GetSoundListByUserContributionIDList 投稿IDリストから音声リストを取得する
func GetSoundListByUserContributionIDList(uID []int) ([]models.UserContributionSound, error) {
	u := models.UserContributionSound{}

	r, _, err := u.GetListByUserContributionIDList(uID)

	return r, err
}

// GetSoundMapByUserContributionIDList 投稿IDリストから音声マップを取得する
func GetSoundMapByUserContributionIDList(uID []int) (map[int]models.UserContributionSound, error) {
	m := map[int]models.UserContributionSound{}

	list, err := GetSoundListByUserContributionIDList(uID)
	if err != nil {
		return m, err
	}

	for _, v := range list {
		m[v.UserContributionID] = v
	}

	return m, nil
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
	var err error

	if b.TalkType <= models.TalkTypeText {
		s, err = ReplaceBodeySound(b.Body)
		if err != nil {
			return err
		}
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
	u.MakeStatus = models.MakeStatusMade

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
	u.MakeStatus = models.MakeStatusMade

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

func getFileName(u models.UserContributionSoundDetail) string {
	return strconv.Itoa(u.UserContributionID) + "_" + strconv.Itoa(u.Priority)
}

// AddTmpSound 一時音声ファイルを追加する
func AddTmpSound(u models.UserContributionSoundDetail) error {
	file := getFileName(u)

	return sound.AddTmpSound(u.BodySound, file, u.VoiceType)
}

// MakeSoundFile 音声ファイルを作成する
func MakeSoundFile(uID int, list []models.UserContributionSoundDetail) error {
	fileList := []string{}

	for _, u := range list {
		if u.BodySound == "" {
			continue
		}

		if u.MakeStatus != models.MakeStatusMade {
			if err := AddTmpSound(u); err != nil {
				return err
			}
		}

		fileList = append(fileList, getFileName(u))
	}

	return sound.Join(fileList, strconv.Itoa(uID))
}

// ExistsSound 音声ファイルの存在判定する
func ExistsSound(uID int) bool {
	dir := beego.AppConfig.String("soundDir")
	root, _ := utils.GetAppPath()

	return utils.ExistsFile(root + "/" + dir + strconv.Itoa(uID) + ".mp3")
}

// UpdateSoundToMakeStatus 投稿IDから作成状態を更新する
func UpdateSoundToMakeStatus(uID int, makeStatus int) error {
	u := models.UserContributionSoundDetail{}

	return u.UpdateToMakeStatusByUserContributionID(uID, makeStatus)
}

// UpdatesSoundToMakeStatusAndVoiceTypeByUserContributionID 投稿IDから作成状態をとボイスタイプを更新する
func UpdatesSoundToMakeStatusAndVoiceTypeByUserContributionID(uID int, makeStatus int, voiceType int) error {
	u := models.UserContributionSoundDetail{}

	return u.UpdatesToMakeStatusAndVoiceTypeByUserContributionID(uID, makeStatus, voiceType)
}

// ReplaceBodeySound 音声本文を置き換える
func ReplaceBodeySound(s string) (string, error) {
	s = getBodySoundFormat(s)

	c := csvModels.ContributionSoundBodyReplace{}
	list, err := c.GetStructAll()
	if err != nil {
		return "", err
	}

	for _, v := range list {
		s = strings.Replace(s, v.Text, v.Replace, -1)
	}

	return s, nil
}

// GetSoundDetailListByMakeStatusMade 作成済みの音声詳細を取得する
func GetSoundDetailListByMakeStatusMade() ([]models.UserContributionSoundDetail, error) {
	u := models.UserContributionSoundDetail{}
	r, _, err := u.GetListByMakeStatusMade()

	return r, err
}

// GetSoudDetailListBySpecifiedDays 指定に日数内の音声詳細を取得する
func GetSoudDetailListBySpecifiedDays(list []models.UserContributionSoundDetail, day int) []models.UserContributionSoundDetail {
	limit := utils.Now().Add(-1 * time.Duration(day) * 24 * time.Hour).Unix()
	r := []models.UserContributionSoundDetail{}

	for _, v := range list {

		if v.UpdatedAt.Unix() > limit {
			continue
		}

		r = append(r, v)
	}

	return r
}

// AddOrSaveSoundLength 長さを追加 or 保存する
func AddOrSaveSoundLength(uID int, second int, length int) error {
	u := models.UserContributionSoundLength{}
	u, _, err := u.GetByUserContributionID(uID)
	if err != nil {
		return err
	}

	if u.ID == uint(0) {
		u = models.UserContributionSoundLength{
			UserContributionID: uID,
			Second:             second,
			Length:             length,
		}

		return u.Add()
	}

	if u.Second == second {
		return nil
	}

	u.UserContributionID = uID
	u.Second = second
	u.Length = length
	return u.Save()
}
