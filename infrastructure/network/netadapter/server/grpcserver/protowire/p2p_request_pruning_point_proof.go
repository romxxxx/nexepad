package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_RequestPruningPointProof) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_RequestPruningPointProof is nil")
	}
	return &appmessage.MsgRequestPruningPointProof{}, nil
}

func (x *nexepadMessage_RequestPruningPointProof) fromAppMessage(_ *appmessage.MsgRequestPruningPointProof) error {
	return nil
}
