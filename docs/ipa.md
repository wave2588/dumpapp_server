## ipa 相关文档

### 某个时间范围内 ipa 的排名情况
GET /ipa/ranking  
```json
{
  "start_at": 1630425600, /// 可不传
  "end_at": 1631343007  /// 可不传
}
```  
Response:
```json
{
    "paging": null,
    "data": [
        {
            "ipa_id": "12345",
            "name": "xxx",
        }
    ]
}  
```


### 获取单个 ipa 信息
GET /ipa/{ipa_id}  
```json
{
  "name": "xxx",
  "bundle_id": "xxx.com",  
  "version": "1.1.1"
}
```  
Response:
```json
{}
```

### 发送获取最新 ipa 信息给管理员
GET /ipa/{ipa_id}/latest
```json
{
  "name": "xxx",
  "bundle_id": "xxx.com",  
  "version": "1.1.1"
}
```  
Response:
```json
{}
```


### 获取 ipa 所有版本
GET /ipa/{country}/{ipa_id}  
例如: /ipa/cn/1111
```
{
    "data": 
}
```


### 获取某个类型所有 ipa 列表
GET /ipa/{ipa_type}/list
```json
{
  "paging": {
    "is_end": true,
    "is_start": true,
    "next": "/api/ipa/normal/list?limit=10&offset=10",
    "previous": "/api/ipa/normal/list?limit=10&offset=0",
    "totals": 239
  },
  "data": [
    {
      "id": "11111",
      "name": "xxxx",
      "bundle_id": "xxxxx",
      "created_at": 1637728013,
      "updated_at": 1639460435,
      "versions": [
        {
          "id": "2777",
          "version": "2.3.0",
          "ipa_type": "normal",
          "created_at": 1637728013,
          "updated_at": 1639406962
        },
        {
          "id": "3088",
          "version": "2.2.0",
          "ipa_type": "normal",
          "describe_url": "",
          "created_at": 1639460435,
          "updated_at": 1639460435
        }
      ]
    }
  ]
}
```