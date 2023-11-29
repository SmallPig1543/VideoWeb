package types

type Filter struct {
	Uid       uint   `json:"uid" form:"uid"`
	Vid       uint   `json:"vid" form:"vid"`
	Title     string `json:"title" form:"title"`
	Types     string `json:"types" form:"types"` //类别
	Views     int    `json:"views" form:"views"` //大于该播放量
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}
