package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// QueryTask 查询任务
type QueryTask struct {
	ID        uint           `gorm:"primarykey;column:id" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"-"`

	TaskName      string     `gorm:"size:100;not null;column:task_name;comment:任务名称" json:"task_name"`
	Databases     string     `gorm:"type:text;not null;column:databases;comment:目标数据库列表(JSON格式，包含instance_id和database_name)" json:"databases"`
	Status        int8       `gorm:"not null;default:0;column:status;comment:任务状态：0-待执行，1-执行中，2-已完成，3-失败" json:"status"`
	TotalDBs      int        `gorm:"not null;default:0;column:total_dbs;comment:数据库总数" json:"total_dbs"`
	CompletedDBs  int        `gorm:"not null;default:0;column:completed_dbs;comment:已完成数据库数" json:"completed_dbs"`
	FailedDBs     int        `gorm:"not null;default:0;column:failed_dbs;comment:失败数据库数" json:"failed_dbs"`
	TotalSQLs     int        `gorm:"not null;default:0;column:total_sqls;comment:SQL语句总数" json:"total_sqls"`
	CompletedSQLs int        `gorm:"not null;default:0;column:completed_sqls;comment:已完成SQL数" json:"completed_sqls"`
	FailedSQLs    int        `gorm:"not null;default:0;column:failed_sqls;comment:失败SQL数" json:"failed_sqls"`
	StartedAt     *time.Time `gorm:"column:started_at;comment:开始执行时间" json:"started_at"`
	CompletedAt   *time.Time `gorm:"column:completed_at;comment:完成时间" json:"completed_at"`
	Description   string     `gorm:"type:text;column:description;comment:任务描述" json:"description"`
	IsFavorite    bool       `gorm:"default:false;column:is_favorite;comment:是否为常用任务" json:"is_favorite"`

	// 关联
	SQLs []QueryTaskSQL `gorm:"foreignKey:TaskID" json:"sqls,omitempty"`
}

// TableName 指定表名
func (QueryTask) TableName() string {
	return "query_task_tasks"
}

// TaskDatabase 任务目标数据库结构
type TaskDatabase struct {
	InstanceID   uint   `json:"instance_id"`   // 实例ID
	DatabaseName string `json:"database_name"` // 数据库名称
	InstanceName string `json:"instance_name"` // 实例名称
}

// TaskDatabases 任务目标数据库列表
type TaskDatabases []TaskDatabase

// Value 实现 driver.Valuer 接口
func (d TaskDatabases) Value() (driver.Value, error) {
	return json.Marshal(d)
}

// Scan 实现 sql.Scanner 接口
func (d *TaskDatabases) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, d)
}
