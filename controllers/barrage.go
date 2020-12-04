package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"go-youku/models"
	"net/http"
)

type BarrageController struct {
	beego.Controller
}

type WsData struct {
	CurrentTime int
	EpisodesId  int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// @router /barrage/ws [*]
func (c *BaseController) BarrageWs() {
	var (
		conn     *websocket.Conn
		err      error
		data     []byte
		barrages []models.BarrageData
	)
	if conn, err = upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil); err != nil {
		goto ERR
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		var wsData WsData
		_ = json.Unmarshal([]byte(data), &wsData)
		endTime := wsData.CurrentTime + 60
		_, barrages, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err == nil {
			if err := conn.WriteJSON(barrages); err != nil {
				goto ERR
			}
		}
	}
ERR:
	conn.Close()
}

// @router /barrage/save [*]
func (c *BaseController) Save() {
	uid, _ := c.GetInt("uid")
	content := c.GetString("content")
	currentTime, _ := c.GetInt("currentTime")
	episodesId, _ := c.GetInt("episodesId")
	videoId, _ := c.GetInt("videoId")
	if content == "" {
		c.Data["json"] = ReturnError(4001, "弹幕不能为空")
		c.ServeJSON()
	}
	if uid == 0 {
		c.Data["json"] = ReturnError(4002, "请先登录")
		c.ServeJSON()
	}
	if episodesId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定剧集")
		c.ServeJSON()
	}
	if videoId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频ID")
		c.ServeJSON()
	}
	if currentTime == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频播放时间")
		c.ServeJSON()
	}
	err := models.SaveBarrage(episodesId, videoId, currentTime, uid, content)
	if err != nil {
		c.Data["json"] = ReturnError(5000, err)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", "", 1)
		c.ServeJSON()
	}
}
