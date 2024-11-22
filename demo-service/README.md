## mysql

```
docker run --name mysql -e MYSQL_ROOT_PASSWORD=admin123 -e
 MYSQL_USER=admin -e MYSQL_PASSWORD=admin -e MYSQL_DATABASE=demodb -p 3306:3306 -d mysql:latest
 ```

## UI for mysql

```
docker run -d --name adminer -p 8089:8080 adminer
```

## To install mysql

```
https://dev.mysql.com/downloads/
```