package main

import (
	"fmt"
	"reflect"
	"strings"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.nexepadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.nexepadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.nexepadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.nexepadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.nexepadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.nexepadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.nexepadMessage_BanRequest{}),
	reflect.TypeOf(protowire.nexepadMessage_UnbanRequest{}),
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
