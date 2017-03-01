package contributions

import "dotstamp_server/models"

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

// AddSoundDetail 音声詳細を追加する
func AddSoundDetail(uID int, b GetBody) error {
	s := ""
	if b.TalkType == models.TalkTypeText {
		s = b.Body
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
