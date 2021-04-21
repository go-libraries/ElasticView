  

技术栈主要用： Vue + golang 

欢迎大家提出自己的issue。

ElasticView
-----------
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/360EntSecGroup-Skylar/ElasticHD/main)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/1340691923/ElasticView/blob/main/LICENSE)
[![Release](https://img.shields.io/github/release/360EntSecGroup-Skylar/ElasticHD.svg?label=Release)](https://github.com/360EntSecGroup-Skylar/ElasticHD/releases/latest)
> ElasticView 是一款用来监控elasticsearch状态和操作elasticsearch索引的web可视化工具。它由golang开发而成，具有部署方便，占用内存小等优点
 * es连接树管理（更方便的切换测试/生产环境）
 * 支持页面级别的权限管理
 * 支持sql转换成dsl语法
 * 更方便的重建索引
 * 任务管理
 * 备份管理
 * 可下载查询内容
 * 可进行索引创建，映射创建，别名创建，索引删除等操作
 
## ElasticView 部分截图

![image](https://raw.githubusercontent.com/1340691923/ElasticView/main/show_img/1.png)
![image](https://raw.githubusercontent.com/1340691923/ElasticView/main/show_img/2.png)
![image](https://raw.githubusercontent.com/1340691923/ElasticView/main/show_img/3.png)
![image](https://raw.githubusercontent.com/1340691923/ElasticView/main/show_img/4.png)
![image](https://raw.githubusercontent.com/1340691923/ElasticView/main/show_img/5.png)
![image](https://raw.githubusercontent.com/1340691923/ElasticView/main/show_img/6.png)
![image](https://raw.githubusercontent.com/1340691923/ElasticView/main/show_img/7.png)

## 应用程序下载
[下载地址]( https://github.com/1340691923/ElasticView/releases/) 支持操作系统：windows，linux

## 安装教程

 第一步: 下载 对应系统的ElasticView.zip
 第二步: 解压 ElasticView.zip，新建mysql数据库 es_view 导入es_view.sql 
 第三步: 修改config.json文件中的 数据库连接信息，日志存放目录和应用启动端口等配置信息
 第四步：windows：双击ElasticView.exe  linux：chmod +x ElasticView && nohup ./ElasticView 
 第五步：浏览器访问对应ip:端口，初始化用户名：admin，密码：admin
 
 ## 手动编译
 第一步：git clone git@github.com:1340691923/ElasticView.git
 第二步：cd views && cnpm install (安装前端依赖)
 第三步：npm run build:stage （打前端正式包）
 第四步：cd .. &&   statik -src=views/dist -f (打包静态资源进二进制)
 第五步：windows: go build -o ElasticView.exe`|| linux : 双击build.bat // 编译go可执行程序 ``