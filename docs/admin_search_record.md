## admin_search_record 后台文档

### 某个时间范围内用户获取 ipa 的排名情况
GET /admin/search/record  
参数:  
start_at: 1630425600  
end_at: 1631343007  
```json
{
    "paging": null,
    "data": [
        {
            "ipa_id": "12345",
            "name": "xxx",
            "count": 735, /// 获取次数
            "latest_at": 1631270330  /// 最新获取的时间
        }
    ]
}
