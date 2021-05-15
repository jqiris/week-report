package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var (
	config = new(Config)
	logger = logrus.WithField("package", "conf")
)

func InitConf(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error("read file: %v error:%v", filename, err)
		return err
	}
	err = json.Unmarshal(content, config)
	if err != nil {
		logger.Error("decode json error: %v", err)
		return err
	}
	logger.Warnf("the conf is:%+v", config)
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
