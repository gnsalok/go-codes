// Mutex Grouping

package main

import (
	"net"
	"sync"
)

type Server struct {
	listenAdd string
	isRunning bool
	/*
		-  Always put Mutex on the top of the type you protect
		- Its good to name like <type>Mu or peersLock
	*/
	peersMu sync.RWMutex
	peers   map[string]net.Conn
}

func main() {

}
