How to write a micro service？

This is the go microservice framework "go micro",the detail you can go:http://btfak.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/2016/03/28/go-micro/

The process to develop a micro service is etc：

First：write the file of proto, define the function and etc,that is user.proto.

Second : use the protobuf protocol which develop by google compile the file .proto,then you can get the .go file :protoc --micro_out=. --go_out=. proto/filename.proto.

Third : write a handler to achieve the user.pb.go interface implement.
	
	(1)set up a file of user.go

	(2)define a struct named UserHandler to achieve the service of user.proto which defined.

Fourth : regist the handler to the micro service, it's success in the main.go.

Fifth : start the service with go run main.go.

Sixth : we need write a web service to realize client visit sercice,which transmit the request of client.(web.go).this step i can't understand it,it's  continued.

![image](http://192.168.207.25/WatermelonIsMe/GO_Micro_demo/raw/master/image/20181109142632.jpg)



