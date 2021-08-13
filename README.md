# API 文档

## 说明
该API设计，不使用RESTful规范

## 协议约定

### 版本（Versioning）
将API的版本号放入URL
```
https://api.yudylaw.com/api/v1/ （client api）(后续统一以client api示例)
https://api.yudylaw.com/v1/ （server api）
```

### 路径（Endpoint）
路径又称"终点"（endpoint），表示API的具体网址。

每个网址代表一种资源（resource），网址中使用名词，不使用动词，而且所用的名词往往与数据库的表名对应。一般来说，数据库中的表都是同种记录的"集合"（collection），因此名词使用复数。

网址中尽可能避免使用_等字符。

举例来说，API提供一套通用代码库，包括用户、消息、关系等，则它的路径应该设计成下面这样。
```
https://api.yudylaw.com/api/v1/users
https://api.yudylaw.com/api/v1/messages
https://api.yudylaw.com/api/v1/relations
```

### HTTP动词
对于资源的具体操作类型，由HTTP动词表示。

由于不使用RESTful，因此建议使用的HTTP动词只有下面2个（括号里是对应的SQL命令）

- GET（SELECT）：从服务器取出资源(一项或多项)
- POST（CREATE、UPDATE、DELETE）：在服务器 新建/修改/删除 资源

下面是一些例子

- GET /users/list    列出所有用户
- POST /users/add    新建一个用户
- POST /users/update 更新某个用户的信息
- POST /users/delete 删除某个用户的信息

### 过滤（Filter）
如果记录数量很多，服务器不可能都将它们返回给调用方。API提供参数，过滤返回结果。
下面是一些常见的参数

- ?count=10：指定返回记录的数量
- ?before=20&count=10：指定返回记录的开始位置。
- ?sort=name：指定返回结果按照哪个属性排序
- ?type=1：指定筛选条件

### 翻页（Pagination）
查询资源集合时，与翻页有关参数如下：

- after 返回比after老的count条数据（不包含after），特殊的，第一次查询时，传入after=0，返回最新的count数据
- before  返回比before新的count条数据（不包含before）
- count  返回数据条数，默认值是20，取值范围[20,100]

返回资源集合时，包含与分页有关的属性如下：
```
{
  "paging":{
    "after":123,  //下一个游标，从after位置往后取
    "before":130, //下一个游标，从before位置往前取
    "has_more":true //是否有下一页数据
  }
}
```

##### 获取最新的用户列表
```
GET /api/v1/users/list?after=0&count=20
```

##### 获取某个用户最新系统消息
```
GET /api/v1/messages/newest?uid=1110&before=300310&count=20
```

### 返回体（Response）
考虑到Client端依赖的部分框架不支持标准的HTTP状态码，譬如403、400，API总是返回200状态码，调用方通过返回体中error_code是否大于0来判断接口成功与否。
请求成功（error_code = 0）
```json
{
    "error_code": 0,
    "error_msg": "操作成功",
    "data":{}
}
```
或者返回
```json
{
    "error_code": 0,
    "error_msg": "操作成功",
    "data":{
        "some_key": "some_value",
        "list": [],
        "paging":{
          "after":123,  
          "before":130,
          "has_more":true
        }
    }
}
```
成功失败（error_code > 0）
```json
{
    "error_code": 499,
    "error_msg": "系统错误",
    "data": []
}
```

## 错误码说明
| error_code    | error_msg   |
| :----------- | :--- |
| 0 | 成功 |
| 499 | 参数错误 |
| 500 | 系统错误 |
| 604 | 认证失败，需重新登录 |
| 611 | 无效的access_token，需重新登录 |
| 612 | 无效的refresh_token，需重新登录 |
| 613 | access_token已过期，需使用refresh_token刷新access_token |
| 614 | refresh_token已过期，需重新登录 |
| 615 | 重放的请求 |
| 616 | 无效的签名 |
| 700 | RPC调用错误 |
| 701 | 数据库异常 |
| 702 | 缓存异常 |
| 703 | 记录未找到 |
| 710 | 发号器发号异常 |
| 2002 | 当天获取验证码次数过多 |
| 2003 | 操作次数频繁 |
| 2004 | 验证码已过期,请重新获取验证码 |
| 2007 | 验证码错误 |
| 2012 | 昵称太长 |
| 2013 | 敏感词 |
| 2014 | 昵称已被占用 |
| 2015 | 签名带有敏感词 |
| 2016 | 手机号或者验证码格式错误 |
| 2017 | 请输入有效手机号 |
| 2018 | 该账号已封 |
| 2025 | 微信access_token无效 |
| 2026 | 该手机号已经绑定微信 |
