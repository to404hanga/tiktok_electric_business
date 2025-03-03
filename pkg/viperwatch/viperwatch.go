package viperwatch

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitViperWatch 初始化viper并读取配置文件，通过命令行config参数传入配置文件路径，默认为config/dev.yaml
func InitViperWatch() {
	cfile := pflag.String("config", "config/dev.yaml", "配置文件路径")
	pflag.Parse()
	viper.SetConfigFile(*cfile)
	viper.WatchConfig()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
