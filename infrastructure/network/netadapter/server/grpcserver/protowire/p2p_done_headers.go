package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_DoneHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_DoneHeaders is nil")
	}
	return &appmessage.MsgDoneHeaders{}, nil
}

func (x *nexepadMessage_DoneHeaders) fromAppMessage(_ *appmessage.MsgDoneHeaders) error {
	return nil
}
