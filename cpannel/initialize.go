package main

import (
	"encoding/json"

	"github.com/spf13/viper"
)

func init() {
	err = InitCpannelConfig()
	if err != nil {
		panic(err.Error())
	}

	err = InitMySQL()
	if err != nil {
		panic(err.Error())
	}
}

func InitCpannelConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&appControl)
	if err != nil {
		println("marshar err ")

		return err
	}

	bts, err := json.Marshal(appControl)
	if err != nil {
		println("marshar err ")
		return err
	}
	println(string(bts))

	return nil
}

func InitMySQL() error {
	DB, err = appControl.DB.New()
	if err != nil {
		return err
	}

	err = DB.DB().Ping()
	if err != nil {
		return err
	}

	DB.DB().SetMaxIdleConns(appControl.DB.MaxIdle)
	DB.DB().SetMaxOpenConns(appControl.DB.MaxOpen)
	DB.LogMode(appControl.DB.Debug)

	return nil
}
