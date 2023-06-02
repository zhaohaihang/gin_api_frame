# Gin_Api_Frame
这是基于Gin实现的web脚手架，基于此脚手架可以快速的搭建一个单体应用。它集成了swagger、cron、logrus、wire、oss、jwt等常用的组件或中间件。

# 特性
1. Swagger 接口文档生成
2. err code 统一定义错误码
3. gorm 数据库组件
4. logrus 日志组件
5. redigo  redis组件
6. 支持 RESTful API 返回值规范
7. 支持 cron 定时任务
8. 支持JWT认证
9. 支持令牌桶限流
10. 集成七牛云对象存储
11. 集成wire实现依赖注入
12. 支持docker-compose部署
13. 引入了自定义参数验证器
14. 集成了email发送
15. 使用logrus替换gorm日志组件，使其日志输出在文件中

后续...
1. ES 用于全文检索

# 项目目录结构
```
cmd  -------        可执行程序的入口
config  -------     项目配置文件
docs  -----         项目swagger接口文档
internal 
  api  ---          restful的接口
  app  ---          总体的服务集合
  cron ---          定时任务
  dao ---           数据库操作
  http  ---         http服务
  model ---         数据模型
  routes --         路由
  serializer  ---   返回值的序列化
  service  ---      业务逻辑
  vaild ---         请求参数验证器
pkg
  consts ---        全局的静态常量
  database ---      数据库配置，目前仅有mysql
  e ---             错误码
  logger ---        统一的日志处理
  mail  ---         email配置
  middleware ---    包含一些Gin中间件
  redis   --        redis配置
  storages  ---     对象存储
    qiniu  ---      七牛云对象存储
  utils ---         一些工具类
  

```

# 功能说明
由于每个人的需求不同，所以项目中仅提供一个user相关的示例。
目前，仅包含登录、更新、上传头像，修改密码等供功能。


# 部署方式
```
cd gin_api_frame
docker-compose up
```

# 参考
搭建此脚手架，参考了如下的开源项目，在此感谢各位大佬无私的分享。

https://github.com/harvardfly/community-blogger
https://github.com/mlogclub/bbs-go
https://github.com/qingwave/weave