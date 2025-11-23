package iammodel

import (
	"gorm.io/gorm"
)

// CompanyEntity 公司管理表结构体
type CompanyEntity struct {
	gorm.Model
	Address                 string `gorm:"column:address;type:varchar(255);;default '';comment: 公司地址"`
	CompanyCode             string `gorm:"column:company_code;type:varchar(32);not null;default '';comment: 公司编码"`
	CompanyName             string `gorm:"column:company_name;type:varchar(128);not null;default '';comment: 公司名称"`
	ContactEmail            string `gorm:"column:contact_email;type:varchar(64);;default '';comment: 联系邮箱"`
	ContactPhone            string `gorm:"column:contact_phone;type:varchar(16);;default '';comment: 联系电话"`
	CreatedBy               uint   `gorm:"column:created_by;type:bigint;not null;default 0;comment: 创建人ID"`
	DeletedBy               uint   `gorm:"column:deleted_by;type:bigint;not null;default 0;comment: 删除人ID"`
	LegalPerson             string `gorm:"column:legal_person;type:varchar(32);;default '';comment: 法人代表"`
	Logo                    string `gorm:"column:logo;type:varchar(255);;default '';comment: 公司Logo"`
	ShortName               string `gorm:"column:short_name;type:varchar(64);;default '';comment: 公司简称"`
	Status                  string `gorm:"column:status;type:varchar(16);;default active;comment: 状态: active-正常 trial-试用 expired-已过期 inactive-停用"`
	TenantID                uint   `gorm:"column:tenant_id;type:bigint;not null;default '';comment: 所属租户ID"`
	UnifiedSocialCreditCode string `gorm:"column:unified_social_credit_code;type:varchar(18);;default '';comment: 统一社会信用代码(18位)"`
	UpdatedBy               uint   `gorm:"column:updated_by;type:bigint;not null;default 0;comment: 更新人ID"`
}

type CompanyEntityList []CompanyEntity

const TableNameCompany = "iam_company"

func (CompanyEntity) TableName() string {
	return TableNameCompany
}

func (l CompanyEntityList) ToMap() map[uint]CompanyEntity {
	m := make(map[uint]CompanyEntity)
	for _, v := range l {
		m[v.ID] = v
	}
	return m
}
