package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/nexepanet/nexepad/stability-tests/common"
	"github.com/nexepanet/nexepad/stability-tests/common/rpc"
	"github.com/nexepanet/nexepad/util/profiling"
)

var timeout = 30 * time.Second

func main() {
	err := parseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config: %+v", err)
		os.Exit(1)
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())
	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	blocks, topBlock, err := prepareBlocks()
	if err != nil {
		log.Errorf("Error preparing blocks: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}

	routes := connectToNode()

	rpcClient, err := rpc.ConnectToRPC(&cfg.Config, cfg.NetParams())
	if err != nil {
		panic(errors.Wrap(err, "error connecting to JSON-RPC server"))
	}

	defer rpcClient.Disconnect()
	err = sendBlocks(routes, blocks, topBlock)
	if err != nil {
		backendLog.Close()
		log.Errorf("Error sending blocks: %+v", err)
		os.Exit(1)
	}

	// Wait a second to let nexepad process orphans
	<-time.After(1 * time.Second)

	err = checkTopBlockIsTip(rpcClient, topBlock)
	if err != nil {
		log.Errorf("Error in checkTopBlockIsTip: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}
}
