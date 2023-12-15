package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_StopNotifyingUtxosChangedRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_StopNotifyingUtxosChangedRequest is nil")
	}
	return x.StopNotifyingUtxosChangedRequest.toAppMessage()
}

func (x *nexepadMessage_StopNotifyingUtxosChangedRequest) fromAppMessage(message *appmessage.StopNotifyingUTXOsChangedRequestMessage) error {
	x.StopNotifyingUtxosChangedRequest = &StopNotifyingUtxosChangedRequestMessage{
		Addresses: message.Addresses,
	}
	return nil
}

func (x *StopNotifyingUtxosChangedRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "StopNotifyingUtxosChangedRequestMessage is nil")
	}
	return &appmessage.StopNotifyingUTXOsChangedRequestMessage{
		Addresses: x.Addresses,
	}, nil
}

func (x *nexepadMessage_StopNotifyingUtxosChangedResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_StopNotifyingUtxosChangedResponse is nil")
	}
	return x.StopNotifyingUtxosChangedResponse.toAppMessage()
}

func (x *nexepadMessage_StopNotifyingUtxosChangedResponse) fromAppMessage(message *appmessage.StopNotifyingUTXOsChangedResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.StopNotifyingUtxosChangedResponse = &StopNotifyingUtxosChangedResponseMessage{
		Error: err,
	}
	return nil
}

func (x *StopNotifyingUtxosChangedResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "StopNotifyingUtxosChangedResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.StopNotifyingUTXOsChangedResponseMessage{
		Error: rpcErr,
	}, nil
}
