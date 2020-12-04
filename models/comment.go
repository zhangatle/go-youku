package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"go-youku/services/mq"
	"time"
)

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Stamp       int
	Status      int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId int, offset int, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	num, _ := o.Raw("select id from comment where status=1 and episodes_id=?", episodesId).QueryRows(&comments)
	_, err := o.Raw("select id, content, add_time, user_id, stamp, praise_count, episodes_id from comment where status=1 and episodes_id=? order by add_time desc limit ?,?", episodesId, offset, limit).QueryRows(&comments)
	return num, comments, err
}

func SaveComment(content string, uid int, episodesId int, videoId int) error {
	o := orm.NewOrm()
	var comment Comment
	comment.Content = content
	comment.UserId = uid
	comment.EpisodesId = episodesId
	comment.VideoId = videoId
	comment.Stamp = 0
	comment.Status = 1
	comment.AddTime = time.Now().Unix()
	_, err := o.Insert(&comment)
	if err == nil {
		_, _ = o.Raw("update video set comment=comment+1 where id=?", videoId).Exec()
		_, _ = o.Raw("update video_episodes set comment=comment+1 where id=?", episodesId).Exec()
		// 通过mq更新redis排行榜
		videoObj := map[string]int{
			"VideoId" : videoId,
		}
		videoJson, _ := json.Marshal(videoObj)
		_ = mq.Publish("", "youku_top", string(videoJson))
		// 延迟增加评论数
		videoCountObj := map[string]int{
			"VideoId" : videoId,
			"EpisodesId": episodesId,
		}
		videoCountJson, _ := json.Marshal(videoCountObj)
		_ = mq.PublishDlx("youku.comment.count", string(videoCountJson))
	}
	return err
}
