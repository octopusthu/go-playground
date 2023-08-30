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
