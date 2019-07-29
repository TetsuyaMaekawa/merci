package model

import "time"

// MerciEmpMaster 社員マスタ
type MerciEmpMaster struct {
	CreatedAt *time.Time `gorm:"column:created_at"`
	CreatedID int64      `gorm:"column:created_id"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	UpdatedID int64      `gorm:"column:updated_id"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (m MerciEmpMaster) tableName() string {
	return "emp_master"
}
