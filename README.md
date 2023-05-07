运行mysql容器
```
docker run -itd --name ji_mysql -p 3306:3306  -v /var/lib/mysql:/var/lib/mysql   -e MYSQL_ROOT_PASSWORD=123456 mysql:8.0
```
处理mysql登陆问题
```
docker exec -it ji_mysql bash
mysql -uroot -p123456
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
FLUSH PRIVILEGES;
```


参考链接
https://blog.csdn.net/piaomiao_/article/details/119241127