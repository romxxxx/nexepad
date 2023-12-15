package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_RequestNextHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_RequestNextHeaders is nil")
	}
	return &appmessage.MsgRequestNextHeaders{}, nil
}

func (x *nexepadMessage_RequestNextHeaders) fromAppMessage(_ *appmessage.MsgRequestNextHeaders) error {
	return nil
}
