## 账号相关 文档

#### 发送邮箱验证码
``` 
POST /email/captcha  
参数:    
email: "11@1.com"
```
 
 
#### 发送手机验证码
```
POST /phone/captcha  
参数:   
phone: "157xxx"
```
 
#### 注册
```
POST /register  
参数:  
email: "11@1.com"               必填  
email_captcha: "888888"         必填  
phone: "157"                    非必填    
phone_captcha: "666666"         非必填   
password: "xxxx"                必填
```

#### 登录
```
POST /login
参数:     email 和 phone 只用传一个即可
email: "xxxx"           
phone: "152"
password: "xxxxxx"  
```
