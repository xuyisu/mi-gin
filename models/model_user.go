package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//User
type User struct {
	Id         uint64     `gorm:"column:id" json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	CreateTime *time.Time `gorm:"column:create_time"  json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime *time.Time `gorm:"column:update_time"  json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser uint64     `gorm:"column:create_user" json:"createUser" comment:"创建人" sql:"bigint(20)"`
	UpdateUser uint64     `gorm:"column:update_user"  json:"updateUser" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag int        `gorm:"column:delete_flag"  json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	Status     int        `gorm:"column:status" json:"status" comment:"启用标志" sql:"tinyint(4)"`
	UserName   string     `gorm:"column:user_name"  json:"userName" comment:"用户名" sql:"varchar(50)"`
	Email      string     `gorm:"column:email"  json:"email" comment:"邮箱" sql:"varchar(50)"`
	Phone      string     `gorm:"column:phone"  json:"phone" comment:"手机号" sql:"varchar(20),UNI"`
	Password   string     `gorm:"column:password"  json:"password" comment:"密码" sql:"varchar(100)"`
}

type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginUser struct {
	Id       uint64 `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

//TableName
func (m *User) TableName() string {
	return "user"
}

//One
func (m *User) One() (one *User, err error) {
	one = &User{}
	err = crudOne(m, one)
	return
}

//All
func (m *User) All(q *PaginationQuery) (list *[]User, total uint, err error) {
	list = &[]User{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *User) Update() (err error) {
	where := User{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *User) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *User) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
