## 按次下载 ipa 文档

#### 用户信息结构里新加了 download_number 字段
```json
{
  "name": "xxx",
  "download_number": 2
}
```

#### 获取单价
GET /v2/member/vip
```json
{
  "price": 123213
}
```

#### 获取支付链接:
POST /v2/member/vip?number=1
```json
{
  "open_url": "xxxxx" 
}
```

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

#### 获取 ipa 的下载地址, 如果没有下载次数会抛错
GET /v2/ipa/{ipa_id}/download_url  
参数:   
version: "1.1.1"    
```json
{
  "open_url": "xxxxx" 
}
```