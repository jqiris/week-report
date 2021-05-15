package conf

import (
	"encoding/json"
	"io/ioutil"
)

var (
	config = new(Config)
)

func InitConf(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, config)
	if err != nil {
		return err
	}
	return nil
}

func GetUserConf() string {
	return config.User
}

func GetProjectConf() map[string][]string {
	return config.Project
}

func GetOutputConf() OutputConf {
	return config.Output
}
