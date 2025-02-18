Transfeed
===

一个简单，易部署的开源RSS翻译器，前端使用`Vue`，后端使用`Go`。


## 特点

- AI翻译（支持大模型）
- 订阅源可分别设置不同翻译器
- 多用户隔离
- 多语言
- 易部署（仅包含一个可执行文件）
- 数据库轻量（使用sqlite）
- 支持自定义刷新时间



查看 👀 [demo](https://wsh233.cn/webapp/transfeed)     用户名：`demo`， 密码：`demo1234`



## 运行

针对不同操作系统，下载可执行文件，运行后，前往浏览器打开网页`http://127.0.0.1:8090`。

查看帮助

```bash
transfeed -h
```



关于配置项

```
-admin-name string
  初始化管理员用户名，默认”admin“
-admin-password string
  初始化管理员密码，默认“admin1234”
-database-url string
  数据库路径，默认当前目录 
-disable-swagger
  是否关闭swagger，默认关闭
-port int
  启动端口，默认8090
-refresh-hours int
  刷新间隔小时 默认4小时
-url-prefix string
  路由前缀，默认无
-version
  显示版本信息
```



## 部署

如果使用`Nginx`转发，必须设置路由前缀`url-prefix`，例如前端访问地址在`http://www.server.com/transfeed`，nginx配置设置如下

```text
location ~ ^/transfeed/(.*)$ {
    proxy_pass http://127.0.0.1:8091/$1;
    proxy_set_header Host            $host:$server_port;
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
}

```



同时设置`url-prefix`为

```bash
transfeed -url-prefix /transfeed
```



Transfeed，为前后端分开，支持自己编写个性化的前端应用，接口部分请查看`swagger`文档。



# 说明

开发过程中用到许多开源项目，非常感谢❤️。



