package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

var GLOBAL_DB *gorm.DB

type Info struct {
	Name string
	Age  int
}

func (i Info) Value() (driver.Value, error) {
	str, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}

func (i Info) Scan(value any) error {
	str, ok := value.([]byte)
	if !ok {
		log.Fatal("wrong")
	}
	json.Unmarshal(str, i)
	return nil
}

func Customize() {
	GLOBAL_DB.AutoMigrate(&Info{})
	GLOBAL_DB.Create(&Info{Name: "i0Ek3", Age: 18})
}
