package main

import "fmt"

func main() {
	s := NewServer("server0")
	fmt.Println(s.Echo("hi choco!"))
	s_sdk, ok := s.(*sdkServer)
	if !ok {
		fmt.Println("Type assertion failed")
		return
	}
	s_sdk.Name = "server0_change_name"
	fmt.Println(s.Echo("hi choco!"))
	
	e := NewServer1("server1")
	fmt.Println(e.Echo("hi xie!"))
	e_sdk, ok := e.(sdkServer)
	if !ok {
		fmt.Println("Type assertion failed")
		return
	}
	e_sdk.Name = "server1_change_name"
	fmt.Println(e.Echo("hi xie!"))
}

func NewServer(name string) Server {
    return &sdkServer{Name: name};
}

func NewServer1(name string) Server {
    return sdkServer{Name: name};
}

type Server interface {
    Echo(s string) string
}

type sdkServer struct {
	Name string
}

func (sdk sdkServer) Echo (str string) string {
    return sdk.Name + " echo: " +str
}
