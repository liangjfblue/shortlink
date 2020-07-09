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


