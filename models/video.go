package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	redisClient "go-youku/services/redis"
	"strconv"
	"time"
)

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

func RedisGetTypeTop(typeId int) (int64, []VideoData, error) {
	var (
		videos []VideoData
		num int64
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()
	// 定义redisKey
	redisKey := "video:top:type:typeId:" + strconv.Itoa(typeId)
	// 判断是否存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		num = 0
		res, _ := redis.Values(conn.Do("zrevrange", redisKey, "0", "10", "WITHSCORES"))
		for k, v := range res {
			if k%2 == 0 {
				videoId, err := strconv.Atoi(string(v.([]byte)))
				videoInfo, err := RedisGetVideoInfo(videoId)
				if err == nil {
					var videoDataInfo VideoData
					videoDataInfo.Id = videoInfo.Id
					videoDataInfo.Img = videoInfo.Img
					videoDataInfo.Img1 = videoInfo.Img1
					videoDataInfo.IsEnd = videoInfo.IsEnd
					videoDataInfo.SubTitle = videoInfo.SubTitle
					videoDataInfo.Title = videoInfo.Title
					videoDataInfo.AddTime = videoInfo.AddTime
					videoDataInfo.Comment = videoInfo.Comment
					videoDataInfo.EpisodesCount = videoInfo.EpisodesCount
					videos = append(videos, videoDataInfo)
					num++
				}
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("select id, title, sub_title, img, img1, add_time, episodes_count, is_end from video where status=1 and type_id=? order by comment desc limit 10", typeId).QueryRows(&videos)
		if err == nil {
			for _, v := range videos {
				conn.Do("zadd", redisKey, v.Comment, v.Id)
			}
			conn.Do("expire", redisKey, 86400*30)
		}
	}
	return num, videos, err
}

func RedisGetChannelTop(channelId int) (int64, []VideoData, error) {
	var (
		videos []VideoData
		num int64
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()
	// 定义redisKey
	redisKey := "video:top:channel:channelId:" + strconv.Itoa(channelId)
	// 判断是否存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		num = 0
		res, _ := redis.Values(conn.Do("zrevrange", redisKey, "0", "10", "WITHSCORES"))
		for k, v := range res {
			if k%2 == 0 {
				videoId, err := strconv.Atoi(string(v.([]byte)))
				videoInfo, err := RedisGetVideoInfo(videoId)
				if err == nil {
					var videoDataInfo VideoData
					videoDataInfo.Id = videoInfo.Id
					videoDataInfo.Img = videoInfo.Img
					videoDataInfo.Img1 = videoInfo.Img1
					videoDataInfo.IsEnd = videoInfo.IsEnd
					videoDataInfo.SubTitle = videoInfo.SubTitle
					videoDataInfo.Title = videoInfo.Title
					videoDataInfo.AddTime = videoInfo.AddTime
					videoDataInfo.Comment = videoInfo.Comment
					videoDataInfo.EpisodesCount = videoInfo.EpisodesCount
					videos = append(videos, videoDataInfo)
					num++
				}
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("select id, title, sub_title, img, img1, add_time, episodes_count, is_end from video where status=1 and channel_id=? order by comment desc limit 10", channelId).QueryRows(&videos)
		if err == nil {
			for _, v := range videos {
				conn.Do("zadd", redisKey, v.Comment, v.Id)
			}
			conn.Do("expire", redisKey, 86400*30)
		}
	}
	return num, videos, err
}

func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end from video where status=1 and channel_id=? order by comment desc limit 10", channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params
	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" {
		qs = qs.Filter("is_end", 0)
	}else if end == "y" {
		qs = qs.Filter("is_end",1)
	}
	if sort == "episodesUpdateTime" {
		qs = qs.OrderBy("-episodes_update_time")
	} else if sort == "comment" {
		qs = qs.OrderBy("-comment")
	} else if sort == "addTime" {
		qs = qs.OrderBy("-add_time")
	} else {
		qs = qs.OrderBy("-add_time")
	}
	nums, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit,offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	return nums, videos, err
}

func GetVideoInfo(videoId int) (Video, error) {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("select * from video where id=? limit 1", videoId).QueryRow(&video)
	return video, err
}

func RedisGetVideoInfo(videoId int) (Video, error) {
	var video Video
	conn := redisClient.PoolConnect()
	defer conn.Close()
	redisKey := "video:id:" + strconv.Itoa(videoId)
	// 判断redis中是否存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(res, &video)
	} else {
		o := orm.NewOrm()
		err := o.Raw("select * from video where id=? limit 1", videoId).QueryRow(&video)
		if err == nil {
			_, err := conn.Do("hmset", redis.Args{redisKey}.AddFlat(video)...)
			if err == nil {
				conn.Do("expire", redisKey, 86400)
			}
		}
	}
	return video, err
}

func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.Raw("select id,title, add_time, num, play_url, comment from video_episodes where video_id=? order by num asc", videoId).QueryRows(&episodes)
	return num, episodes, err
}

func RedisGetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	var (
		episodes []Episodes
		num int64
		err error
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()
	redisKey := "video:episodes:videoId:" + strconv.Itoa(videoId)
	// 判断redisKey是否存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		num, err = redis.Int64(conn.Do("llen", redisKey))
		if err == nil {
			values, _ := redis.Values(conn.Do("lrange", redisKey, "0", "-1"))
			var episodesInfo Episodes
			for _, v := range values {
				err = json.Unmarshal(v.([]byte), &episodesInfo)
				if err == nil {
					episodes = append(episodes, episodesInfo)
				}
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("select id, title, add_time, num, play_url, comment, aliyun_video_id from video_episodes where video_id=? order by num asc", videoId).QueryRows(&episodes)
		if err == nil {
			// 遍历获取到的数据，json化存入redis
			for _, v := range episodes {
				jsonValue, err := json.Marshal(v)
				if err == nil {
					conn.Do("rpush", redisKey, jsonValue)
				}
			}
			conn.Do("expire", redisKey, 86400)
		}
	}
	return num, episodes, err
}

func GetUserVideo(uid int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, img, img1, add_time, episodes_count, is_end from video where user_id=? order by add_time desc", uid).QueryRows(&videos)
	return num, videos, err
}

func SaveAliyunVideo(videoId string, log string) error {
	o := orm.NewOrm()
	_, err := o.Raw("insert into aliyun_video (video_id, log, add_time) values (?,?,?)", videoId, log, time.Now().Unix()).Exec()
	fmt.Println(err)
	return err
}

func SaveVideo(title string, subTitle string, channelId int, regionId int, typeId int, playUrl string, uid int, aliyunVideoId string) error {
	o := orm.NewOrm()
	var video Video
	tm := time.Now().Unix()
	video.Title = title
	video.SubTitle = subTitle
	video.AddTime = tm
	video.Img = ""
	video.Img1 = ""
	video.EpisodesCount = 1
	video.IsEnd = 1
	video.ChannelId = channelId
	video.Status = 1
	video.RegionId = regionId
	video.TypeId = typeId
	video.EpisodesUpdateTime = tm
	video.Comment = 0
	video.UserId = uid
	videoId, err := o.Insert(&video)
	if err == nil {
		if aliyunVideoId != "" {
			playUrl = ""
		}
		_, err = o.Raw("insert into video_episodes (title, add_time, num, video_url, status, comment, aliyun_video_id) values (?,?,?,?,?,?,?,?)", subTitle, tm, 1, videoId, playUrl, 1, 0, aliyunVideoId).Exec()
	}
	return err
}