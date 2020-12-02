package models

import "github.com/astaxie/beego/orm"

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int64
	Comment            int
	UserId             int
	IsRecommend        int
}

type VideoData struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
	Comment       int
}

type Episodes struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	PlayUrl       string
	Comment       int
	AliyunVideoId string
}

func init() {
	orm.RegisterModel(new(Video))
}

func GetChannelHot(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end from video where status=1 and is_hot=1 and channel_id=? order by episodes_update_time desc limit 9", channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelRecommendRegionList(channelId int, regionId int) (int64, []VideoData, error){
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end from video where status=1 and is_recommend=1 and region_id=? and channel_id=? order by episodes_update_time desc limit 9", regionId, channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelRecommendTypeList(channelId int, typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end from video where status=1 and is_recommend=1 and type_id=? and channel_id=? order by episodes_update_time desc limit 9", typeId, channelId).QueryRows(&videos)
	return num, videos, err
}

func GetTypeTop(typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end from video where status=1 and type_id=? order by comment desc limit 10", typeId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end from video where status=1 and channel_id=? order by comment desc limit 10", channelId).QueryRows(&videos)
	return num, videos, err
}