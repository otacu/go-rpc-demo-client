go微服务客户端demo

服务端和客户端的proto是一样的，是对微服务接口的定义。
通过修改anime.proto文件然后用命令protoc --go_out=plugins=micro:. anime.proto生成go文件（要先安装protoc.exe）。