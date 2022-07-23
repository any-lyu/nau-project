# nau 毕业进阶作业
## 家校通融合教育移动WEB项目

---

##### family-joint-school 服务端代码 （golang+gin+ent）

##### family-joint-school-front 前端代码 (vue.js+vant)

---
演示地址 http://124.221.108.233:8080/ （基于移动设备设计的界面，建议用手机打开或者浏览器设置手机显示模式）
学生账号
账号：stu 
密码 123456

老师账号
账号：teacher
密码 123456

---
应用截图

![img.png](image/img_5.png) ![img.png](image/img.png) ![img.png](image/img_1.png)
![img.png](image/img_2.png) ![img.png](image/img_3.png) ![img.png](image/img_4.png)
![img.png](image/img_6.png) ![img.png](image/img_7.png) 

---
服务端接口包括

```
登录
POST http://localhost:8888/login
Content-Type: application/json

{
  "account": "stu",
  "password": "123456"
}

> {"code": 0, "msg": "","data": {"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4","user": {"id": 1,"nickname": "stu_test","avatar": "https://img.syt5.com/2019/0912/20190912111829683.jpg.420.240.jpg","mobile": "13555667878","class_id": 1, "created_at": "2022-07-21T15:16:01+08:00","updated_at": "2022-07-21T15:17:36+08:00"},"permission": 1,"permission_name": "student"}}
```
```
班级列表
POST http://localhost:8888/api/class/list
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{}
```
```
通知列表
POST http://localhost:8888/api/notice/list
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{
  "cursor": ""
}
```
```
作业列表
POST http://localhost:8888/api/homework/list
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{
  "cursor": ""
}
```
```
提交作业
POST http://localhost:8888/api/homework/submit
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{
  "homework_id": 1,
  "title": "语文背诵作业",
  "content": "鹅鹅鹅，曲项向天歌，白毛浮绿水，红掌拨清波",
  "pics": [
    {
      "url": "https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo.jpg"
    }
  ]
}
```
```
已提交作业详情 status = 0 未提交作业 status = 1 已提交作业
POST http://localhost:8888/api/homework/user/detail
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{
  "homework_id": 1
}

发布通知
POST http://localhost:8888/cms/notice/publish
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{
  "title": "紧急通知",
  "content": "全员放假",
  "level" : 1
}
```
```
发布作业
POST http://localhost:8888/cms/homework/publish
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{
  "title": "语文作业",
  "content": "背诵唐诗三百首",
  "pics": [
    {
      "url": "https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo.jpg"
    }
  ]
}
```
```
学生已提交作业列表
POST http://localhost:8888/cms/homework/user/list
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjY4NzU3NTAwLCJpc3MiOiJzeXN0ZW0ifQ.pW8A5E3QQq4ZeYvJ2uQobrsHKXGbIyJJyz_MqGXiXD4

{
  "cursor": ""
}
```

