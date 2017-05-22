# 《USTC高级软件工程》课程实验
一共7次实验，分别在7个文件夹中，含C源码、实验报告、Go实现

> 说明：
> 1. lab1 ~ lab3 在终端执行，需要配置环境变量
> 2. lab4 需要配置 beego & bee



# Lab-gui 说明
《USTC 高级软件工程》实验分为 lab1-7，从 `lab4` 开始已经实现 GUI 版本，传送门：[https://github.com/csxiaoyaojianxian/GoStudy/tree/master/gowork/src/lab-gui](https://github.com/csxiaoyaojianxian/GoStudy/tree/master/gowork/src/lab-gui)

![截图](https://raw.githubusercontent.com/csxiaoyaojianxian/GoStudy/master/gowork/src/lab-gui/static/images/pic.png)

## 开发框架

使用 beego 框架 + beego-orm +bee

## 安装说明

### 1. 使用 go get 下载安装 beego 与 bee 工具

```
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee
```
### 2. 使用 bee 工具初始化 beego 项目

```
$ bee new myapp
```
### 3. 使用 bee 工具热编译 beego 项目

```
$ bee run myapp
```
### 4. 使用 go get 下载安装 beego ORM

```
$ go get github.com/astaxie/beego/orm
```


> Powered by SUNSHINE STUDIO
>
> Author : CS逍遥剑仙
>
> Homepage : [www.csxiaoyao.com](http://www.csxiaoyao.com)](www.csxiaoyao.com)