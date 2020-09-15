package main

import (
	"fmt"
	"time"

	getty "github.com/AlexStocks/getty/transport"
)

func main() {
	server := getty.NewTCPServer(getty.WithLocalAddress(":8888"))
	server.RunEventLoop(func(session getty.Session) error {
		session.SetPkgHandler(echoPkgHandler)
		session.SetEventListener(newEchoMessageHandler())
		return nil
	})
	fmt.Println(server.EndPointType())
	fmt.Println(server.ID())
	fmt.Println(server.IsClosed())
	fmt.Println(server)

	time.Sleep(time.Second * 100)
}

var (
	echoPkgHandler = NewEchoPackageHandler()
)

type EchoPackageHandler struct{}

func NewEchoPackageHandler() *EchoPackageHandler {
	return &EchoPackageHandler{}
}

func (h *EchoPackageHandler) Read(ss getty.Session, data []byte) (interface{}, int, error) {
	fmt.Printf("read %s\n", data)
	return string(data), len(data), nil
}

func (h *EchoPackageHandler) Write(ss getty.Session, pkg interface{}) ([]byte, error) {
	str, _ := pkg.(string)
	return []byte(str), nil
}

type EchoMessageHandler struct {
}

func newEchoMessageHandler() *EchoMessageHandler {
	return new(EchoMessageHandler)
}

func (h *EchoMessageHandler) OnOpen(session getty.Session) error {
	fmt.Println("OnOpen")
	return nil
}

func (h *EchoMessageHandler) OnError(session getty.Session, err error) {
	fmt.Println("OnError")
}

func (h *EchoMessageHandler) OnClose(session getty.Session) {
	fmt.Println("OnClose")
}

func (h *EchoMessageHandler) OnMessage(session getty.Session, pkg interface{}) {
	fmt.Printf("OnMessage %v\n", pkg)
}

func (h *EchoMessageHandler) OnCron(session getty.Session) {
	fmt.Println("OnCron")
}
