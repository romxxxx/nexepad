package protowire

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
)

func (x *nexepadMessage_PruningPoints) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "nexepadMessage_PruningPoints is nil")
	}

	if x.PruningPoints == nil {
		return nil, errors.Wrapf(errorNil, "x.PruningPoints is nil")
	}

	blockHeaders := make([]*appmessage.MsgBlockHeader, len(x.PruningPoints.Headers))
	for i, blockHeader := range x.PruningPoints.Headers {
		var err error
		blockHeaders[i], err = blockHeader.toAppMessage()
		if err != nil {
			return nil, err
		}
	}
	return &appmessage.MsgPruningPoints{
		Headers: blockHeaders,
	}, nil
}

func (x *nexepadMessage_PruningPoints) fromAppMessage(msgPruningPoints *appmessage.MsgPruningPoints) error {
	blockHeaders := make([]*BlockHeader, len(msgPruningPoints.Headers))
	for i, blockHeader := range msgPruningPoints.Headers {
		blockHeaders[i] = &BlockHeader{}
		err := blockHeaders[i].fromAppMessage(blockHeader)
		if err != nil {
			return err
		}
	}

	x.PruningPoints = &PruningPointsMessage{
		Headers: blockHeaders,
	}
	return nil
}
