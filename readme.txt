1.curl
a.查找用户
curl -X GET 'http://localhost:8001/users/findUserById?id=58'
b.添加用户
curl -X POST 'http://localhost:8001/users/addUser' -d '{"name":"hongyuqin11","password":"sssddd"}'
c.更新用户
curl -X POST 'http://localhost:8001/users/updateUser' -d '{"id":19,"password":"ssss","tag":2}'
d.删除用户
curl -X GET 'http://localhost:8001/users/delUser?id=58'

2.测试流程
 go test

