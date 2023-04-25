# 需求分析文档

## 简介

本项目是一个互联网广告管理系统。互联网通过展示企业客户广告获得服务收益，企业客户向互联网服务提供商购买广告，购买的广告展示在服务使用界面上展示
  

## 用户需求

* ####    用户可以浏览应用页，浏览到展示的广告
* ####    管理员可以终止广告展示
* ####    企业客户可以修改密码
* ####    管理员可以查看企业信息，包括基本信息，账户信息，持有广告位，消费记录
* ####    管理员可以查看广告位信息，包括广告位状态，定价，所属企业，剩余时长等
* ####    管理员可以给广告位定价	
* ####    企业用户可以预充值
* ####    管理员可以审核企业客户注册信息，修改信息，提交的广告
* ####    管理员可以增加新的管理员账号
* ####     企业客户可以注销账号
* ####     企业客户可以查询审核进度
* ####     没有账号的企业客户可以注册账号，提供企业信息以供审核
* ####     企业客户可以修改个人信息并提交审核
* ####     企业客户可以购买广告位，上传广告信息并提交审核
* ####     企业客户可以查看个人账户余额，消费记录
* ####     企业客户可以查看广告位定价，广告位状态
* ####     用户可以选择以管理员/企业客户身份登录
  

## 系统需求描述
 
* ####    用户可以浏览应用页，浏览到展示的广告
* ####    管理员可以终止广告展示
* ####    企业客户可以修改密码
    * 系统应该提供密码修改界面
    * 系统对操作鉴权
* ####    管理员可以查看企业信息，包括基本信息，账户信息，持有广告位，消费记录
* ####    管理员可以查看广告位信息，包括广告位状态，定价，所属企业，剩余时长等
* ####    管理员可以给广告位定价
    * 系统应该提供广告展示界面
    * 管理员可以通过展示页面给相应广告位置定价
* ####    企业用户可以预充值
    * 系统提供个人账户界面
    * 系统可以根据广告价格扣除账户里的金额
* ####    管理员可以审核企业客户注册信息，修改信息，提交的广告
    * 系统需要将待审核的用户存放起来供管理员查询
    * 管理员可以批准或者拒绝用户的注册或者信息修改
* ####    管理员可以增加新的管理员账号
    * 系统应该提供新建管理员的相关接口
    * 管理员可以指定新建管理员的相关信息
* ####     企业客户可以注销账号
    * 系统应该提供注销按钮
    * 对于注销按钮需要二次提醒
    * 系统应该对操作进行鉴权
* ####     企业客户可以查询审核进度
    * 系统应该为每个企业客户提供一个消息界面用于储存系统推送的消息
    * 系统应该提供审核进度在消息界面中
* ####     没有账号的企业客户可以注册账号，提供企业信息以供审核
    * 系统应该提供一个企业账号注册界面，允许没有账号的企业客户进行注册
    * 系统应该要求企业客户提供企业信息，包括公司名称、公司地址、负责人姓名、负责人电话，营业执照
    * 系统应该要求企业用户设置账号密码
    * 系统应该对密码加密存储
* ####     企业客户可以修改个人信息并提交审核
    * 系统应该提供个人信息展示界面
    * 系统应该支持的修改个人信息的信息种类：企业名称，企业地址，负责人，负责人电话，营业执照
    * 系统应该提供个人信息修改界面，允许用户随时修改个人信息
    * 系统应该对该操作鉴权，只有企业客户能修改
* ####     企业客户可以购买广告位，上传广告信息并提交审核
    * 系统应该提供广告购买界面
    * 系统应该提供广告信息上传功能，允许企业客户上传广告素材、目标链接，并发送到后台审核
    * 系统提供支付功能
    * 系统应该对该操作鉴权
    * 系统应该展示广告位的详细信息，包括定价，位置等
* ####     企业客户可以查看个人账户余额，消费记录
    * 系统应该提供查看个人账户的界面
    * 系统应该记录每个用户的个人账户余额，消费记录的信息
    * 系统应该对操作进行鉴权
    * 系统应该实时更新账户余额，消费记录
* ####     企业客户可以查看广告位定价，广告位状态
    * 系统应该提供一个展示广告位信息的界面供企业用户查看
    * 系统应该提供广告位的详细信息，包括状态，定价，到期时间等
    * 系统应该实时更新广告位状态，避免客户查看到过时信息
    * 系统应该对查看操作进行鉴权，只有企业客户和管理员用户能获取到信息
* ####     用户可以选择以管理员/企业客户身份登录
    * 系统应该提供一个首页供用户选择登录方式
    * 系统应该提供管理员登录页面，用表单形式提交用户登录信息
    * 系统应该提供企业客户登录界面，用表单形式提交用户登录信息
    * 系统应该对管理员和企业客户分开管理权限
