1.curl
a.查找用户
curl -X GET 'http://localhost:8001/users/findById?id=58'
b.添加用户
curl -X POST 'http://localhost:8001/users/addUser' -d '{"name":"hongyuqin11","password":"sssddd"}'
c.更新用户
curl -X POST 'http://localhost:8001/users/updateUser' -d '{"id":19,"password":"ssss","tag":2}'
d.删除用户
curl -X GET 'http://localhost:8001/users/delUser?id=58'

2.测试用例说明
  测试用例测试了增删改查4个http接口的逻辑，所以跑测试用例时，需要先开启main.go,启动http服务。

2.测试流程
a.启动main.go
b.运行main_test.go的测试用例
  go test -count=1 -v  main_test.go main.go

