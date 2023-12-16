package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *NexepadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NexepadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *NexepadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
