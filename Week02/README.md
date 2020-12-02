##学习笔记

问题解答：

Q:我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A:在Dao层中应该Wrap掉这个error,因为Dao层已经属于业务底层，再往下就是不受控制的基础框架层，所以在Dao层应该带上堆栈信息。
Dao层之上理论上不应带任何error stack信息，而基础框架层尽量也带原始error信息，避免双份的stack信息。

代码说明：

```javascript
├── README.md
├── api
│   └── authcontroller.go 认证有关API代码，对外暴露
├── dao
│   ├── db
│   │   └── mock.go 模拟数据库操作
│   └── user.go 用户DAO代码
├── go.mod
├── go.sum
├── main.go Http服务器启动代码
└── service
    └── userservice.go 用户Service代码
```



先启动服务端代码

```
go run main.go
```



然后客户端请求

```
http http://localhost:1234/users\?uid\=1
```



客户端返回如下,这个是正常返回信息

![](https://chenxf.org/usr/uploads/2020/12/3084231056.png)

再次请求触发错误信息

```
http http://localhost:1234/users\?uid\=4
```

客户端返回如下，这个是服务端捕获到异常的返回

![](https://chenxf.org/usr/uploads/2020/12/1449773572.png)

再看下服务端打印的堆栈信息

![image-20201202104706098](/Users/nbboy/Library/Application Support/typora-user-images/image-20201202104706098.png)

实际项目中可以把这些错误信息记录到log中，然后统一汇聚到ES中进行结构化和检索。