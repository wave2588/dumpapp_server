## ipa sign 相关文档

### 1. 创建签名任务
POST /ipa/sign  
certificate_id string   /// 证书 id   
ipa_version_id string   /// 需要注意，这个是 ipa version id


### 2. 获取签名列表  
GET /ipa/sign  
```json
{
  "data": [{
    "id": "111",
    "ipa": {
      "id": "222",
      "name": "hhhh"
    },
    "status": "processing",
    "current_ipa_version": "1.1.1",
    "current_ipa_type": "normal"
  }],
  "paging": {
    "is_end": false
  }
}
```  

### 3. 获取签名下载地址  
GET /ipa/sign/{ipa_sign_id}/url  
```json
{
  "open_url": "https://xxxx"
}
```