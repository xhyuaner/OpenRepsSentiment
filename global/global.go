package global

import (
	"SDDS/config"
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	DB *gorm.DB

	Trans ut.Translator
	//
	//ServerConfig *config.ServerConfig = &config.ServerConfig{}
	//
	//NacosConfig *config.NacosConfig = &config.NacosConfig{}
	//
	//UserSrvClient proto.UserClient
)
