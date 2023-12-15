package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_IbdBlockLocatorHighestHashNotFound) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_IbdBlockLocatorHighestHashNotFound is nil")
	}
	return &appmessage.MsgIBDBlockLocatorHighestHashNotFound{}, nil
}

func (x *nexepadMessage_IbdBlockLocatorHighestHashNotFound) fromAppMessage(message *appmessage.MsgIBDBlockLocatorHighestHashNotFound) error {
	x.IbdBlockLocatorHighestHashNotFound = &IbdBlockLocatorHighestHashNotFoundMessage{}
	return nil
}
