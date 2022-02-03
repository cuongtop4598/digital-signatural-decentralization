package utils

import (
	"encoding/json"
	"io/ioutil"
)

func Json2struct(jsonfile string, dst interface{}) error {
	file, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(file), &dst)
	if err != nil {
		return err
	}
	return nil
}
