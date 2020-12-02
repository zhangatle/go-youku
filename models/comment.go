package models

import (
	"github.com/astaxie/beego/orm"
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
