
## 用法

### 服务器初始化

```go
s := wkrpc.New("tcp://127.0.0.1:10000")
defer s.Stop()

err = s.Start()
```

### 客户端初始化

```go
cli := client.New("tcp://127.0.0.1:10000")
defer cli.Stop()

err = cli.Start()
```

### 用法一: Request/Response

server

```go
s.Route("/hi", func(c *wkrpc.Context) {
	c.Write([]byte("reply"))
})
```

client 

```go
resp,_ := cli.Request("/hi", []byte("hi"))

fmt.Println(string(resp.Body)) // reply
```

### 用法二: Send/Receive

server

```go
s.OnMessage(func(conn gnet.Conn, m *proto.Message) {
    fmt.Println(string(msg.Body)) // hi
})
```

client

```go
err = cli.Send(&proto.Message{
	MsgType: 1,
	Content: []byte("hi"),
})
```