/*
项目配置逻辑处理包，初始化所有的配置内容，其它代码需要配置时都从此包内获取
*/
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// 读取配置文件
// configFileName：配置文件名称
// 返回值：
// 配置文件内容
func ReadConfigFile(configFileName string) map[string]interface{} {
	// 读取配置文件（一次性读取整个文件，则使用ioutil）
	bytes, err := ioutil.ReadFile(configFileName)
	if err != nil {
		panic(errors.New(fmt.Sprintf("读取配置文件%s的内容出错，错误信息为：%s", configFileName, err)))
	}

	// 使用json反序列化
	config := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &config); err != nil {
		panic(errors.New(fmt.Sprintf("反序列化配置文件%s的内容出错，错误信息为：%s", configFileName, err)))
	}

	return config
}

// 从config配置中获取int类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
func GetIntConfig(config map[string]interface{}, configName string) int {
	configValue, ok := config[configName]
	if !ok {
		panic(errors.New(fmt.Sprintf("不存在名为%s的配置或配置为空", configName)))
	}
	configValue_float, ok := configValue.(float64)
	if !ok {
		panic(errors.New(fmt.Sprintf("%s必须为int型", configName)))
	}

	return int(configValue_float)
}

// 从config配置中获取string类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
func GetStringConfig(config map[string]interface{}, configName string) string {
	configValue, ok := config[configName]
	if !ok {
		panic(errors.New(fmt.Sprintf("不存在名为%s的配置或配置为空", configName)))
	}
	configValue_string, ok := configValue.(string)
	if !ok {
		panic(errors.New(fmt.Sprintf("%s必须为string型", configName)))
	}

	return configValue_string
}
