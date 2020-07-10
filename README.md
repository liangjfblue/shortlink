# 短连接系统

[![Build Status](https://travis-ci.org/liangjfblue/shortlink.svg?branch=master)](https://travis-ci.org/liangjfblue/shortlink)
[![codecov](https://codecov.io/gh/liangjfblue/shortlink/branch/master/graph/badge.svg)](https://codecov.io/gh/liangjfblue/shortlink)
![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

[我的短连接系统开源项目](https://github.com/liangjfblue/shortlink)

## 短连接算法
### 进制法
### redis原子递增算法
- 通过redis原子自增作为发号器, redis是cluster集群架构, 避免单点; 
- 通过10进制转62进制缩短短码; 
- 借助mysql持久化长短连接映射避免丢失数据, 借助缓存提高查询效率


### 基于etc/zkd的元数据中心发号算法
- etcd作为元数据中心存储已分发id的id段区间(比如区间是10000, 0~10000 10001~20000...)
- 各个发号服务启动时到etcd/zk取id区间, 并且更新etcd/zk的id段区间为下一阶段(同一个事务/分布式锁)
- 即使有发号服务挂了/重启也只是丢失一段id区间, id还是保持递增

- web api
- core
- generate
- db/redis

web api:
- 提供短连接生成/获取接口

core:
- 提供长短连接转换逻辑处理接口
- GetShortLinkByLongLink(string) (string, error)
- GetShortLinkByLongLink(string) (string, error)

generate:
- 提供短连接id生成接口
- Create(opts ...Option) (int64, error)

db/redis:
- 持久化存储长短连接映射
- 缓存长短连接映射

### 单机压测结果
#### 机器配置
- cpu: 4  Intel(R) Core(TM) i5-4210M CPU @ 2.60GHz
- 内存: 8G


## 项目组件
- gin
- redis
- mysql
- grpc
- etcd
- gorm

## 使用
### Http
#### 请求生成短连接
```[POST] http://172.16.7.16:9099/v1/shorten```

**body:**

    {
        "longLink":"https://www.google.com/search?q=%E7%9F%AD%E9%93%BE%E6%8E%A5%E7%B3%BB%E7%BB%9F%E8%AE%BE%E8%AE%A1&sxsrf=ALeKk01rFpwiLcx4dNPmy5Fylgy5lvHZRg:1594121387265&ei=q1wEX8_mD5vr-Qbf2ZmYCA&start=10&sa=N&ved=2ahUKEwiP6JD4hLvqAhWbdd4KHd9sBoMQ8NMDegQIDBBG&biw=1745&bih=852"
    }

**respond:**

    {
        "code": 1,
        "msg": "ok",
        "data": "172.16.7.16:9099/sl/g9"
    }
    

#### 请求短连接跳转
```[GET] 172.16.7.16:9099/sl/g9```


#### 根据短连接获取长连接信息
```[GET] http://172.16.7.16:9099/v1/shorten?shortLink=172.16.7.16:9099/sl/g9```

**respond:**

    {
        "code": 1,
        "msg": "ok",
        "data": {
            "shortId": 1001,
            "createdAt": "2020-07-09T11:40:56+08:00",
            "updatedAt": "2020-07-09T11:40:56+08:00",
            "shortCode": "g9",
            "longLink": "https://www.google.com/search?q=%E7%9F%AD%E9%93%BE%E6%8E%A5%E7%B3%BB%E7%BB%9F%E8%AE%BE%E8%AE%A1&sxsrf=ALeKk01rFpwiLcx4dNPmy5Fylgy5lvHZRg:1594121387265&ei=q1wEX8_mD5vr-Qbf2ZmYCA&start=10&sa=N&ved=2ahUKEwiP6JD4hLvqAhWbdd4KHd9sBoMQ8NMDegQIDBBG&biw=1745&bih=852"
        }
    }


### 自定义短码生成短连接
```[POST] http://172.16.7.16:9099/v1/shorten/customize```

body:

    {
        "shortCode":"waini",
        "longLink":"https://www.baidu.com"
    }

respond:

    {
        "code": 1,
        "msg": "ok",
        "data": "172.16.7.16:9099/sl/waini"
    }


### grpc
#### 请求生成短连接
#### 根据短连接获取长连接信息


### 切换发号器算法
通过插件化的方式, 支持用户自定义底层的发号算法

默认是redis的原子递增算法, 现在要切换为基于etcd的元数据中心发号算法

步骤:
- 1.实现基于etcd的元数据中心发号算法
实现 ```service/generate/generate.go``` 的```IGenerate```接口即可, 详细可查看 ````service/generate/etcd_section_generator.go````

- 2.在server/server.go注册

```go
generate.RegisterGenerate("etcd_section_generator", generate.NewEtcdSectionGenerate())
generate.SetDefaultGenerate("etcd_section_generator")
```



## 分布式部署
### redis原子自增分布式
负载均衡器+n个短连接服务

比如起10个短连接服务, 每个分别以0~9结尾, 负债均衡器采用轮训方式转发请求, 每次短连接服务由单点的自增1改为自增10

这样, 就可以分布式部署并且保持id递增, 即使机器挂了, 可以重启机器并且设置起始生成id为当前同个数量级的或者后面数量级的数字, 继续作为发号器服务


### etcd/zk元数据中心号段发放分布式
负载均衡器+n个短连接服务

短连接服务启动时会先获取分布式锁, 然后去元数据中心取id号段, 并且更新元数据中心id号段

基于etcd+发号服务组成的分段分发, etcd作为元数据中心存储已分发id的id段区间(比如区间是10000), 各个发号服务到etcd取id区间, 并且更新etcd的id段区间为下一阶段(同一个事务)
所以, 即可有发号服务挂了也只是丢失一段id区间, id还是保持递增


## 优化
- 分库分表
- 读写分离
- 缓存
- 防攻击
- 限流
- 短码反推?(不希望:洗牌算法, 希望:固定位插入字符,还原时删除就行)