package playground

import (
	"fmt"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	parsedTime, err := time.Parse(time.DateTime, "2023-06-08 10:10:10")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("parsedTime: ", parsedTime)
}

func TestParseTimeLayout(t *testing.T) {
	parsedTime, err := time.Parse("2006-01-02 00:00:00", "2023-09-19 00:00:00")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("parsedTime: ", parsedTime)
}

func TestParseTimeInLocation(t *testing.T) {
	localLocation, _ := time.LoadLocation("Local")
	parsedTime, err := time.ParseInLocation(time.DateTime, "2023-06-08 10:10:10", localLocation)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("parsedTime: ", parsedTime)
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("20060102"))
}

func TestTimeFormatLayout(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 01:01:01"))
}

func TestAddMinusOneMonthToJanuary(t *testing.T) {
	january, _ := time.Parse(time.DateTime, "2023-01-01 01:01:01")
	fmt.Println(lastMonth(january).Format(time.DateTime))
}

func lastMonth(t time.Time) time.Time {
	return t.AddDate(0, -1, 0)
}

func TestLastMonth(t *testing.T) {
	var m time.Time
	m, _ = time.Parse(time.DateTime, "2023-05-01 00:00:00")
	fmt.Printf("lastMonth(%v): %v\n", m, lastMonth(m))
	m, _ = time.Parse(time.DateTime, "2023-05-31 00:00:00")
	fmt.Printf("lastMonth(%v): %v\n", m, lastMonth(m))
	m, _ = time.Parse(time.DateTime, "2023-03-01 00:00:00")
	fmt.Printf("lastMonth(%v): %v\n", m, lastMonth(m))
}

func lastDayOfMonth(t time.Time) time.Time {
	firstDayOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	lastDayOfThisMonth := firstDayOfNextMonth.AddDate(0, 0, -1)
	return lastDayOfThisMonth
}

func TestLastDayOfMonth(t *testing.T) {
	var m time.Time
	m, _ = time.Parse(time.DateTime, "2023-05-01 00:00:00")
	fmt.Printf("lastDayOfMonth(%v): %v\n", m, lastDayOfMonth(m))
	m, _ = time.Parse(time.DateTime, "2023-05-31 00:00:00")
	fmt.Printf("lastDayOfMonth(%v): %v\n", m, lastDayOfMonth(m))
	m, _ = time.Parse(time.DateTime, "2023-03-01 00:00:00")
	fmt.Printf("lastDayOfMonth(%v): %v\n", m, lastDayOfMonth(m))
}

func TestMultiplyDurationByVariable(t *testing.T) {
	month := 31 * 24 * time.Hour
	multiplier := 3
	duration := time.Duration(multiplier) * month
	fmt.Printf("%d * month  = %v\n", multiplier, duration)
}

func TestAddToTimeDoesNotChangeTheOriginalTime(t *testing.T) {
	now := time.Now()
	originalTime := now
	fmt.Printf("original time: %v\n", originalTime)
	now.Add(time.Hour)
	timeAfterCallingAdd := now
	fmt.Printf("time after calling time.Add(): %v\n", timeAfterCallingAdd)
	if timeAfterCallingAdd != originalTime {
		t.Errorf("timeAfterCallingAdd %v not equal to originalTime %v", timeAfterCallingAdd, originalTime)
	}
}
