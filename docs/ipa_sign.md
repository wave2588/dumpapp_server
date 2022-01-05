## ipa sign 相关文档

### 创建签名任务
POST /ipa/sign  
certificate_id string   /// 证书 id 
ipa_version_id string   /// 需要注意，这个是 ipa version id


### 获取签名列表  
GET /ipa/sign  
```json
{
  "data": [{
    "id": "111",
    "ipa_id": "222",
    "status": "processing"
  }],
  "paging": {
    "is_end": false
  }
}
```