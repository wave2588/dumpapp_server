# 用户 certificate 相关接口

### 生成证书
POST /certificate
Request body
```
{
    "udid": "xxxx"
}
``` 
Response:
```json
{
}
```   


### 下载 p12
GET /certificate/p12
Request body
```
{
    "device_id": "xxxx",
    "cer_id": "xxxx",
}
``` 



### 下载 mobileprovision
GET /certificate/mobileprovision
Request body
```
{
    "device_id": "xxxx",
    "cer_id": "xxxx",
}
``` 
