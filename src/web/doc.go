/*
提供Web服务的包，用于提供一下功能：
1、封号、解封
2、禁言、解禁
3、推送消息
4、重新加载敏感词
在config.go中的SetParam方法被调用之前不能使用其中定义的配置文件，所以不能定义需要用到config.go中的配置的init方法
*/
package web
