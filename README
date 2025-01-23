
## 用法

### 用法一： Request/Response

server

```go
s := wkrpc.New("tcp://127.0.0.1:10000")
s.Route("/hi", func(c *wkrpc.Context) {
	c.Write([]byte("reply"))
})
defer s.Stop()

_ = s.Start()

```

client 

```go

cli := client.New("tcp://127.0.0.1:10000")

defer cli.Stop()

_ = cli.Start()

resp,_ := cli.Request("/hi", []byte("hi"))

fmt.Println(string(resp.Body)) // reply

```

### 用法二 Send/Receive

server

```go
s := wkrpc.New("tcp://127.0.0.1:10000")
cli.OnMessage(func(conn gnet.Conn, m *proto.Message) {
    fmt.Println(string(msg.Body)) // hi
})
defer s.Stop()

_ = s.Start()

```

client

```go

cli := client.New("tcp://127.0.0.1:10000")

defer cli.Stop()

_ = cli.Start()


err = cli.Send(&proto.Message{
		MsgType: 1,
		Content: []byte("hi"),
})

```