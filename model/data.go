package model

import (
	"sync"

	"github.com/googollee/go-socket.io"
)

var mutex = sync.Mutex{}

type ClientData struct {
	Title string
	ID    string
}
type ServerData struct {
	Title string
	ID    string
}
type Orient struct {
	ID    string  `json:"id"`
	Gamma float64 `json:"gamma"`
	Beta  float64 `json:"beta"`
	Alpha float64 `json:"alpha"`
}

type Conn map[string]socketio.Socket

func (c Conn) Add(id string, socket socketio.Socket) {
	mutex.Lock()
	c[id] = socket
	mutex.Unlock()
}
func (c Conn) Remove(id string) {
	mutex.Lock()
	delete(c, id)
	mutex.Unlock()
}
