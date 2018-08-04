build:
#一定要注意Makefile中的缩进，否则执行makebuild可能报错：make:Nothingtobedonefor`build'.
#protoc命令前边是一个Tab，不是四个空格
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. proto/user.proto
#告知Go编译器生成二进制文件的目标环境：amd64CPU的Linux系统
	GOOS=linux GOARCH=amd64 go build
#根据当前目录下的Dockerfile生成名为consignment-service的镜像
	docker build -t user-service .
run:
#在Dockeralpine容器的50001端口上运行consignment-service服务
#可添加-d参数将微服务放到后台运行
	docker run -p 50051:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns \
	user-service
