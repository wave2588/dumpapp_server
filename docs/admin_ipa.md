
#### 上传 ipa
POST /admin/ipa  
参数:  
```json
{
  "ipas": [
    {
      "ipa_id": "111",
      "name": "xxxxx",
      "bundle_id": "xxxx",
      "versions": [
        {
          "version": "1.11",
          "token": "xxx.ipa",
          "ipa_type": "normal", /// normal 或者 crack
          "is_temporary": true, /// 标记三天后自动删除
          "describe_url": "https://xxx" /// 一般情况下 ipa_type == crack 才会用到
        }   
      ]   
    } 
  ],
  "is_send_email": true
}
```

#### 删除 ipa
DELETE /admin/ipa
```json
{
  "ipa_id": "xxx",
  "ipa_type": "normal",
  "ipa_version": "1.11",  /// 指定删除某个版本
  "is_retain_latest_version": true  /// 是否保留最新版本
}
```


#### 批量删除 ipa
DELETE /admin/batch_ipa
```json
{
  "ipa_ids": ["111","222"],
  "ipa_type": "normal",
  "is_retain_latest_version": true /// 是否保留最新版本
}
```
