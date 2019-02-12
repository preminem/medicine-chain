package main

//物流方注册登记
type RegistrationLogistics struct {
	EnterpriseNumber           string                     //企业编号
	EnterpriseName             string                     //企业名称
	EnterpriseAddress          string                     //企业地址
	EnterpriseCategory         string                     //企业类别（生产/代理）（Production/Agent）
	AuditInformation           AuditInformation           //审核信息
	BusinessLicenseInformation BusinessLicenseInformation //营业执照信息
	DeliveryClerk              []Contacts                 //配送员				---可重复记录，包括分管院长和财务负责人
}

//货物送达
type DeliveryOfGoods struct {
	OrderID       string //订单编号
	OrderStat     string //订单状态
	DeliveryClerk string //送货员
	Signatory     string //签收人
	ServiceTime   string //送达时间
}
