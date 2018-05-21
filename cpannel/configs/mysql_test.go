package configs

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Pallinder/go-randomdata"
)

var (
	objs []*MySQLConfig
)

func setUpStringData() {
	objs = make([]*MySQLConfig, 10)
	for i := 0; i < 10; i++ {
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

func tearDownStringData() {
	objs = []*MySQLConfig{}

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
