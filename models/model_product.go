package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Product
type Product struct {
	Id         uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime *time.Time `gorm:"column:update_time"  json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser uint64     `gorm:"column:create_user"  json:"createUser" comment:"创建人" sql:"bigint(20)"`
	UpdateUser uint64     `gorm:"column:update_user"  json:"updateUser" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag int        `gorm:"column:delete_flag"  json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	ProductId  uint64     `gorm:"column:product_id"  json:"productId,string" comment:"商品id" sql:"bigint(20)"`
	CategoryId uint64     `gorm:"column:category_id"  json:"categoryId,string" comment:"品类id" sql:"bigint(20)"`
	Name       string     `gorm:"column:name"  json:"name" comment:"商品名称" sql:"varchar(60)"`
	SubTitle   string     `gorm:"column:sub_title"  json:"subTitle" comment:"简要描述" sql:"varchar(100)"`
	MainImage  string     `gorm:"column:main_image"  json:"mainImage" comment:"商品图片地址" sql:"varchar(100)"`
	SubImages  string     `gorm:"column:sub_images" json:"subImages" comment:"子图片列表" sql:"varchar(100)"`
	ActivityId uint64     `gorm:"column:activity_id"  json:"activityId,string" comment:"活动id" sql:"bigint(20)"`
	Status     int        `gorm:"column:status" json:"status" comment:"商品状态" sql:"tinyint(4)"`
	Price      float64    `gorm:"column:price" json:"price" comment:"商品单价" sql:"decimal(20,2)"`
	Stock      int        `gorm:"column:stock"  json:"stock" comment:"库存数" sql:"int(11)"`
}

//TableName
func (m *Product) TableName() string {
	return "product"
}

//One
func (m *Product) One() (one *Product, err error) {
	one = &Product{}
	err = crudOne(m, one)
	return
}

//All
func (m *Product) All(q *PaginationQuery) (list *[]Product, total uint, err error) {
	list = &[]Product{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *Product) Update() (err error) {
	where := Product{Id: m.Id}
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	return crudUpdate(m, where)
}

//Create
func (m *Product) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *Product) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
