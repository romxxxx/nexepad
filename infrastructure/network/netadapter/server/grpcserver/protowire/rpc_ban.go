package protowire

import (
	"github.com/nexepanet/nexepad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *nexepadMessage_BanRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_BanRequest is nil")
	}
	return x.BanRequest.toAppMessage()
}

func (x *BanRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "BanRequestMessage is nil")
	}
	return &appmessage.BanRequestMessage{
		IP: x.Ip,
	}, nil
}

func (x *nexepadMessage_BanRequest) fromAppMessage(message *appmessage.BanRequestMessage) error {
	x.BanRequest = &BanRequestMessage{Ip: message.IP}
	return nil
}

func (x *nexepadMessage_BanResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_BanResponse is nil")
	}
	return x.BanResponse.toAppMessage()
}

func (x *BanResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "BanResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.BanResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *nexepadMessage_BanResponse) fromAppMessage(message *appmessage.BanResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.BanResponse = &BanResponseMessage{
		Error: err,
	}
	return nil
}
