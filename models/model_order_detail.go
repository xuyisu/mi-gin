package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//OrderDetail
type OrderDetail struct {
	Id                uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	CreateTime        *time.Time `gorm:"column:create_time"  json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime        *time.Time `gorm:"column:update_time"  json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser        uint64     `gorm:"column:create_user" json:"createUser" comment:"创建人" sql:"bigint(20)"`
	UpdateUser        uint64     `gorm:"column:update_user"  json:"updateUser" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag        int        `gorm:"column:delete_flag"  json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	OrderNo           string     `gorm:"column:order_no"  json:"orderNo" comment:"订单编号" sql:"varchar(60)"`
	OrderDetailNo     string     `gorm:"column:order_detail_no"  json:"orderDetailNo" comment:"订单明细编号" sql:"varchar(60)"`
	ActivityId        uint64     `gorm:"column:activity_id"  json:"activityId,string" comment:"活动id" sql:"bigint(20)"`
	ActivityName      string     `gorm:"column:activity_name"  json:"activityName" comment:"活动名称" sql:"varchar(50)"`
	ActivityMainImage string     `gorm:"column:activity_main_image"  json:"activityMainImage" comment:"活动图片地址" sql:"varchar(100)"`
	ProductId         uint64     `gorm:"column:product_id"  json:"productId,string" comment:"商品id" sql:"bigint(20)"`
	ProductName       string     `gorm:"column:product_name"  json:"productName" comment:"商品名称" sql:"varchar(50)"`
	ProductMainImage  string     `gorm:"column:product_main_image"  json:"productMainImage" comment:"商品图片地址" sql:"varchar(100)"`
	CurrentUnitPrice  float64    `gorm:"column:current_unit_price"  json:"currentUnitPrice" comment:"单价" sql:"decimal(20,2)"`
	Quantity          int        `gorm:"column:quantity"  json:"quantity" comment:"数量" sql:"int(11)"`
	TotalPrice        float64    `gorm:"column:total_price"  json:"totalPrice" comment:"总价" sql:"decimal(20,2)"`
	UserId            uint64     `gorm:"column:user_id" json:"userId,string" comment:"购买人id" sql:"bigint(20)"`
	Status            int        `gorm:"column:status" json:"status" comment:"订单状态" sql:"tinyint(4)"`
	StatusDesc        string     `gorm:"column:status_desc" json:"statusDesc" comment:"状态描述" sql:"varchar(20)"`
	CancelTime        *time.Time `gorm:"column:cancel_time"  json:"cancelTime,omitempty" comment:"取消时间" sql:"datetime"`
	CancelReason      int        `gorm:"column:cancel_reason" json:"cancelReason" comment:"取消原因" sql:"int(4)"`
	SendTime          *time.Time `gorm:"column:send_time" json:"sendTime,omitempty" comment:"发货时间" sql:"datetime"`
	ReceiveTime       *time.Time `gorm:"column:receive_time"  json:"receiveTime,omitempty" comment:"签收时间" sql:"datetime"`
}

//TableName
func (m *OrderDetail) TableName() string {
	return "order_detail"
}

//One
func (m *OrderDetail) One() (one *OrderDetail, err error) {
	one = &OrderDetail{}
	err = crudOne(m, one)
	return
}

//All
func (m *OrderDetail) All(q *PaginationQuery) (list *[]OrderDetail, total uint, err error) {
	list = &[]OrderDetail{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *OrderDetail) Update() (err error) {
	now := time.Now()
	where := OrderDetail{Id: m.Id}
	m.Id = 0
	m.UpdateTime = &now
	return crudUpdate(m, where)
}

//Create
func (m *OrderDetail) Create() (err error) {
	now := time.Now()
	m.Id = 0
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *OrderDetail) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
