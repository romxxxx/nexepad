package main

import (
	"context"
	"fmt"

	"github.com/romxxxx/nexepad/cmd/nexelliawallet/daemon/client"
	"github.com/romxxxx/nexepad/cmd/nexelliawallet/daemon/pb"
	"github.com/romxxxx/nexepad/cmd/nexelliawallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.Formatnexe(addressBalance.Available), utils.Formatnexe(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, nexe %s %s%s\n", utils.Formatnexe(response.Available), utils.Formatnexe(response.Pending), pendingSuffix)

	return nil
}
