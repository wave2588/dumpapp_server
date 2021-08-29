## admin 后台文档


#### 增加次数
POST /admin/member/download_number  
参数:  
email: "xxx@163.com"  
number: 10
```json
{
}
```

#### 删除次数
DELETE /admin/member/download_number 
参数:  
email: "xxx@163.com"  
number: 10
```json
{
}
```

#### 删除 ipa
DELETE /admin/ipa  
参数:  
ipa_id: "111"  必填  
version: "1.1.1"  选填  
is_retain_latest_version: true  选填  是否保留最新版本


#### 批量删除 ipa 
DELETE /admin/batch_ipa  
参数:  
ipa_ids: ["1", "2"]  
is_retain_latest_version: true  选填  是否保留最新版本


