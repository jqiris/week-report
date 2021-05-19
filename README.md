# week-report
使用git记录生成周报工具


## 安装

### 方式一：源码安装
```
git clone https://github.com/jqiris/week-report 
cd week-report & go build & go install
```

### 方式二：直接下载编译文件


## 使用

### 调用week-report命令，确保出现以下内容，保证工具已经正确安装
```
NAME:
   week-report - 通过git使用记录产生周报

USAGE:
   week-report.exe [global options] command [command options] [arguments...]

COMMANDS:
   run, r   产生周报
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --conf value, -c value   指定配置文件 (default: "config.json")
   --sdate value, -s value  日报开始日期 (default: (*time.Time)(nil))
   --edate value, -e value  日报结束日期 (default: (*time.Time)(nil))
   --help, -h               show help (default: false)
```

### 选择一个周报目录，然后设置配置文件config.json，范例如下:
```
{
    "user": "jqiris", 
    "project": {
        "周报项目": [
            "H:\\project\\week-report"
        ]
    },
    "output": {
        "title": "上周",
        "dir": "report",
        "filter": [
            "test",
            "Merge"
        ]
    }
}
```
格式说明:
- user-git账号名称
- project-项目目录，同一个项目可能有多个子项目，产出按照项目名称归类
- output-输出设置
    - title-周报标题
    - dir-周报输出目录，如果没有，尝试自动生成
    - filter-如果git记录里面包含这些关键字会过滤掉

### 执行周报生成命令

- 标准：week-report run ,根据配置文件信息自动生成周报，默认开始时间5天前，结束时间当前时间
- 指定配置文件：week-report -c=xxx.json run ,xxx.json配置文件路径
- 指定开始日期或者结束日期 week-report -s=20210518 -e=20210519 run 日期可选，不设置默认开始时间5天前，结束时间当天


## 说明
命令执行完，会在输出目录生成周报文件，格式week_开始日期_结束日期.txt，同一个项目提交的同一个记录会去重，遇到过滤关键字字段会忽略，
所以如果你的git提交的时候做的的是同一个功能，可以用同一个注释提交，如果你不想日志加入周报生成，可以在记录里加入过滤关键字，生成效果如下，希望大家用的开心，并给我点个赞：

```
上周:
1-周报项目-Initial commit
2-周报项目-周报工具
```