package main

import (
	"context"

	"github.com/ConsenSys/handel"
	"github.com/ConsenSys/handel/network"
	"github.com/ConsenSys/handel/simul/lib"
	"github.com/ConsenSys/handel/simul/p2p"
)

// MakeUDP is a p2p.Adaptor that returns a list of node using UDP as their
// network and that sends the packet to all identities.
func MakeUDP(ctx context.Context, list lib.NodeList, ids []int, threshold int, opts p2p.Opts) (handel.Registry, []p2p.Node) {
	created := len(ids)
	encoding := extractEncoding(opts)
	nodes := make([]p2p.Node, 0, created)
	for _, n := range list {
		if p2p.IsIncluded(ids, int(n.ID())) {
			udpNode := NewNode(n.SecretKey, n.Identity, list.Registry(), encoding)
			nodes = append(nodes, udpNode)
		}
	}
	//fmt.Println(list.Registry())
	return list.Registry(), nodes
}

func extractEncoding(opts p2p.Opts) network.Encoding {
	return network.NewGOBEncoding()
}
