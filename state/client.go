package state

import (
	"github.com/coreos/etcd/clientv3"
)

type Client interface {
	// Notify the state server new events. Client MAY notify
	// state server duplicated events. It is server's responsibility
	// to de-dup.
	Notify(e Event) error
}

type client struct {
	clientv3.Client
}
