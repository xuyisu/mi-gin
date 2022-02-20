package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//OrderInfo
type OrderInfo struct {
	Id              uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	CreateTime      *time.Time `gorm:"column:create_time"  json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime      *time.Time `gorm:"column:update_time"  json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser      uint64     `gorm:"column:create_user"  json:"createUser" comment:"创建人" sql:"bigint(20)"`
	UpdateUser      uint64     `gorm:"column:update_user"  json:"updateUser" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag      int        `gorm:"column:delete_flag" json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	OrderNo         string     `gorm:"column:order_no"  json:"orderNo" comment:"订单编号" sql:"varchar(60)"`
	Payment         float64    `gorm:"column:payment"  json:"payment" comment:"支付金额" sql:"decimal(20,2)"`
	PaymentType     int        `gorm:"column:payment_type"  json:"paymentType" comment:"支付类型" sql:"tinyint(4)"`
	PaymentTypeDesc string     `gorm:"column:payment_type_desc"  json:"paymentTypeDesc" comment:"支付类型描述" sql:"varchar(20)"`
	Postage         float32    `gorm:"column:postage" json:"postage" comment:"邮费" sql:"decimal(20,2)"`
	Status          int        `gorm:"column:status"  json:"status" comment:"订单状态" sql:"tinyint(4)"`
	StatusDesc      string     `gorm:"column:status_desc"  json:"statusDesc" comment:"状态描述" sql:"varchar(20)"`
	PaymentTime     *time.Time `gorm:"column:payment_time"  json:"paymentTime,omitempty" comment:"支付时间" sql:"datetime"`
	AddressId       uint64     `gorm:"column:address_id" json:"addressId,string" comment:"地址id" sql:"bigint(20)"`
	ReceiveName     string     `gorm:"column:receive_name" json:"receiveName" comment:"收货人" sql:"varchar(50)"`
	ReceivePhone    string     `gorm:"column:receive_phone"  json:"receivePhone" comment:"联系号码" sql:"varchar(20)"`
	Province        string     `gorm:"column:province"  json:"province" comment:"省份" sql:"varchar(20)"`
	City            string     `gorm:"column:city"  json:"city" comment:"城市" sql:"varchar(20)"`
	Area            string     `gorm:"column:area"  json:"area" comment:"区" sql:"varchar(20)"`
	Street          string     `gorm:"column:street"  json:"street" comment:"详细地址" sql:"varchar(50)"`
	PostalCode      string     `gorm:"column:postal_code"  json:"postalCode" comment:"邮编" sql:"varchar(255)"`
	UserId          uint64     `gorm:"column:user_id"  json:"userId,string" comment:"购买人id" sql:"bigint(20)"`
}

//订单发那会列表新对象
type OrderInfoVo struct {
	OrderInfo
	Details []OrderDetail `json:"details"`
}

//支付请求
type PayReq struct {
	OrderNo string `json:"orderNo"`
	PayTool int    `json:"payTool"`
}

//TableName
func (m *OrderInfo) TableName() string {
	return "order_info"
}

//One
func (m *OrderInfo) One() (one *OrderInfo, err error) {
	one = &OrderInfo{}
	err = crudOne(m, one)
	return
}

//All
func (m *OrderInfo) All(q *PaginationQuery) (list *[]OrderInfo, total uint, err error) {
	list = &[]OrderInfo{}

	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *OrderInfo) Update() (err error) {
	now := time.Now()
	where := OrderInfo{Id: m.Id}
	m.Id = 0
	m.UpdateTime = &now
	return crudUpdate(m, where)
}

//Create
func (m *OrderInfo) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *OrderInfo) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
