/**
 * FileName:		config.go
 * Description:		read config file
 * Author:			Qianno.Xie
 * Email:			qianno.xie@appcoachs.com
**/
package config

import (
	"fmt"
)

import(
	"sync"
	"github.com/spf13/viper"
)

var (
	once sync.Once
	conf *viper.Viper
)

func init(){
	once.Do(initilaizing)
}

func initilaizing(){
	fmt.Println("start config")
	conf = viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("toml")
	conf.AddConfigPath("config")
	
	err := conf.ReadInConfig()
	if err != nil{
		return
	}
}

func GetConfig() *viper.Viper{
	return conf
}
