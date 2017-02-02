package controllersWork

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils"
)

// WorkMakeController 作品作成Controller
type WorkMakeController struct {
	controllers.BaseController
}

// Post 一時作品作成する
func (w *WorkMakeController) Post() {
	uID := 1
	name := w.GetString("name")
	categoryID, _ := w.GetInt("category_id")
	authorID, _ := w.GetInt("author_id")
	countryID, _ := w.GetInt("country_id")
	releaseDate, _ := utils.StringToDate(w.GetString("release_date"))

	r := map[string]interface{}{
		"user_id":      uID,
		"name":         name,
		"category_id":  categoryID,
		"author_id":    authorID,
		"country_id":   countryID,
		"release_date": releaseDate,
	}

	w.Data["json"] = r
	w.ServeJSON()
}
