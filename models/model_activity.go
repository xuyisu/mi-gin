package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Activity
type Activity struct {
	Id         uint64     `gorm:"column:id"  json:"id,string" comment:"主键" sql:"bigint(20) unsigned,PRI"`
	CreateTime *time.Time `gorm:"column:create_time"  json:"createTime,omitempty" comment:"创建时间" sql:"datetime"`
	UpdateTime *time.Time `gorm:"column:update_time"  json:"updateTime,omitempty" comment:"更新时间" sql:"datetime"`
	CreateUser uint64     `gorm:"column:create_user"  json:"createUser" comment:"创建人" sql:"bigint(20)"`
	UpdateUser uint64     `gorm:"column:update_user" json:"updateUser" comment:"更新人" sql:"bigint(20)"`
	DeleteFlag int        `gorm:"column:delete_flag"  json:"deleteFlag" comment:"删除标志" sql:"tinyint(4)"`
	ActivityId uint64     `gorm:"column:activity_id"  json:"activityId,string" comment:"活动id" sql:"bigint(20)"`
	Name       string     `gorm:"column:name"  json:"name" comment:"活动名称" sql:"varchar(60)"`
	Status     int        `gorm:"column:status" json:"status" comment:"活动状态" sql:"tinyint(4)"`
	MainImage  string     `gorm:"column:main_image"  json:"mainImage" comment:"活动图片地址" sql:"varchar(100)"`
	StartTime  *time.Time `gorm:"column:start_time"  json:"startTime,omitempty" comment:"开始时间" sql:"datetime"`
	EndTime    *time.Time `gorm:"column:end_time"  json:"endTime,omitempty" comment:"结束时间" sql:"datetime"`
}

//TableName
func (m *Activity) TableName() string {
	return "activity"
}

//One
func (m *Activity) One() (one *Activity, err error) {
	one = &Activity{}
	err = crudOne(m, one)
	return
}

//All
func (m *Activity) All(q *PaginationQuery) (list *[]Activity, total uint, err error) {
	list = &[]Activity{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *Activity) Update() (err error) {
	where := Activity{Id: m.Id}
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	return crudUpdate(m, where)
}

//Create
func (m *Activity) Create() (err error) {
	m.Id = 0
	now := time.Now()
	m.UpdateTime = &now
	m.CreateTime = &now
	return mysqlDB.Create(m).Error
}

//Delete
func (m *Activity) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

//GetActivityByTime
func (m *Activity) GetActivityByTime(date time.Time) (one *Activity, err error) {
	one = &Activity{}
	if mysqlDB.Where("start_time<? and end_time>? and delete_flag =? and status=?", date, date, 0, 1).First(one).RecordNotFound() {
		err = errors.New("resource is not found")
	}
	return
}
