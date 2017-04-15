package cg

import (
	"encoding/json"
	"errors"
	"ipc"
	"sync"
)

var _ ipc.Server = &CenterServer{} //确认实现了Server接口

type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}
