## 按次下载 ipa 文档

#### 获取单个 ipa, 如果没有下载次数会抛错
GET /ipa/{ipa_id}  
参数:  
name: "ipa_name"  
```json
{
  "id": "1111",
  "versions": [
    {
      "version": "1.1.1"
    }
  ]
}
```

#### 查看是否可以下载  
GET  /v2/ipa/{ipa_id}/check_can_download  
参数:   
version: "1.1.1"    
```json
{
  "can_download": true
}
```

#### 获取 ipa 的下载地址, 如果没有下载次数会抛错
GET /v2/ipa/{ipa_id}/download_url  
参数:   
version: "1.1.1"    
```json
{
  "open_url": "xxxxx" 
}
```
