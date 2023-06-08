# 部署运行

## 方法一

1. 运行mysql容器
```
docker run -itd --name ji_mysql -p 3306:3306  -v /var/lib/mysql:/var/lib/mysql  -e MYSQL_ROOT_PASSWORD=123456 mysql:8.0
```
2. 运行服务
```
cd JI/cmd
./main
```

## 方法二

```
cd ji
docker-compose up
```

# 常见问题

1. mysql容器无法登陆问题
```
参考链接
https://blog.csdn.net/piaomiao_/article/details/119241127
docker exec -it ji_mysql bash
mysql -uroot -p123456
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
FLUSH PRIVILEGES;
```


# swag
```
https://blog.csdn.net/joychenwenyu/article/details/126935706

cd ji #进入项目根路径
swag init -g cmd/main.go -o docs  #初始化文档
```
docker restart ji_redis ji_mysql ji_kibana ji_es ji_rabbitmq