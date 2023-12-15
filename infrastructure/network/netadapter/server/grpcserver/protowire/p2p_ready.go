package protowire

import (
	"github.com/nexepanet/nexepad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *nexepadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *nexepadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
