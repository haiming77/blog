package controllers

import (
	"personal.com/blog/cpannel/configs"

	"github.com/jinzhu/gorm"
)

type ApplicationController struct {
	Config *configs.ApplicationConfig
	DB     *gorm.DB
}

type ApiResponse struct {
	ErrorCode        int         `json:"error_code"`
	ErrorMessage     string      `json:"err_msg"`
	ErrorDescription string      `json:"error_description`
	Data             interface{} `json:"response"`
	PageInfo         interface{} `json:"page_info"`
}

func (a ApplicationController) NewResponse() *ApiResponse {
	return &ApiResponse{}
}
