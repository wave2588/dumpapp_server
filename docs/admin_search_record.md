## admin_search_record 后台文档

### 某个时间范围内用户获取 ipa 的排名情况
GET /admin/search/record  
参数:  
offset: 0  
limit: 10
start_at: 1630425600        /// 可不传  
end_at: 1631343007          /// 可不传
```json
{
    "paging": null,
    "data": [
        {
            "ipa_id": "12345",
            "name": "xxx",
            "count": 735, /// 获取次数
        }
    ]
}
