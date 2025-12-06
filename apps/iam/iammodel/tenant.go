package iammodel

import (
	"gorm.io/gorm"
)

// TenantEntity 租户管理表结构体
type TenantEntity struct {
	gorm.Model
	CreatedBy   uint   `gorm:"column:created_by;type:bigint;not null;default 0;comment: 创建人ID"`
	DeletedBy   uint   `gorm:"column:deleted_by;type:bigint;not null;default 0;comment: 删除人ID"`
	Domain      string `gorm:"column:domain;type:varchar(255);;default '';comment: 租户域名"`
	Logo        string `gorm:"column:logo;type:varchar(255);;default '';comment: 租户logo"`
	Description string `gorm:"column:description;type:varchar(255);;default '';comment: 租户描述"`
	SortOrder   int32  `gorm:"column:sort_order;type:int;;default 0;comment: 排序"`
	Status      string `gorm:"column:status;type:varchar(16);;default active;comment: 状态: active-正常 inactive-停用"`
	TenantCode  string `gorm:"column:tenant_code;type:varchar(32);not null;default '';comment: 租户编码"`
	TenantName  string `gorm:"column:tenant_name;type:varchar(64);not null;default '';comment: 租户名称"`
	UpdatedBy   uint   `gorm:"column:updated_by;type:bigint;not null;default 0;comment: 更新人ID"`
}

type TenantEntityList []TenantEntity

const TableNameTenant = "iam_tenant"

func (TenantEntity) TableName() string {
	return TableNameTenant
}

func (l TenantEntityList) ToMap() map[uint]TenantEntity {
	m := make(map[uint]TenantEntity)
	for _, v := range l {
		m[v.ID] = v
	}
	return m
}
