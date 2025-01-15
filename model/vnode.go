package model

type Model struct {
	Id   int64 `gorm:"primaryKey"`
	Name string
}
