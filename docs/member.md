
### 获取用户已下载列表
GET member/self/download_record
```json
{
  "paging": {
    "is_end": false,
    "is_start": true,
    "next": "/api/member/self/download_record?limit=10&offset=10",
    "previous": "/api/member/self/download_record?limit=10&offset=0",
    "totals": 119
  },
  "data": [
    {
      "id": "7502",
      "version": "11.11",
      "status": "used",
      "ipa_type": "normal",
      "ipa": {
        "id": "112233",
        "name": "xxxx",
        "bundle_id": "xxx.com",
        "created_at": 1636017484,
        "updated_at": 1640346316,
        "versions": [
          {
            "id": "3341",
            "version": "1.28.0.46101",
            "ipa_type": "normal",
            "describe_url": "",
            "describe": "",
            "created_at": 1640346316,
            "updated_at": 1640497962
          }
        ]
      },
      "created_at": 1640237958,
      "update_at": 1640343982
    }
  ]
}
```