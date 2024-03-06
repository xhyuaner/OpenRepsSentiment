package initialize

import (
	"SDDS/global"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}

func InitConfig() {
	debug := GetEnvInfo("SDDS_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("developer-sentiment/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("developer-sentiment/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)

	////这个对象如何在其他文件中使用 - 全局变量
	//if err := v.Unmarshal(global.NacosConfig); err != nil {
	//	panic(err)
	//}
	//zap.S().Infof("配置信息: %v", global.NacosConfig)
	////从nacos中读取配置信息
	//sc := []constant.ServerConfig{
	//	{
	//		IpAddr: global.NacosConfig.Host,
	//		Port:   global.NacosConfig.Port,
	//	}, //可设置多个nacos配置中心
	//}
	//
	//cc := constant.ClientConfig{
	//	NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "tmp/nacos/log", //如果文件夹创建不成功可以手动提前创建好
	//	CacheDir:            "tmp/nacos/cache",
	//	RotateTime:          "1h",
	//	MaxAge:              3,
	//	LogLevel:            "debug",
	//}
	//
	//configClient, err := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": sc,
	//	"clientConfig":  cc,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//content, err := configClient.GetConfig(vo.ConfigParam{
	//	DataId: global.NacosConfig.DataId,
	//	Group:  global.NacosConfig.Group})
	//
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Println(content) //字符串 - yaml
	////想要将一个json字符串转换成struct，需要去设置这个struct的tag
	//err = json.Unmarshal([]byte(content), &global.ServerConfig)
	//if err != nil {
	//	zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	//}
	//fmt.Println(&global.ServerConfig)

}
