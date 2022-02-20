package models

import (
	"errors"
	"go-gin/pkg/lib"
	"time"
)

var _ = time.Thursday

//Cart
type Cart struct {
	Id                uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	CreateTime        *time.Time `gorm:"column:create_time" json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime        *time.Time `gorm:"column:update_time"  json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser        uint64     `gorm:"column:create_user"  json:"createUser,string" comment:"创建人" sql:"bigint(20)"`
	UpdateUser        uint64     `gorm:"column:update_user"  json:"updateUser,string" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag        int        `gorm:"column:delete_flag" json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	UserId            uint64     `gorm:"column:user_id"  json:"userId,string" comment:"用户id" sql:"bigint(11)"`
	ActivityId        uint64     `gorm:"column:activity_id"  json:"activityId,string" comment:"活动id" sql:"bigint(11)"`
	ActivityName      string     `gorm:"column:activity_name"  json:"activityName" comment:"活动名称" sql:"varchar(255)"`
	ProductId         uint64     `gorm:"column:product_id"  json:"productId,string" comment:"商品id" sql:"bigint(11)"`
	ProductName       string     `gorm:"column:product_name"  json:"productName" comment:"商品名称" sql:"varchar(255)"`
	ProductSubtitle   string     `gorm:"column:product_subtitle"  json:"productSubtitle" comment:"商品简要描述" sql:"varchar(255)"`
	ProductMainImage  string     `gorm:"column:product_main_image"  json:"productMainImage" comment:"商品图片地址" sql:"varchar(255)"`
	Quantity          int        `gorm:"column:quantity"  json:"quantity" comment:"数量" sql:"int(10) unsigned"`
	ProductUnitPrice  float64    `gorm:"column:product_unit_price"  json:"productUnitPrice" comment:"单价" sql:"decimal(20,2) unsigned"`
	Selected          int8       `gorm:"column:selected"  json:"selected" comment:"是否已选择 1是 0 否" sql:"tinyint(4)"`
	ProductTotalPrice float64    `gorm:"column:product_total_price"  json:"productTotalPrice" comment:"总价格" sql:"decimal(20,2)"`
}

type CartResp struct {
	//购物车总价
	CartTotalPrice string `json:"cartTotalPrice"`
	//总数量
	CartTotalQuantity int `json:"cartTotalQuantity"`
	//是否全选
	SelectedAll bool `json:"selectedAll"`

	//购物车列表
	CartProductList interface{} `json:"cartProductList"`
}

type CartReq struct {
	//数量
	//Quantity int `json:"quantity"`
	//是否选中
	Selected interface{} `json:"selected"`
	//类型
	Type interface{} `json:"type"`
}

//TableName
func (m *Cart) TableName() string {
	return "cart"
}

//One
func (m *Cart) One() (one *Cart, err error) {
	one = &Cart{}
	err = crudOne(m, one)
	return
}

//All
func (m *Cart) All(q *PaginationQuery) (list *[]Cart, total uint, err error) {
	list = &[]Cart{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *Cart) Update() (err error) {
	where := Cart{Id: m.Id}
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	return crudUpdate(m, where)
}

//Create
func (m *Cart) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *Cart) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

//查询当前用户购物车数量
func (m *Cart) GetCartCount(userId uint64) (c uint) {
	var count uint
	mysqlDB.Model(m).Where("delete_flag =? and user_id=?", lib.Zero, userId).Count(&count)
	return count
}
