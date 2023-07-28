# dousheng
### 1.更改配置

更改 constants/constant.go 中的地址配置

### 2.建立基础依赖

```shell
docker-compose up
```

### 3.运行feed微服务

```shell
cd cmd/feed
sh build.sh
sh output/bootstrap.sh
```

### 4.运行publish微服务

```shell
cd cmd/publish
sh build.sh
sh output/bootstrap.sh
```

### 5.运行user微服务

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 6.运行favorite微服务

```shell
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

### 7.运行comment微服务

```shell
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

### 8.运行relation微服务

```shell
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

### 9.运行api微服务

```shell
cd cmd/api
chmod +x run.sh
./run.sh
```

### 10.Jaeger链路追踪 

在浏览器上查看`http://127.0.0.1:16686/`
