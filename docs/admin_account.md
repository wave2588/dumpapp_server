#### 添加账号
POST /admin/account  
参数:  
email: "xxx.com"  
phone: "157"        /// 国外地区手机号不用填  
password: "12321"
```json
ok
```

#### 修改账号，暂时只支持修改密码
PUT /admin/account  
参数:  
email: "xxx.com"  
password: "12321"
```json
ok
```
