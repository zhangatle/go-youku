package send

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-youku/models"
	"go-youku/services/mq"
)

func main() {
	_ = beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)
	mq.Consumer("", "youku_send_message_user", callback)
}

func callback(s string) {
	type Data struct {
		UserId int
		MessageId int64
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	if err == nil {
		_ = models.SendMessageUser(data.UserId, data.MessageId)
	}
	fmt.Printf("msg is %s\n", s)
}