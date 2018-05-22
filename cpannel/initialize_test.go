package main

import (
	"encoding/json"
	"testing"

	"personal.com/blog/cpannel/configs"
)

func TestInitCpannelConfig(t *testing.T) {
	InitCpannelConfig()
	if appControl == nil {
		t.Log("appcontrol is nil")
	}
	env := appControl.ENV
	if env == "" {
		t.Logf("failed init, env should not be empty")
	} else if env != "debug" {
		t.Log("failed init, env should be debug")
	}

	dbConfig := &configs.MySQLConfig{
		UserName:   "root",
		Password:   "root",
		Database:   "blog",
		Host:       "tcp(localhost:3306)",
		Parameters: "charset=utf8mb4&parseTime=True",
		MaxIdle:    10,
		MaxOpen:    10,
		Debug:      true,
	}

	jsonBytes, err := json.Marshal(dbConfig)
	if err != nil {
		t.Logf("MySQLConfig init error, json marshal failed \n%v", appControl)
	}
	str := string(jsonBytes)

	configDbBytes, err := json.Marshal(appControl.DB)
	if err != nil {
		t.Logf("config init error, json marshal failed \n%v", appControl)
	}
	configStr := string(configDbBytes)

	if str == configStr {
		t.Log("success init config file")
	} else {
		t.Logf("failed init config file, the config is :%s", configStr)
	}
}
