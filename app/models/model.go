package models

import (
	"github.com/spf13/cast"
	"time"
)

type BaseModel struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement;"json:"id,omitempty"`
}
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;"json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;"json:"updated_at,omitempty"`
	//DeletedAt time.Time `gorm:"column:deleted_at;index;"json:"deleted_at,omitempty"`
}

func (b BaseModel) GetStringId() string {
	return cast.ToString(b.ID)
}
