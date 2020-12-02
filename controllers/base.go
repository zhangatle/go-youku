package controllers

import (
	"github.com/astaxie/beego"
	"go-youku/models"
)

type BaseController struct {
	beego.Controller
}

// @router /channel/region [get]
func (c *BaseController) ChannelRegion() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	num, regions, err := models.GetChannelRegion(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", regions, num)
		c.ServeJSON()
	}
}

// @router /channel/type [get]
func (c *BaseController) ChannelType() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	num, regions, err := models.GetChannelType(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", regions, num)
		c.ServeJSON()
	}
}
