package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *nexepadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
