package main

type BusinessLicens struct {
	No                    string //编号
	Name                  string //名称
	Type                  string //类别
	Address               string //住所
	LegalPerson           string //法定代表人
	RegisteredCapital     int    //注册资本
	DateOfEstablishment   string //成立日期
	BusinessTerm          string //营业期限
	BusinessScope         string //经营范围
	RegistrationAuthority string //登记机关
	RegistrationDate      string //登记日期
	HashOfBusinessLicens  string // 营业执照扫描件哈希值
}

type BankInformation struct {
	BankAccount string //银行账号
	OpeningBank string //开户银行
}

type Contacts struct {
	IDNumber string //身份证号
	Name     string //姓名
	mobile   string //手机号
	Phone    string //座机号
}

//采购方注册登记
type RegistrationPurchaser struct {
	UnitId              string            //机构编号
	UnitName            string            //机构名称
	UnitClass           string            //机构类别（资质等级）
	Businesstype        string            //机构性质（所有制形式）
	BusinessLicens      BusinessLicens    //营业执照信息
	UnitAddress         string            //机构地址
	ReceivingAddress    string            //收货地址
	NumberOfBeds        int               //床位数
	ManagementCondition string            //经营状况（最近两年营业收入，财务报表及相关资质文件）
	BankInformation     []BankInformation // 银行信息				---可重复记录
	Contacts            []Contacts        //联系人	---可重复记录，包括分管院长和财务负责人
}

//采购订单
type PurchaseOrder struct {
	OrderID            string  //订单编号
	Unitid             string  //医疗机构编号
	ProductID          string  //产品ID
	ProductName        string  //产品名称
	ApprovalSymbol     string  //批准文号
	Model              string  //型号
	Specifications     string  //规格
	ComponentName      string  //组件名称
	EnterpriseNumber   string  //生产企业编号
	EnterpriseName     string  //生产企业名称
	DistributionNo     string  //配送企业编号
	DistributionName   string  //配送企业名称
	TransactionPrice   float32 //成交价
	PurchasePrice      float32 //采购价
	PackageNumber      int     //包装数
	PurchaseQuantity   int     //采购数量
	TotalPurchasePrice float32 //采购总价
	OrderStatus        string  //订单状态
	OrderTime          string  //下单时间
}

type WarehouseReceipt struct {
	SerialNumbe            string  //序号,
	Name                   string  //名称,
	Specifications         string  //规格,
	Unit                   string  //单位
	Number                 int     //数量
	Price                  float32 //单价
	AmountOfMoney          float32 //金额
	Remarks                string  //备注
	HashOfWarehouseReceipt string  //入库单扫描件哈希值
}

type TaxableName struct {
	Name           string  //名称
	Specifications string  //规格
	Model          string  //型号
	Number         int     //数量
	Price          float32 //单价
	AmountOfMoney  float32 //金额
	TaxRate        float32 //税率
	TaxAmount      float32 //税额
}

type SalesUnitInformation struct {
	Name                         string //名称
	TaxpayerIdentificationNumber string //纳税人识别号
	Address                      string //地址
	PhoneNumber                  string //电话
	OpeningBank                  string //开户行
	BankAccount                  string //账号
}

type InvoiceInformation struct {
	Number                       string               //编号
	TicketNumber                 string               //票号
	Name                         string               //名称
	TaxpayerIdentificationNumber string               //纳税人识别号
	Address                      string               //地址
	PhoneNumber                  string               //电话
	OpeningBank                  string               //开户行
	BankAccount                  string               //账号
	TaxableName                  TaxableName          //应税名称
	SalesUnitInformation         SalesUnitInformation //销货单位信息
	InvoiceUnit                  string               //开票单位
	InvoicePerson                string               //开票人
	InvoiceDate                  string               //开票日期
	HashOfInvoice                string               //发票扫描件哈希值
}

//收货入库
type ReceivingGoods struct {
	OrderID            string             //订单编号
	OrderStatus        string             //订单状态
	Sign               string             //签收人
	WarehouseReceipt   WarehouseReceipt   //入库单
	InvoiceInformation InvoiceInformation //发票信息
	SendDate           string             //收货时间
	WarehousingTime    string             //入库时间
}
