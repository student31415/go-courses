package main

import (
	"fmt"
	"github.com/CodingSquire/go-courses/lection_4/calendar/cal"
	"testing"
	"time"
)

//TDD
//func TestSimpleTest(t *testing.T){
//setup
//cases := []struct {
//	month   string
//	quarter int
//}{
//	{month: "01", quarter: 1},
//	{month: "02", quarter: 1},
//	{month: "03", quarter: 1},
//	{month: "04", quarter: 2},
//	{month: "05", quarter: 2},
//	{month: "06", quarter: 2},
//	{month: "07", quarter: 3},
//	{month: "08", quarter: 3},
//	{month: "09", quarter: 3},
//	{month: "10", quarter: 4},
//	{month: "11", quarter: 4},
//	{month: "12", quarter: 4},
//}
//RUN
//t.Run("SimpleTest",func(){
//

//equal
//
//})
//
//}

func TestCurrentQuarter(t *testing.T) {
	cases := []struct {
		month   string
		quarter int
	}{
		{month: "01", quarter: 1}, // 1/3=0 +1=1
		{month: "02", quarter: 1}, // 2/3=0 +1=1
		{month: "03", quarter: 1}, // 3/3=1 +1=2
		{month: "04", quarter: 2}, // 4/3=1 +1=2
		{month: "05", quarter: 2}, // 5/3=1 +1=2
		{month: "06", quarter: 2}, // 6/3=2 +1=3
		{month: "07", quarter: 3},
		{month: "08", quarter: 3},
		{month: "09", quarter: 3},
		{month: "10", quarter: 4},
		{month: "11", quarter: 4},
		{month: "12", quarter: 4},
	}

	//TODO Реализовать Календарь

	for _, test := range cases {
		parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", test.month))
		calendar := cal.NewCalendar(parsed)
		//calendar.CurrData=time.Now()
		actual := calendar.CurrentQuarter()
		calendar2 := cal.NewCalendarWithName(parsed, "cal name")
		calendar2.CurrentQuarter()

		groupCal := make([]cal.CalenarDO, 2)
		groupCal[1] = calendar
		groupCal[2] = calendar2

		//calendar.Value=1

		if actual != test.quarter {
			t.Error("Month:", test.month,
				"Expected Quarter:", test.quarter,
				"Actual Quarter:", actual)
		}
	}
}
