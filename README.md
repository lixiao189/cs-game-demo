# A client / server game demo

我的一个用于测试服务器和客户端通信的测试项目, 全程采用 go 编写

## 如何运行 🚀

**客户端**

```bash
make run-client  # start up first client
make run-client1 # start up another client
```

**服务端**

```bash
make run-server
```

## 前后端交互

为了简单，这里使用 Json 交互

```json
{
  "type": "",
  "data": {}
}
```
