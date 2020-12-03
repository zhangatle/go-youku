package controllers

import (
	"github.com/astaxie/beego"
	"go-youku/models"
)

type VideoController struct {
	beego.Controller
}

// @router /channel/advert [get]
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
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", adverts, num)
		c.ServeJSON()
	}
}

// @router /channel/hot [get]
func (c *VideoController) ChannelHot() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	num, videos, err := models.GetChannelHot(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	}
}

// @router /channel/recommend/region [get]
func (c *VideoController) ChannelRecommendRegionList() {
	channelId, _ := c.GetInt("channelId")
	regionId, _ := c.GetInt("regionId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	if regionId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道地区")
		c.ServeJSON()
	}
	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	}
}

// @router /channel/recommend/type [get]
func (c *VideoController) GetChannelRecommendTypeList() {
	channelId, _ := c.GetInt("channelId")
	typeId, _ := c.GetInt("typeId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	if typeId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道类型")
		c.ServeJSON()
	}
	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	}
}

// @router /channel/video [get]
func (c *VideoController) ChannelVideo() {
	channelId, _ := c.GetInt("channelId")
	regionId, _ := c.GetInt("regionId")
	typeId, _ := c.GetInt("typeId")
	end := c.GetString("end")
	sort := c.GetString("sort")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}
	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	}
}

// @router /video/info [get]
func (c *VideoController) VideoInfo() {
	videoId, _ := c.GetInt("videoId")
	if videoId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频ID")
		c.ServeJSON()
	}
	video, err := models.RedisGetVideoInfo(videoId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", video, 1)
		c.ServeJSON()
	}
}

// @router /video/episodes/list [get]
func (c *VideoController) VideoEpisodesList() {
	videoId, _ := c.GetInt("videoId")
	if videoId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频ID")
		c.ServeJSON()
	}
	num, episodes, err := models.RedisGetVideoEpisodesList(videoId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		c.ServeJSON()
	}
}

// @router /user/video [get]
func (c *VideoController) UserVideo() {
	uid, _ := c.GetInt("uid")
	if uid == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定用户")
		c.ServeJSON()
	}
	num, videos, err := models.GetUserVideo(uid)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	}
}

// @router /video/save [*]
func (c *VideoController) VideoSave() {
	playUrl := c.GetString("playUrl")
	title := c.GetString("title")
	subTitle := c.GetString("subTitle")
	channelId, _ := c.GetInt("channelId")
	typeId, _ := c.GetInt("typeId")
	regionId, _ := c.GetInt("regionId")
	uid, _ := c.GetInt("uid")
	aliyunVideoId := c.GetString("aliyunVideoId")
	if uid == 0 {
		c.Data["json"] = ReturnError(4001, "请先登录")
		c.ServeJSON()
	}
	if playUrl == "" {
		c.Data["json"] = ReturnError(4002, "视频地址不能为空")
		c.ServeJSON()
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid, aliyunVideoId)
	if err != nil {
		c.Data["json"] = ReturnError(5000, err)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		c.ServeJSON()
	}
}
