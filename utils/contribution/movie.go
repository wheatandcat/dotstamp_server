package contributions

import (
	"dotstamp_server/models"
)

// AddMovie 動画を追加する
func AddMovie(uID int, mID string, t int, s int) error {
	u := models.UserContributionMovie{
		UserContributionID: uID,
		MovieID:            mID,
		MovieType:          t,
		MovieStatus:        s,
	}

	return u.Add()
}

// AddOrSaveMovie 追加か保存する
func AddOrSaveMovie(uID int, mID string, t int, s int) error {
	u, err := GetMovie(uID, t)
	if err != nil {
		return err
	}

	if u.ID == uint(0) {
		return AddMovie(uID, mID, models.MovieTypeYoutube, s)
	}

	u.MovieID = mID
	u.MovieStatus = s

	return u.Save()
}

// GetMovie 動画を取得する
func GetMovie(uID int, t int) (models.UserContributionMovie, error) {
	u := models.UserContributionMovie{}
	r, _, err := u.GetByUserContributionID(uID, t)

	return r, err
}
