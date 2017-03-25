package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"strconv"

	"github.com/astaxie/beego"
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

// ExistsMovie 動画ファイルの存在判定する
func ExistsMovie(uID int) bool {
	dir := beego.AppConfig.String("movieDir")
	root, _ := utils.GetAppPath()

	return utils.ExistsFile(root + "/" + dir + strconv.Itoa(uID) + ".mp4")
}
