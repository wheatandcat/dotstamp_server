package csvModels

// ContributionSoundBodyReplace 音声テキスト置き換え
type ContributionSoundBodyReplace struct {
	ID      string
	Text    string
	Replace string
	Delete  string
}

// GetStructAll 全てを取得する
func (c *ContributionSoundBodyReplace) GetStructAll() (r []ContributionSoundBodyReplace, err error) {
	err = GetAll("contribution_sound_body_replace.csv", &r)

	return r, err
}
