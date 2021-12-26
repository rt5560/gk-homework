# Kratos Project Template

## Environment Requirements
* go
* protoc
* protoc-gen-go

The GO111MODULE should be enabled.
> go env -w GO111MODULE=on

## Install Kratos
```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade
```
## Create a service
```
# 创建项目模板
kratos new server

cd server
# 拉取项目依赖
go mod download

# 生成proto模板
kratos proto add api/server/server.proto
# 生成proto源码
kratos proto client api/server/server.proto
# 生成server模板
kratos proto server api/server/server.proto -t internal/service

#install wire
go get github.com/google/wire/cmd/wire

#生成所有proto源码、wire等等
go generate ./...

# 运行程序
kratos run
```

