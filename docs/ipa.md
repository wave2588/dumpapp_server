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
