package types

import (
	"gorm.io/gorm"
)

type Scopes func(db *gorm.DB) *gorm.DB

type PageOptions struct {
	Page     uint32
	PageSize uint32
	Scopes   Scopes
}

type CreateModel struct {
	ID        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64  `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
}

type BaseModel struct {
	ID        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64  `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updated_at,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteModel struct {
	ID        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64     `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updated_at,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"-" gorm:"index;comment:删除时间"`
}
