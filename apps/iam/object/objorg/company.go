package objorg

type CompanyBaseInfo struct {
	// Address 公司地址
	Address string `json:"address" form:"address"`
	// CompanyCode 公司编码
	CompanyCode string `json:"companyCode" form:"companyCode"`
	// CompanyName 公司名称
	CompanyName string `json:"companyName" form:"companyName"`
	// ContactEmail 联系邮箱
	ContactEmail string `json:"contactEmail" form:"contactEmail"`
	// ContactPhone 联系电话
	ContactPhone string `json:"contactPhone" form:"contactPhone"`
	// LegalPerson 法人代表
	LegalPerson string `json:"legalPerson" form:"legalPerson"`
	// Logo 公司Logo
	Logo string `json:"logo" form:"logo"`
	// ShortName 公司简称
	ShortName string `json:"shortName" form:"shortName"`
	// Status 状态: active-正常 trial-试用 expired-已过期 inactive-停用
	Status string `json:"status" form:"status"`
	// TenantID 所属租户ID
	TenantID uint `json:"tenantID" form:"tenantID"`
	// UnifiedSocialCreditCode 统一社会信用代码(18位)
	UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode" form:"unifiedSocialCreditCode"`
}
