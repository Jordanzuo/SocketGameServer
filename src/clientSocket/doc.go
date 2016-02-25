/*
为客户端提供服务的Socket端；
在config.go中的SetParam方法被调用之前不能使用其中定义的配置文件，所以不能定义需要用到config.go中的配置的init方法
*/
package clientSocket
