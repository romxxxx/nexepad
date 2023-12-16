package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/romxxxx/nexepad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.NexepadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.NexepadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.NexepadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.NexepadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.NexepadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.NexepadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.NexepadMessage_BanRequest{}),
	reflect.TypeOf(protowire.NexepadMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
