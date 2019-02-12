package main

type AuditInformation struct {
	AuditStat    string // 审核状态
	AuditOpinion string //审核意见
}

type BusinessLicenseInformation struct {
	Namber                string //编号
	Name                  string //名称
	Category              string //类别
	Address               string //住所
	LegalPerson           string //法定代表人
	RegisteredCapital     int    //注册资本
	DateOfEstablishment   string //成立日期
	BusinessTerm          string //营业期限
	BusinessScope         string //经营范围
	RegistrationAuthority string //登记机关
	registrationDate      string //登记日期
	HashOfBusinessLicens  string //营业执照扫描件哈希值
}

type ProductionLicense struct {
	Number                  string //编号
	EnterpriseName          string //企业名称
	LegalPerson             string //法定代表人
	HeadOfEnterprise        string //企业负责人
	TypesOfEnterprises      string //企业类型
	RegisteredAddress       string //注册地址
	ProductionAddress       string //生产地址
	ProductionScope         string //生产范围
	TermOfValidity          string //有效期
	CertificationAuthority  string //发证机关
	DateOfIssue             string //发证日期
	HashOfProductionLicense string //生产许可证扫描件哈希值
}

//供应方注册登记
type RegistrationSupplier struct {
	EnterpriseNumber           string                     //企业编号
	EnterpriseName             string                     //企业名称
	EnterpriseAddress          string                     //企业地址
	EnterpriseCategory         string                     //企业类别（生产/代理）(Production/Agent)
	AuditInformation           AuditInformation           //审核信息
	hospitalAccountNumber      string                     //医院账号
	BusinessLicenseInformation BusinessLicenseInformation //营业执照信息
	ProductionLicense          ProductionLicense          //生产许可证信息
	BankInformation            []BankInformation          // 银行信息   ---可重复记录
	Contacts                   []Contacts                 //业务联系人  ---可重复记录，包括分管院长和财务负责人
	Tenderer                   []Contacts                 //投标人	     ---可重复记录，包括分管院长和财务负责人
}

type OutboundOrder struct {
	SerialNumbe         string  //序号,
	Name                string  //名称,
	Specifications      string  //规格,
	Unit                string  //单位
	Number              int     //数量
	Price               float32 //单价
	AmountOfMoney       float32 //金额
	Remarks             string  //备注
	HashOfOutboundOrder string  //出库单扫描件哈希值
}

//订单发货
type OrderDelivery struct {
	OrderID            string             //订单编号
	OrderStatus        string             //订单状态
	OutboundOrder      OutboundOrder      //出库单
	InvoiceInformation InvoiceInformation //发票信息
	DeliveryTime       string             //发货时间
}

//融资申请
type FinancingApplication struct {
	EnterpriseNumber           string                     //企业编号
	EnterpriseName             string                     //企业名称
	EnterpriseAddress          string                     //企业地址
	EnterpriseCategory         string                     //企业类别（生产/代理）(Production/Agent)
	AuditInformation           AuditInformation           //审核信息
	accountOfHospital          string                     //医院账号
	BusinessLicenseInformation BusinessLicenseInformation //营业执照信息
	ProductionLicense          ProductionLicense          //生产许可证信息
	BankInformation            []BankInformation          // 银行信息				---可重复记录
	FinancingScale             float64                    //拟融资规模
	AmountOfAccountsReceivable float64                    //应收账款金额
}
