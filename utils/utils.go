package utils

import (
	"os"
	"time"
)

/**
判断文件是否存在
*/
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/**
创建文件夹
*/
func CreateDir(dirName string, mod os.FileMode) error {
	err := os.Mkdir(dirName, mod)
	if err != nil {
		return err
	}
	return nil
}

func TimeToDayBegin(t *time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TimeToDayEnd(t *time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}
