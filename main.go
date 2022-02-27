package main

import (
	"github.com/aimerneige/web-paste/config"
	"github.com/aimerneige/web-paste/database"
	"github.com/aimerneige/web-paste/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig("config", "yml", "./config")
	database.InitDatabase(database.MysqlDatabase{
		UserName: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		Database: viper.GetString("mysql.database"),
		CharSet:  viper.GetString("mysql.charset"),
	})

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	route.AllRouteCollection(r)
	port := viper.GetString("common.port")
	panic(r.Run(":" + port))
}
