package model

type Comment struct {
	ID       uint `gorm:"primarykey"`
	Uid      uint
	Vid      uint
	Content  string
	CreateAt string
}
