如何编写一个微服务？

这里用的是go的微服务框架go micro，具体的情况可以查阅：http://btfak.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/2016/03/28/go-micro/

开发微服务的步骤：

1、书写proto文件，定义函数等 ,即user.proto。

2、采用Google开发的protobuf协议编译.proto文件，生成.go文件：protoc --micro_out=. --go_out=. proto/filename.proto。

3、书写一个handler实现user.pb.go定义的接口。

	（1）先创建一个user.go文件。

	（2）定义一个结构体，命名为UserHandler，实现user.proto文件所有定义的service。

4、将handler注册进微服务，这一步在main中实现，即main.go。

5、通过go run main.go启动服务即可。

6、客户端访问服务，需要写一个web服务，将客户端的请求进行转发，（此步骤暂时没看明白，待续。）模板奉上（web.go），有待理解。

![image](https://github.com/WatermelonIsMe/GO_Micro_demo/blob/master/image/20181109142632.jpg)