package model
import "time"

type Item struct {
	Id int `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Status string `json:"status" gorm:"column:status;"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func(Item) TableName() string {return "boxes"}