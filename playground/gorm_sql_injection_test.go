package playground

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

const (
	dsn = "test:1@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
)

type fooTable struct {
	A int `gorm:"column:a"`
}

func (v fooTable) TableName() string {
	return "foo"
}

func TestWhereWithoutArgs(t *testing.T) {
	dangerousSql := "a=1 order by a;drop table bar;"
	db, _ := open(dsn)
	foo := fooTable{}
	db.Debug().Where(dangerousSql).First(&foo)
}

func TestWhereWithMap(t *testing.T) {
	paramMap := map[string]interface{}{"a": "'1';drop table bar;"}
	db, _ := open(dsn)
	foo := fooTable{}
	db.Debug().Where(paramMap).First(&foo)
}

func open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("DB is nil")
	}
	return db, nil
}
