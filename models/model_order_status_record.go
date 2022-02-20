package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//OrderStatusRecord
type OrderStatusRecord struct {
	Id            uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20),PRI"`
	CreateTime    *time.Time `gorm:"column:create_time" json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	OrderNo       string     `gorm:"column:order_no"  json:"orderNo" comment:"订单编号" sql:"varchar(60)"`
	OrderDetailNo string     `gorm:"column:order_detail_no"  json:"orderDetailNo" comment:"订单明细编号" sql:"varchar(60)"`
	ProductId     uint64     `gorm:"column:product_id"  json:"productId,string" comment:"商品id" sql:"int(11)"`
	ProductName   string     `gorm:"column:product_name"  json:"productName" comment:"商品名称" sql:"varchar(60)"`
	Status        int        `gorm:"column:status"  json:"status" comment:"订单状态" sql:"tinyint(4)"`
	StatusDesc    string     `gorm:"column:status_desc"  json:"statusDesc" comment:"状态描述" sql:"varchar(60)"`
}

//TableName
func (m *OrderStatusRecord) TableName() string {
	return "order_status_record"
}

//One
func (m *OrderStatusRecord) One() (one *OrderStatusRecord, err error) {
	one = &OrderStatusRecord{}
	err = crudOne(m, one)
	return
}

//All
func (m *OrderStatusRecord) All(q *PaginationQuery) (list *[]OrderStatusRecord, total uint, err error) {
	list = &[]OrderStatusRecord{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *OrderStatusRecord) Update() (err error) {
	where := OrderStatusRecord{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *OrderStatusRecord) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *OrderStatusRecord) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
