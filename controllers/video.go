package controllers

import (
	"github.com/astaxie/beego"
	"go-youku/models"
)

type VideoController struct {
	beego.Controller
}

func (c *VideoController) ChannelAdvert() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	num, adverts, err := models.GetChannelAdvert(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "查询失败,请稍后重试")
		c.ServeJSON()
	}else{
		c.Data["json"] = ReturnSuccess(0, "success", adverts, num)
		c.ServeJSON()
	}
}