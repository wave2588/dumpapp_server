
#### 查看未砸壳列表
GET /admin/ipa/dump_order  
参数:  
offset: 0  
limit: 10
```json
{
  "data": [{
    "member": {
      "name": "xxxx"
    },
    "ipa_id": 111,
    "ipa_name": "xxx",
    "ipa_version": "1.1.1" 
  }],
  "paging": {}
}
```

#### 删除无需砸壳的订单
Delete /admin/ipa/dump_order    
参数:  
ipa_id string  
ipa_version string
