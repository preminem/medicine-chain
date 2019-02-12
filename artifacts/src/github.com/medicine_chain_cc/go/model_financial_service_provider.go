package main

//金融服务方注册登记
type RegistrationFinancialServiceProvider struct {
	EnterpriseNumber           string                     //企业编号
	BusinessLicenseInformation BusinessLicenseInformation //营业执照信息
	Salesman                   []Contacts                 //业务员	 ---可重复记录
}

//企业征信
//type CorporateCredit struct {
//	BasicEnterpriseInformation RegistrationSupplier //企业基本信息
//	应付款订单信息
//{
//	医疗机构编号医疗机构名称合计金额收货确认时间
//	...
//}
//	招投标信息
//{
//}
//	合同信息
//{
//}
//	收付款历史信息
//{
//	医疗机构编号医疗机构名称合计金额收货确认时间付款时间
//	...
//}
//}
