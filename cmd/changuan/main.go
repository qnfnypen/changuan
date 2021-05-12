package main

func main() {
	// cmd exec
	cmdExec()
	// http 服务端
	go generateHTTPServer()
	// rpc 服务端
}
