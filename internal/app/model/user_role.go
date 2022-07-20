package model

import (
	"gorm.io/gorm"
)

// Role 角色表
type UserRole struct {
	BaseModel
	UserId uint64 `gorm:"column:user_id;type:int(10);not null;comment:用户id"`
	RoleId uint64 `gorm:"column:user_id;type:int(10);not null;comment:角色id"`
}

func (u UserRole) TableName() string {
	return "role"
}

func (u UserRole) Migrate(db *gorm.DB) error {
	if err := db.Set(
		"gorm:table_options",
		"ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表'",
	).AutoMigrate(u); err != nil {
		return err
	}

	return nil
}
