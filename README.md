
##  项目结构

```shell
├─.idea
├─api
├─conf
├─database
│  ├─cache
│  └─db
│      ├─dao
│      └─model
├─es
├─middleware
├─mq
├─response
├─route
├─service
├─types
└─util
```
- api 放置路由相关接口
- conf mysql， redis，oss等相关配置
- database mysql，redis相关操作
- es elastic操作
- middleware JWT认证操作
- mq 聊天模块的操作，使用rabbitmq实现
- response 返回数据结构
- route 路由组
- service 具体的各种操作
- types 存放请求相关结构体
- util 生成token


###  项目完成情况

- bonus部分完成了1，2，6，7


### 接口文档地址

- https://apifox.com/apidoc/shared-2f99121d-895b-4576-a9f2-2e985a8087e9