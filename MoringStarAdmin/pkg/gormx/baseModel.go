package gormx

import (
	"gorm.io/gorm"
	"time"
)

type BaseDBStruct struct {
	ID        string         `json:"id" gorm:"size:20;primaryKey;"`  // Unique ID
	CreatedAt time.Time      `json:"created_at" gorm:"index;"`       // Create time
	UpdatedAt time.Time      `json:"updated_at" gorm:"index;"`       // Update time
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;"`       // Delete time
	CreateBy  string         `json:"create_by" gorm:"size:20;index"` //创建人
	DeleteBy  string         `json:"delete_by" gorm:"size:20;index"` //删除人
}
