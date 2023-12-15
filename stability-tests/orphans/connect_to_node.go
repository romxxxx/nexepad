package main

import (
	"fmt"
	"os"

	"github.com/romxxxx/nexepad/infrastructure/config"
	"github.com/romxxxx/nexepad/infrastructure/network/netadapter/standalone"
)

func connectToNode() *standalone.Routes {
	cfg := activeConfig()

	nexepadConfig := config.DefaultConfig()
	nexepadConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(nexepadConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating minimalNetAdapter: %+v", err)
		os.Exit(1)
	}
	routes, err := minimalNetAdapter.Connect(cfg.NodeP2PAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to node: %+v", err)
		os.Exit(1)
	}
	return routes
}
