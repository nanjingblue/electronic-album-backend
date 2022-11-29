# electronic-gallery-backend
电子相册的后端，实现照片云存储，并支持社交圈分享。
## TODO
* 功能需求
    - [x] 登录注册
    - [x] 账号功能
    - [x] 相册增删改查
    - [ ] 图片增删改查
    - [ ] Post增删改查
    - [ ] 管理员功能
* 项目开发
  - [ ] 接口文档
  - [ ] 日志
  - [x] RESTFUL API
  - [x] JWT鉴权
  - [x] OSS对象存储
  - [x] 跨域处理
  - [x] Dockerfile
## 配置
配置文件在 `configs/example.config.yaml`，需要重命名为 `config.yaml` 。    
数据库支持 Mysql 和 Sqlite，可在 type 项填写。   
支持阿里云对象存储，需要先去阿里云创建Bucket，并将相关信息填入配置项。
## 运行
```Bash
git clone https://github.com/nanjingblue/electronic-gallery-backend.git
cd electronic-gallery-backend
go mod tidy
go run main.go
```
## 效果

![image-20221129202719860](https://raw.githubusercontent.com/nanjingblue/gallery/master/images/202211292027964.png)

![image-20221129202656834](https://raw.githubusercontent.com/nanjingblue/gallery/master/images/202211292026910.png)
