package controllers

import (
	"github.com/astaxie/beego"
	"go-youku/models"
)

type TopController struct {
	beego.Controller
}

// @router /type/top [get]
func (c *TopController) TypeTop() {
	typeId, _ := c.GetInt("typeId")
	if typeId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定类型")
		c.ServeJSON()
	}
	num, videos, err := models.GetTypeTop(typeId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	}else{
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	}
}

// @router /channel/top [get]
func (c *TopController) ChannelTop() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	num, videos, err := models.RedisGetChannelTop(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	}else{
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	}
}