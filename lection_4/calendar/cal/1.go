package cal

import (
	"time"
)

type CalenarDO interface {
	CurrentQuarter() (result int)
}
type CalenarName interface {
	GetName() (name string)
	CurrentQuarter() (result int)
}

type calendar struct {
	CurrData time.Time
}

type calendarWithName struct {
	CurrData time.Time
	Name     string
}

func (c *calendarWithName) GetName() (name string) {
	return c.Name
}

func (c *calendar) CurrentQuarter() (result int) {
	curMonth := int(c.CurrData.Month())
	result = (curMonth-1)/3 + 1
	return
}

func (c *calendarWithName) CurrentQuarter() (result int) {
	curMonth := int(c.CurrData.Month())
	result = (curMonth - 1) / 3
	return
}

func NewCalendar(dataTime time.Time) CalenarDO {
	return &calendar{CurrData: dataTime}
}

func NewCalendarWithName(dataTime time.Time, name string) CalenarName {
	return &calendarWithName{
		CurrData: dataTime,
		Name:     name,
	}
}

//
//type Calendar struct{
//	month   string
//	quarter int
//}
//
//func (c *Calendar)CurrentQuarter() int{
//	return c.quarter
//}
//
//func NewCalendar(time time.Time) *Calendar{
//	v:=time.Month()
//	vInt:=int(v)-1
//	return &Calendar{
//		month:   v.String(),
//		quarter: (vInt/3)+1,
//	}
//}
