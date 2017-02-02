package controllersWork

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils"
)

// PersonMakeController 人物作成コントローラ
type PersonMakeController struct {
	controllers.BaseController
}

// Post 人物作成する
func (p *PersonMakeController) Post() {
	uID := 1
	name := p.GetString("name")
	categoryID, _ := p.GetInt("category_id")
	authorID, _ := p.GetInt("author_id")
	countryID, _ := p.GetInt("country_id")
	releaseDate, _ := utils.StringToDate(p.GetString("release_date"))

	w := map[string]interface{}{
		"user_id":      uID,
		"name":         name,
		"category_id":  categoryID,
		"author_id":    authorID,
		"country_id":   countryID,
		"release_date": releaseDate,
	}

	p.Data["json"] = w
	p.ServeJSON()
}
