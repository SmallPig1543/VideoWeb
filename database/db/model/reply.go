package model

type Reply struct {
	ID       uint `gorm:"primarykey"`
	Uid      uint
	Cid      uint //评论id
	Content  string
	CreateAt string
}
