package controllers

import (
	"encoding/json"
	"testing"
)

type UserTest struct {
	User   string `json:"user"`
	Mobile string `json:"mobile"`
}

type PageInfoTest struct {
	TotalPages   int `json:"total_pages`
	CurrentPages int `josn:"current_pages"`
}

var (
	user UserTest
	page PageInfoTest
)

func setup() {
	user = UserTest{
		User:   "user1",
		Mobile: "111111111111",
	}

	page = PageInfoTest{TotalPages: 1, CurrentPages: 1}
}

func tearDown() {
	user = UserTest{}
	page = PageInfoTest{}
}

func TestNewResponse(t *testing.T) {
	setup()
	defer tearDown()
	ac := ApplicationController{}
	resp := ac.NewResponse()
	if resp == nil {
		t.Log("new api Response struct failed")
		return
	}

	resp.ErrorCode = 1
	resp.ErrorMessage = "test"
	resp.ErrorDescription = "test new response struct"

	respData := make([]UserTest, 1)
	respData[0] = user

	resp.Data = respData
	resp.PageInfo = page

	bytes, err := json.Marshal(resp)
	if err != nil {
		t.Log("response marshal failed")
		return
	}

	t.Logf("response :%s", string(bytes))
}
