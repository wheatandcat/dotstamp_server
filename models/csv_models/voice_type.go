package csvModels

import "strconv"

// VoiceType 音声タイプ
type VoiceType struct {
	ID              string
	Name            string
	VoiceSystemType string
	VoiceType       string
	Delete          string
}

const (
	// VoiceSystemTypeOpenjtalk タイプ：open-jtalk
	VoiceSystemTypeOpenjtalk = 1
	// VoiceSystemTypeAquesTalk タイプ：AquesTalk
	VoiceSystemTypeAquesTalk = 2
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
	// VoiceTypeYukkuri 音声タイプ:ゆっくり
	VoiceTypeYukkuri = 7
)

// GetStructAll 全てを取得する
func (c *VoiceType) GetStructAll() (r []VoiceType, err error) {
	err = GetAll("voice_type.csv", &r)

	return r, err
}

// GetStruct 取得する
func (c *VoiceType) GetStruct(voiceType int) (r VoiceType, err error) {
	list, err := c.GetStructAll()
	if err != nil {
		return r, err
	}

	for _, v := range list {
		id, err := strconv.Atoi(v.ID)
		if err != nil {
			return r, err
		}

		if id == voiceType {
			return v, nil
		}
	}

	return r, nil

}
