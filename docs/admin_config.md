# admin config 相关接口

### 设置配置
POST /admin/config
```json
{
  "admin_busy": true,   // 管理员是否忙碌
  "daily_free_count":10  // 每日免费限额
}
```  
Response:
```json
{
  "admin_busy": true,
  "daily_free_count": 10
}
```

### 获取配置
GET /admin/config
Response:
```json
{
  "admin_busy": true,
  "daily_free_count": 10
}
```