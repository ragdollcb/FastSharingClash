package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

//配置文件写法

func init() {

	viper.SetConfigType("ini")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("无法打开配置文件: %s \n", err))
	}
}

func GetConfig(key string) string {
	return viper.GetString(key)
}

func GetUUid() string {
	uuid := viper.GetString("default.uuid")
	fmt.Println(uuid)
	if uuid == "" {
		uuid = CreateUuid()
		viper.Set("default.uuid", uuid)
		_ = viper.WriteConfig()
	}
	return uuid
}
