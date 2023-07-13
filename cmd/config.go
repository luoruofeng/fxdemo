package cmd

import "flag"

var ConfigFilePathMap map[string]string

//go run  . -cnf="./conf/conf.json"
func GetConfigFilePath() map[string]string {
	if ConfigFilePathMap != nil {
		return ConfigFilePathMap
	}
	ConfigFilePathMap = make(map[string]string)
	configPath := flag.String("cnf", "", "Config file path,The default value is in the current folder.")
	//TODO 如果添加的组件也需要启动命令时设置配置文件路径，可以在这里添加
	flag.Parse()

	ConfigFilePathMap["cnf"] = *configPath
	//TODO 如果添加的组件也需要启动命令时设置配置文件路径，可以在这里添加

	return ConfigFilePathMap
}

func NewConfigPathMap() map[string]string {
	return GetConfigFilePath()
}
