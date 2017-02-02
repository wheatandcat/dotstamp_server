package work

import (
	"github.com/mitchellh/mapstructure"

	"dotstamp_server/models"
)

// AddPerson 人物を追加する
func AddPerson(m map[string]interface{}) error {
	tp := models.TmpPerson{}

	err := mapstructure.Decode(m, &tp)
	if err != nil {
		return err
	}

	return tp.Add()
}

// AddWork 作品を追加する
func AddWork(m map[string]interface{}) error {
	tw := models.TmpWork{}

	err := mapstructure.Decode(m, &tw)
	if err != nil {
		return err
	}

	return tw.Add()
}
