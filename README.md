# 仿小米商城后端服务
## 介绍

模拟小米官网的后端服务

项目运行环境

| go          | 1.15+   |
| ------------ | ------ |
| mysql        | 5.7+   |
| redis        | 6.2.1  |
| gin        | latest |
| gorm        | latest |



## 代码结构说明
本项目使用了[ginbro](https://github.com/dejavuzhou/ginbro)脚手架

```
mi-gin

——conf      #配置加载

——controller   #接口控制层

——models   #实体层和db 的 crud

——pkg   #通用组件

——static    #静态文件

——swagger   #在线接口文档(ginbro 生成的)

——task   #定时任务

```

## 后端启动

[代码地址Gitee](https://gitee.com/gitxys/mi_gin)

[代码地址GitHub](https://github.com/xuyisu/mi_gin)

###### 1.先下载代码

###### 2.配置代理

![image-20211230101706606](doc/images/image-20211230101706606.png)

###### 3.下载依赖组件

```
go mod download
```

###### 4.将doc 下面的mi-mall.sql 导入到mysql数据库,同时  启动mysql  和 redis

**后端运行先配置数据库（mysql 和redis）**

```
[app]
    name = "go-gin"
    addr ="localhost:8081"
    secret = "qazwsxecd"
    env = "local" # only allows local/dev/test/prod
    log_level = "debug" # only allows debug info warn error fatal panic
    enable_not_found = true # if true and static_path is not empty string, all not found route will serve static/index.html
    enable_swagger = false
    enable_cors = true  # true will case 403 error in swaggerUI  may cause api perform decrease
    enable_sql_log = true # show gorm sql in terminal
    enable_https = false # if addr is a domain enable_https will works
    enable_cron = false # is enable buildin schedule job
    time_zone = "Asia/Shanghai"
    api_prefix = "" #  api_prefix could be empty string,            the api uri will be api/v1/resource
    static_path = "./static/"  # path must be an absolute path or relative to the go-build-executable file, may cause api perform decrease
    mem_expire_min = 60 # memory cache expire in 60 minutes
    mem_max_count = 1024000 # memory cache maxium store count
[mysql]
    addr = "127.0.0.1:3306"
    user = "root"
    password = "123456"
    database = "mi-mall"
    charset = "utf8mb4"
[redis]
    addr = "127.0.0.1:6379" # 127.0.0.1:6379 empty string will not init the redis db in models package
    password = "123456"
    db_idx = 0
    session_expire = 3600
```

###### 5.启动

切到命令行窗口，并输入 `bee run`，看到以下结果即代表启动成功

```
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v2.0.2
2021/12/30 10:17:54 WARN     ▶ 0001 Running application outside of GOPATH
2021/12/30 10:17:54 INFO     ▶ 0002 Using 'mi-beego' as 'appname'
2021/12/30 10:17:54 INFO     ▶ 0003 Initializing watcher...
mi-beego/pkg/utils
mi-beego/pkg/lib
mi-beego/pkg/log
mi-beego/third_party/mysql
mi-beego/third_party/swagger
mi-beego/third_party/redis
mi-beego/models
mi-beego/pkg/filter
mi-beego/controllers
mi-beego/routers
mi-beego
2021/12/30 10:17:58 SUCCESS  ▶ 0004 Built Successfully!
2021/12/30 10:17:58 INFO     ▶ 0005 Restarting 'mi-beego.exe'...
2021/12/30 10:17:58 SUCCESS  ▶ 0006 './mi-beego.exe' is running...
2021/12/30 10:17:58.768 [I] [parser.go:85]  D:\workspaces\mywork\gitee\mi-beego\controllers no changed

2021/12/30 10:17:58.769 [I] [server.go:241]  http server Running on http://:8081

```

浏览器输入 http://127.0.0.1:8081/swagger

![image-20211230102018377](doc/images/swagger.png)

## 前端启动

项目地址https://gitee.com/gitxys/mi_vue

[代码地址Gitee](https://gitee.com/gitxys/mi_vue)

[代码地址GitHub](https://github.com/xuyisu/mi_vue)

1.下载代码到本地

2. 控制台先安装依赖包

```
npm install 
```

3.启动

```
npm run serve
```

## 页面介绍

浏览器输入http://localhost:8080 将看到一下页面

![](images/index.png)

登录:**用户名/密码**  admin/123456

![image-20211219223115929](doc/images/login.png)

购物车

![image-20211219223220837](doc/images/cart.png)

订单确认

![image-20211219223323684](doc/images/order-confirm.png)

订单结算(彩蛋！！！！   这里的结算做了特殊处理)

![image-20211219223406482](doc/images/pay.png)

订单列表

![image-20211219223507791](doc/images/order.png)





亲，留个star 吧

