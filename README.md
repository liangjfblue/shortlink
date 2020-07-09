# 短连接系统

## 进制法
- 通过redis原子自增作为发号器, redis是cluster集群架构, 避免单点; 
- 通过10进制转62进制缩短短码; 
- 借助mysql持久化长短连接映射避免丢失数据, 借助缓存提高查询效率

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


## 使用
- 请求生成短连接
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
    

- 请求短连接跳转
```[GET] 172.16.7.16:9099/sl/g9```


- 根据短连接获取长连接信息
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