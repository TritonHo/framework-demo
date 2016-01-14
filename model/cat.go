package model

import "time"

type Cat struct {
	Id     string `xorm:"pk" json:"id" validate:"fixed"`
	UserId string `json:"UserId" validate:"fixed"`

	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required,enum=MALE/FEMALE"`

	CreateTime time.Time `xorm:"created" json:"createTime" validate:"zerotime"`
	UpdateTime time.Time `xorm:"updated" json:"updateTime" validate:"zerotime"`
}

func (c Cat) TableName() string {
	return "cats"
}
