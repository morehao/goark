package objorg

type CompanyBaseInfo struct {
	Address                 string `json:"address" form:"address"`                                 // 公司地址
	CompanyCode             string `json:"companyCode" form:"companyCode"`                         // 公司编码
	CompanyName             string `json:"companyName" form:"companyName"`                         // 公司名称
	ContactEmail            string `json:"contactEmail" form:"contactEmail"`                       // 联系邮箱
	ContactPhone            string `json:"contactPhone" form:"contactPhone"`                       // 联系电话
	LegalPerson             string `json:"legalPerson" form:"legalPerson"`                         // 法人代表
	Logo                    string `json:"logo" form:"logo"`                                       // 公司Logo
	ShortName               string `json:"shortName" form:"shortName"`                             // 公司简称
	Status                  string `json:"status" form:"status"`                                   // 状态: active-正常 trial-试用 expired-已过期 inactive-停用
	TenantID                uint   `json:"tenantId" form:"tenantId"`                               // 所属租户ID
	UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode" form:"unifiedSocialCreditCode"` // 统一社会信用代码(18位)
}
