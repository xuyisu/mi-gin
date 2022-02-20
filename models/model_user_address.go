package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//UserAddress
type UserAddress struct {
	Id           uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	AddressId    uint64     `gorm:"column:address_id"  json:"addressId,string" comment:"地址id" sql:"bigint(20) unsigned"`
	CreateTime   *time.Time `gorm:"column:create_time"  json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime   *time.Time `gorm:"column:update_time" json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser   uint64     `gorm:"column:create_user"  json:"createUser" comment:"创建人" sql:"bigint(20)"`
	UpdateUser   uint64     `gorm:"column:update_user"  json:"updateUser" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag   int        `gorm:"column:delete_flag"  json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	DefaultFlag  int        `gorm:"column:default_flag"  json:"defaultFlag" comment:"默认标志" sql:"tinyint(4)"`
	ReceiveName  string     `gorm:"column:receive_name"  json:"receiveName" comment:"收货人" sql:"varchar(60)"`
	ReceivePhone string     `gorm:"column:receive_phone"  json:"receivePhone" comment:"联系号码" sql:"varchar(20)"`
	Province     string     `gorm:"column:province"  json:"province" comment:"省份" sql:"varchar(20)"`
	ProvinceCode string     `gorm:"column:province_code"  json:"provinceCode" comment:"省份编码" sql:"varchar(10)"`
	City         string     `gorm:"column:city"  json:"city" comment:"城市" sql:"varchar(20)"`
	CityCode     string     `gorm:"column:city_code"  json:"cityCode" comment:"城市编码" sql:"varchar(10)"`
	Area         string     `gorm:"column:area"  json:"area" comment:"区" sql:"varchar(20)"`
	AreaCode     string     `gorm:"column:area_code"  json:"areaCode" comment:"区编码" sql:"varchar(10)"`
	Street       string     `gorm:"column:street"  json:"street" comment:"详细地址" sql:"varchar(100)"`
	PostalCode   string     `gorm:"column:postal_code"  json:"postalCode" comment:"邮编" sql:"varchar(10)"`
	AddressLabel int        `gorm:"column:address_label"  json:"addressLabel" comment:"地址标签" sql:"tinyint(4)"`
}

//TableName
func (m *UserAddress) TableName() string {
	return "user_address"
}

//One
func (m *UserAddress) One() (one *UserAddress, err error) {
	one = &UserAddress{}
	err = crudOne(m, one)
	return
}

//All
func (m *UserAddress) All(q *PaginationQuery) (list *[]UserAddress, total uint, err error) {
	list = &[]UserAddress{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *UserAddress) Update() (err error) {
	where := UserAddress{Id: m.Id}
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	return crudUpdate(m, where)
}

//Create
func (m *UserAddress) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *UserAddress) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
