
#### 查看未砸壳列表
GET /admin/ipa/dump_order  
参数:  
offset: 0  
limit: 10
```json
{
    "paging": {
        "is_end": true,
        "is_start": true,
        "next": "/api/admin/ipa/dump_order?limit=10&offset=10",
        "previous": "/api/admin/ipa/dump_order?limit=10&offset=0",
        "totals": 1
    },
    "data": [
        {
            "demander_member": {
                "id": "1431956649500741632",
                "email": "15711367321@163.com",
                "status": "normal",
                "phone": "15711367321",
                "download_count": 24,
                "vip": {
                    "is_vip": false
                },
                "invite_url": "https://www.dumpapp.com/register?invite_code=yTWJab",
                "devices": [
                    {
                        "id": "1457748144790966272",
                        "udid": "00008101-001959310E92001E",
                        "product": "iPhone13,4",
                        "created_at": 1636389216,
                        "updated_at": 1636389216,
                        "certificates": [
                            {
                                "id": "1457748418188283904",
                                "created_at": 1636389281,
                                "updated_at": 1636389281,
                                "p12_is_active": true
                            }
                        ]
                    }
                ],
                "created_at": 1630240044,
                "updated_at": 1637079616
            },
            "other_demander_member": [
                {
                    "id": "1431956649500741632",
                    "email": "15711367321@163.com",
                    "status": "normal",
                    "phone": "15711367321",
                    "download_count": 24,
                    "vip": {
                        "is_vip": false
                    },
                    "invite_url": "https://www.dumpapp.com/register?invite_code=yTWJab",
                    "devices": [
                        {
                            "id": "1457748144790966272",
                            "udid": "00008101-001959310E92001E",
                            "product": "iPhone13,4",
                            "created_at": 1636389216,
                            "updated_at": 1636389216,
                            "certificates": [
                                {
                                    "id": "1457748418188283904",
                                    "created_at": 1636389281,
                                    "updated_at": 1636389281,
                                    "p12_is_active": true
                                }
                            ]
                        }
                    ],
                    "created_at": 1630240044,
                    "updated_at": 1637079616
                }
            ],
            "operator_member": null,
            "ipa_id": 1441513936,
            "ipa_version": "",
            "ipa_name": "CHAO - 潮流 × 兴趣 × 记录",
            "ipa_bundle_id": "",
            "ipa_app_store_link": "",
            "created_at": 1637249147,
            "updated_at": 1637249147
        }
    ]
}
```

#### 删除无需砸壳的订单
Delete /admin/ipa/dump_order    
参数:  
ipa_id string  
ipa_version string
