package main

import (
	"dotstamp_server/tasks"
	"dotstamp_server/utils/contribution"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func main() {
	sm := create()

	contribution(sm)

	sm.Finalize().PingSearchEngines()
}

// create 作成する
func create() *stm.Sitemap {
	sm := stm.NewSitemap()
	sm.SetDefaultHost(beego.AppConfig.String("topurl"))
	// sm.SetCompress(true)
	// sm.SetVerbose(true)
	sm.SetVerbose(false)
	sm.Create()

	sm.Add(stm.URL{"loc": "/about", "changefreq": "daily"})

	return sm
}

// contribution 投稿を確認する
func contribution(sm *stm.Sitemap) {
	contributionIDList, err := contributions.GetViewStatusPublicIDList()
	if err != nil {
		tasks.Err(err, "sitemap")
	}

	for _, id := range contributionIDList {
		sm.Add(stm.URL{"loc": "/contribution/show/" + strconv.Itoa(id), "changefreq": "daily"})
	}
}
