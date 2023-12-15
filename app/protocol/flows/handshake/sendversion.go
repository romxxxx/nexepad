package handshake

import (
	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/app/appmessage"
	"github.com/romxxxx/nexepad/app/protocol/common"
	peerpkg "github.com/romxxxx/nexepad/app/protocol/peer"
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/infrastructure/network/netadapter/router"
	"github.com/romxxxx/nexepad/version"
)

var (
	// userAgentName is the user agent name and is used to help identify
	// ourselves to other nexepa peers.
	userAgentName = "nexepad"

	// userAgentVersion is the user agent version and is used to help
	// identify ourselves to other nexepa peers.
	userAgentVersion = version.Version()

	// defaultServices describes the default services that are supported by
	// the server.
	defaultServices = appmessage.DefaultServices

	// defaultRequiredServices describes the default services that are
	// required to be supported by outbound peers.
	defaultRequiredServices = appmessage.SFNodeNetwork
)

type sendVersionFlow struct {
	HandleHandshakeContext

	incomingRoute, outgoingRoute *router.Route
	peer                         *peerpkg.Peer
}

// SendVersion sends a version to a peer and waits for verack.
func SendVersion(context HandleHandshakeContext, incomingRoute *router.Route,
	outgoingRoute *router.Route, peer *peerpkg.Peer) error {

	flow := &sendVersionFlow{
		HandleHandshakeContext: context,
		incomingRoute:          incomingRoute,
		outgoingRoute:          outgoingRoute,
		peer:                   peer,
	}
	return flow.start()
}

func (flow *sendVersionFlow) start() error {
	onEnd := logger.LogAndMeasureExecutionTime(log, "sendVersionFlow.start")
	defer onEnd()

	log.Debugf("Starting sendVersionFlow with %s", flow.peer.Address())

	// Version message.
	localAddress := flow.AddressManager().BestLocalAddress(flow.peer.Connection().NetAddress())
	subnetworkID := flow.Config().SubnetworkID
	if flow.Config().ProtocolVersion < minAcceptableProtocolVersion {
		return errors.Errorf("configured protocol version %d is obsolete", flow.Config().ProtocolVersion)
	}
	msg := appmessage.NewMsgVersion(localAddress, flow.NetAdapter().ID(),
		flow.Config().ActiveNetParams.Name, subnetworkID, flow.Config().ProtocolVersion)
	msg.AddUserAgent(userAgentName, userAgentVersion, flow.Config().UserAgentComments...)

	// Advertise the services flag
	msg.Services = defaultServices

	// Advertise our max supported protocol version.
	msg.ProtocolVersion = flow.Config().ProtocolVersion

	// Advertise if inv messages for transactions are desired.
	msg.DisableRelayTx = flow.Config().BlocksOnly

	err := flow.outgoingRoute.Enqueue(msg)
	if err != nil {
		return err
	}

	// Wait for verack
	log.Debugf("Waiting for verack")
	_, err = flow.incomingRoute.DequeueWithTimeout(common.DefaultTimeout)
	if err != nil {
		return err
	}
	log.Debugf("Got verack")
	return nil
}
