// Copyright (c) 2021 zacksleo <zacksleo@gmail.com>
// MIT Licence - http://opensource.org/licenses/MIT

/**
* timestamp alfred wordflow
 */
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	aw "github.com/deanishe/awgo"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func getCurrentTime() {
	now := time.Now()
	nowSecond := fmt.Sprintf("%d", now.Unix())
	nowMillSecond := fmt.Sprintf("%d", now.UnixNano()/1000000)
	wf.NewItem(nowSecond).
		Subtitle("当前时间戳(s)").
		Valid(true).
		Arg(fmt.Sprintf("%d", now.Unix())).
		Var("timestamp", nowSecond)
	wf.NewItem(nowMillSecond).
		Subtitle("当前时间戳(ms)").
		Valid(true).
		Arg(nowMillSecond).
		Var("timestamp", nowMillSecond)
	timeStr := now.Format("2006-01-02 15:04:05")
	wf.NewItem(timeStr).Subtitle("当前时间").Valid(true).Var("timestamp", timeStr)

	wf.SendFeedback()
}

func timestampDecode(query string) {
	timestamp, err := strconv.Atoi(query)
	if err != nil {
		wf.NewItem("请重新输入").Subtitle("解析错误")
		wf.SendFeedback()
		return
	}

	if len(query) == 10 {
		currentTime := time.Unix(int64(timestamp), 0)
		timeStr := currentTime.Format("2006-01-02 15:04:05")
		wf.NewItem(timeStr).Subtitle("解析后的时间").Valid(true).Var("timestamp", timeStr)
		wf.SendFeedback()
		return
	}

	if len(query) == 13 {
		currentTime := time.Unix(int64(timestamp/1000), 0)
		timeStr := currentTime.Format("2006-01-02 15:04:05")
		wf.NewItem(timeStr).Subtitle("解析后的时间").Valid(true).Var("timestamp", timeStr)
		wf.SendFeedback()
		return
	}

	wf.NewItem("请重新输入").Subtitle("格式不正确")

	wf.SendFeedback()
}

func timestampEncode(query string) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", query, loc)
	timeStr := fmt.Sprintf("%d", tt.Unix())
	wf.NewItem(timeStr).Subtitle("解析后的时间戳").Valid(true).Var("timestamp", timeStr)
	wf.SendFeedback()
}

func help() {
	wf.NewItem("ts").Subtitle("查询当前日期和时间戳")
	wf.NewItem("ts help").Subtitle("查询帮助")
	wf.NewItem("ts {1616514467}").Subtitle("将时间戳转化为日期")
	wf.NewItem("ts {1616514689011}").Subtitle("将时间戳转化为日期")
	wf.NewItem("ts {2006-01-02 15:04:05}").Subtitle("将日期转化为时间戳")
	wf.SendFeedback()
}

func run() {

	query := ""
	if len(wf.Args()) > 0 {
		query = wf.Args()[0]
	}

	if query == `help` {
		help()
		return
	}

	// 默认展示当前时间戳
	if len(query) < 1 {
		getCurrentTime()
		return
	}

	re := regexp.MustCompile(`^\d{10,}$`)
	if re.Match([]byte(query)) {
		timestampDecode(query)
		return
	}

	re = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)
	if re.Match([]byte(query)) {
		timestampEncode(query)
		return
	}

	wf.NewItem("格式不正确").Subtitle("请重新输入")

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
