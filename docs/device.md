## 设备相关
#### 获取绑定设备的二维码
GET /device/config/qr_code  
Response:  
image 实体


## 用户 certificate 相关接口

### 生成证书
POST /certificate  
Request body
```
{
    "udid": "xxxx"
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
