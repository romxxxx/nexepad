package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_DoneBlocksWithTrustedData) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_DoneBlocksWithTrustedData is nil")
	}
	return &appmessage.MsgDoneBlocksWithTrustedData{}, nil
}

func (x *nexepadMessage_DoneBlocksWithTrustedData) fromAppMessage(_ *appmessage.MsgDoneBlocksWithTrustedData) error {
	return nil
}
