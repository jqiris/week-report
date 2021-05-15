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

//日报产出
func Report(c *cli.Context) error {
	s, e, n := c.Timestamp("sdate"), c.Timestamp("edate"), time.Now()
	if s == nil {
		s = &n
	}
	if e == nil {
		e = &n
	}
	sb, ee := utils.TimeToDayBegin(s), utils.TimeToDayEnd(e)
	sdate, edate := sb.Format("2006-01-02 15:04:05"), ee.Format("2006-01-02 15:04:05")
	osdate, oedate := sb.Format("20060102"), ee.Format("20060102")
	logger.Info("week-report:", osdate, "==>", oedate)
	//获取当前路径
	cdir, err := os.Getwd()
	if err != nil {
		return err
	}
	//分析日志
	result := make(map[string][]string)
	projects := conf.GetProjectConf()
	for pname, pdirs := range projects {
		logger.Info("开始分析日志, 项目:" + pname)
		for _, pdir := range pdirs {
			res, err := analyse(cdir, pdir, pname, sdate, edate)
			if err != nil {
				return err
			}
			if len(res) > 0 {
				result[pname] = append(result[pname], res...)
			}
		}
		logger.Info("结束分析日志, 项目:" + pname)
	}
	//去重并且过滤关键字
	for k, v := range result {
		result[k] = filters.Filtering(v)
	}
	//输出到目录
	ocfg := conf.GetOutputConf()
	return output(ocfg.Title, ocfg.Dir, osdate, oedate, result)
}

//analyse 解析日报信息
func analyse(cdir, pdir, name, sdate, edate string) ([]string, error) {
	err := os.Chdir(pdir)
	if err != nil {
		return nil, err
	}
	logcmd := cmd.NewCmd("git", "log", "--author="+conf.GetUserConf(), "--pretty=format:%s", "--after="+sdate, "--before="+edate)
	s := <-logcmd.Start()
	err = os.Chdir(cdir)
	if err != nil {
		return nil, err
	}
	return s.Stdout, nil
}

//output 输出到目录文件
func output(title, dir, sdate, edate string, result map[string][]string) error {
	//开始写入结果
	filename := "week_" + sdate + "_" + edate + ".txt"
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
	for pname, plist := range result {
		for _, item := range plist {
			vw := fmt.Sprintf("%d-%s-%s\n", num, pname, item)
			file.Write([]byte(vw))
			num += 1
		}
	}
	return nil
}
