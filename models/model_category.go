package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Category
type Category struct {
	Id         uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime *time.Time `gorm:"column:update_time"  json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser uint64     `gorm:"column:create_user"  json:"createUser" comment:"创建人" sql:"bigint(20)"`
	UpdateUser uint64     `gorm:"column:update_user"  json:"updateUser" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag int        `gorm:"column:delete_flag"  json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	ParentId   uint64     `gorm:"column:parent_id"  json:"parentId,string" comment:"父id" sql:"bigint(20)"`
	Name       string     `gorm:"column:name"  json:"name" comment:"名称" sql:"varchar(100)"`
	Status     int        `gorm:"column:status"  json:"status" comment:"启用禁用状态 1启用 0禁用" sql:"tinyint(4)"`
	SortOrder  int        `gorm:"column:sort_order"  json:"sortOrder" comment:"排序" sql:"int(11)"`
}

//TableName
func (m *Category) TableName() string {
	return "category"
}

//One
func (m *Category) One() (one *Category, err error) {
	one = &Category{}
	err = crudOne(m, one)
	return
}

//All
func (m *Category) All(q *PaginationQuery) (list *[]Category, total uint, err error) {
	list = &[]Category{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *Category) Update() (err error) {
	where := Category{Id: m.Id}
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	return crudUpdate(m, where)
}

//Create
func (m *Category) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *Category) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
