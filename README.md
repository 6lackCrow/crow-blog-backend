# crow-blog-backend

crow-blog-backend是crow-blog-vue3的后端  

启动时在src同级目录下创建application.yml
```yaml
server:
  port: 端口号

data-source:
  mysql-dsn: mysql连接url

# 添加sys-user配置实现启动时创建用户(表中已有用户时不会执行创建,注意需要所有配置均符合校验规则才会创建)
sys-user:
  account: 账号 #以大小写字母开头，由大小写字母，数字和下划线组成的长度为4到15的字符串
  password: 密码 #以大小写字母开头，由大小写字母，数字组成的长度为5到17的字符串
  email: 邮箱
  nickname: 昵称
```