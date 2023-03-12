package rpcserver

import (
	"bytes"
	"io"
	"net/rpc/jsonrpc"
)

// RpcRequest represents a RPC request.
// RpcRequest implements the io.ReadWriteCloser interface.
type RpcRequest struct {
	r    io.Reader     // holds the JSON formated RPC request
	rw   io.ReadWriter // holds the JSON formated RPC response
	done chan bool     // signals then end of the RPC request
}

// NewRpcRequest returns a new RpcRequest.
func NewRpcRequest(r io.Reader) *RpcRequest {
	var buf bytes.Buffer
	done := make(chan bool)
	return &RpcRequest{r, &buf, done}
}

// Read implements the io.ReadWriteCloser Read method.
func (r *RpcRequest) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

// Write implements the io.ReadWriteCloser Write method.
func (r *RpcRequest) Write(p []byte) (n int, err error) {
	return r.rw.Write(p)
}

// Close implements the io.ReadWriteCloser Close method.
func (r *RpcRequest) Close() error {
	r.done <- true
	return nil
}

// Call invokes the RPC request, waits for it to complete, and returns the results.
func (r *RpcRequest) Call() io.Reader {
	go jsonrpc.ServeConn(r)
	<-r.done
	return r.rw
}
