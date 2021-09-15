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
