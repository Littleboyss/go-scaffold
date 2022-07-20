package model

import (
	"gorm.io/gorm"
)

// Role 角色表
type Role struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(64);not null;default:'';comment:角色"`
	Remark string `gorm:"column:remark;type:text;comment:描述"`
}

func (u Role) TableName() string {
	return "role"
}

func (u Role) Migrate(db *gorm.DB) error {
	if err := db.Set(
		"gorm:table_options",
		"ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表'",
	).AutoMigrate(u); err != nil {
		return err
	}

	return nil
}
