// +build zmq

package tools

import (
	"time"

	"github.com/pebbe/zmq4"
)

// ZMQ stores a ZMQ4 Socket
type ZMQ struct {
	Socket *zmq4.Socket
}

// NewZMQ generates a new instance of ZMQ
func NewZMQ(socket *zmq4.Socket) *ZMQ {
	return &ZMQ{
		Socket: socket,
	}
}

// SendDirect - Takes a resource, method, and an object to send it to all live
// clients.
// To achieve this, it uses ZeroMQ.
func (zmq *ZMQ) SendDirect(url, res, method string, object interface{}) {
	msg := struct {
		Res    string      `json:"res"`
		Data   interface{} `json:"data,omitempty"`
		Method string      `json:"method"`
		Time   time.Time   `json:"time"`
	}{
		Res:    res,
		Data:   object,
		Method: method,
		Time:   time.Now(),
	}

	encoded := Stringify(msg)

	zmq.Socket.SendMessageDontwait(url, encoded)
}
