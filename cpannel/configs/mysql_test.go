package configs

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Pallinder/go-randomdata"
)

var (
	objs   []*MySQLConfig
	sqlObj *MySQLConfig
)

func setUpStringData() {
	objs = make([]*MySQLConfig, 10)
	objs[0] = &MySQLConfig{
		UserName:   "root",
		Password:   "root",
		Database:   "blog",
		Host:       "tcp(localhost:3306)",
		Parameters: "charset=utf8mb4&parseTime=True",
		MaxIdle:    10,
		MaxOpen:    10,
		Debug:      true,
	}
	for i := 1; i < 10; i++ {
		objs[i] = &MySQLConfig{
			UserName:   randomdata.SillyName(),
			Password:   randomdata.SillyName(),
			Database:   randomdata.RandStringRunes(3),
			Host:       fmt.Sprintf("tcp(%s:3306)", randomdata.IpV4Address()),
			Parameters: "charset=utf8mb4&parseTime=True",
			MaxIdle:    randomdata.Number(1, 5),
			MaxOpen:    randomdata.Number(1, 5),
			Debug:      true,
		}
	}

	jsonStr, _ := json.Marshal(objs)
	println(string(jsonStr))

}

func setUpNewSqlData() {
	sqlObj = &MySQLConfig{
		UserName:   "root",
		Password:   "root",
		Database:   "blog",
		Host:       "tcp(localhost:3306)",
		Parameters: "charset=utf8mb4&parseTime=True",
		MaxIdle:    10,
		MaxOpen:    10,
		Debug:      true,
	}

	jsonStr, _ := json.Marshal(sqlObj)
	println(string(jsonStr))

}

func tearDownStringData() {
	objs = []*MySQLConfig{}

}

func tearDownNewSqlData() {
	sqlObj = &MySQLConfig{}

}

func TestString(t *testing.T) {

	setUpStringData()
	defer tearDownStringData()

	jsonStr, _ := json.Marshal(objs)
	println(string(jsonStr))
	for _, obj := range objs {
		str := obj.String()

		if str == fmt.Sprintf("%s:%s@%s/%s?%s", obj.UserName, obj.Password, obj.Host, obj.Database, obj.Parameters) {
			t.Logf("the method handle success and return str is :%s", str)
		} else {
			t.Logf("the method handle failed \n return str is :%s \n right string is:root:root@/blog?charset=utf8mb4&parseTime=True ", str)
		}
	}
}

func TestNew(t *testing.T) {
	setUpNewSqlData()
	defer tearDownNewSqlData()

	_, err := sqlObj.New()
	if err != nil {
		t.Logf("it should be init success,it error is %s", err.Error())
	} else {
		t.Log("db init success")
	}
}
