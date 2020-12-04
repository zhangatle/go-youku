package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go-youku/services/mq"
)

func main() {
	_ = beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)
	mq.ConsumerDlx("youku.comment.count", "youku_comment_count", "youku_comment.count.dlx", "youku_comment_count_dlx", 10000, callback)
}

func callback(s string) {
	type Data struct {
		VideoId int
		EpisodesId int
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	if err == nil {
		o := orm.NewOrm()
		// 修改视频的总评论数
		_, _ = o.Raw("update video set comment=comment+1 where id=?", data.VideoId).Exec()
		// 修改剧集的评论数
		_, _ = o.Raw("update video_episodes set comment=comment+1 where id=?", data.EpisodesId).Exec()
		videoObj := map[string]int{
			"VideoId": data.VideoId,
		}
		videoJson, _ := json.Marshal(videoObj)
		_ = mq.Publish("", "youku_top", string(videoJson))
	}
	fmt.Printf("msg is : %s\n", s)
}
