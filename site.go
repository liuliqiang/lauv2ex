package v2ex

/*
* File Name:	v2ex.go
* Description:  www.v2ex.com API
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-10-19
 */

// https://www.v2ex.com/p/7v9TEc53
// http://www.v2ex.com/t/85402

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidUsername = errors.New("v2ex: invalid user name")
	ErrInvalidNodename = errors.New("v2ex: invalid node name")
)

const v2exAPI = "https://www.v2ex.com/api"

// =======================================
// 社区信息 API
type Stats struct {
	TopicMax  *int `json:"topic_max"`  //当前社区话题数量
	MemberMax *int `json:"member_max"` //当前社区用户数量
}
type Info struct {
	Title       string `json:"title"`       //当前社区站名
	Slogan      string `json:"slogan"`      //当前社区口号
	Description string `json:"description"` //当前社区描述
	Domain      string `json:"domain"`      //社区网址
}
type SiteService interface {
	//获取社区统计信息
	GetStats() (stats *Stats, err error)
	//获取社区介绍
	GetInfo() (info *Info, err error)
}

func getAPIData(url string, v interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(v)
	return
}
