package main

import (
	"log"
	"os"

	"github.com/jqiris/week-report/core"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "week-report",
		Usage: "通过git使用记录产生周报",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "conf",
				Aliases: []string{"c"},
				Value:   "config.json",
				Usage:   "指定配置文件",
			},
			&cli.TimestampFlag{
				Name:    "sdate",
				Aliases: []string{"s"},
				Layout:  "20060102",
				Usage:   "日报开始日期",
			},
			&cli.TimestampFlag{
				Name:    "edate",
				Aliases: []string{"e"},
				Layout:  "20060102",
				Usage:   "日报结束日期",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "产生周报",
				Before:  core.Before,
				Action:  core.Report,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
