package core

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/jqiris/week-report/conf"
	"github.com/jqiris/week-report/filters"
	"github.com/jqiris/week-report/utils"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	logger = logrus.WithField("package", "cmd")
)

//Before 使用前初始化
func Before(c *cli.Context) error {
	//初始化配置
	cfg := c.String("conf")
	if len(cfg) == 0 {
		return errors.New("请指定配置文件名")
	}
	err := conf.InitConf(cfg)
	if err != nil {
		return err
	}
	//初始化输出目录
	output := conf.GetOutputConf()
	if !utils.IsDir(output.Dir) {
		return utils.CreateDir(output.Dir, 0755)
	}
	return nil
}

//Report 日报产出
func Report(c *cli.Context) error {
	s, e := c.Timestamp("sdate"), c.Timestamp("edate")
	if s == nil {
		n := time.Now().AddDate(0, 0, -4)
		s = &n
	}
	if e == nil {
		n := time.Now()
		e = &n
	}
	sb, ee := utils.TimeToDayBegin(s), utils.TimeToDayEnd(e)
	sDate, eDate := sb.Format("2006-01-02 15:04:05"), ee.Format("2006-01-02 15:04:05")
	osDate, oeDate := sb.Format("20060102"), ee.Format("20060102")
	logger.Info("week-report:", osDate, "==>", oeDate)
	//获取当前路径
	cDir, err := os.Getwd()
	if err != nil {
		return err
	}
	//分析日志
	result := make(map[string][]string)
	projects := conf.GetProjectConf()
	for pName, pDirs := range projects {
		logger.Info("开始分析日志, 项目:" + pName)
		for _, pDir := range pDirs {
			res, err := analyse(cDir, pDir, pName, sDate, eDate)
			if err != nil {
				return err
			}
			if len(res) > 0 {
				result[pName] = append(result[pName], res...)
			}
		}
		logger.Info("结束分析日志, 项目:" + pName)
	}
	//去重并且过滤关键字
	for k, v := range result {
		result[k] = filters.Filtering(v)
	}
	//输出到目录
	oCfg := conf.GetOutputConf()
	return output(oCfg.Title, oCfg.Dir, osDate, oeDate, result)
}

//analyse 解析日报信息
func analyse(cDir, pDir, name, sDate, eDate string) ([]string, error) {
	err := os.Chdir(pDir)
	if err != nil {
		return nil, err
	}
	logCmd := cmd.NewCmd("git", "log", "--author="+conf.GetUserConf(), "--pretty=format:%s", "--after="+sDate, "--before="+eDate)
	s := <-logCmd.Start()
	err = os.Chdir(cDir)
	if err != nil {
		return nil, err
	}
	return s.Stdout, nil
}

//output 输出到目录文件
func output(title, dir, sDate, eDate string, result map[string][]string) error {
	//开始写入结果
	filename := "week_" + sDate + "_" + eDate + ".txt"
	filepath := path.Join(dir, filename)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if len(title) > 0 {
		file.Write([]byte(title + ":\n"))
	}
	num := 1
	for pName, plist := range result {
		for _, item := range plist {
			vw := fmt.Sprintf("%d-%s-%s\n", num, pName, item)
			file.Write([]byte(vw))
			num += 1
		}
	}
	return nil
}
