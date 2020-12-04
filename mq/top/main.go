package top

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-youku/models"
	"go-youku/services/mq"
	redisClient "go-youku/services/redis"
	"strconv"
)

func main() {
	_ = beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)
	mq.Consumer("", "youku_top", callback)
}

func callback(s string) {
	type Data struct {
		VideoId int
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	videoInfo, err := models.RedisGetVideoInfo(data.VideoId)
	if err == nil {
		conn := redisClient.PoolConnect()
		defer conn.Close()
		redisChannelKey := "video:top:channel:channelId:" + strconv.Itoa(videoInfo.ChannelId)
		redisTypeKey := "video:top:type:typeId:" + strconv.Itoa(videoInfo.TypeId)
		conn.Do("zincrby", redisChannelKey, 1, data.VideoId)
		conn.Do("zincrby", redisTypeKey, 1, videoInfo)
	}
	fmt.Printf("msg is : %s\n", s)
}
