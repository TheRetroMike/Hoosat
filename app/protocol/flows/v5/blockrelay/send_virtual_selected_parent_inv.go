package blockrelay

import (
	"github.com/Hoosat-Oy/HTND/app/appmessage"
	peerpkg "github.com/Hoosat-Oy/HTND/app/protocol/peer"
	"github.com/Hoosat-Oy/HTND/app/protocol/protocolerrors"
	"github.com/Hoosat-Oy/HTND/domain"
	"github.com/Hoosat-Oy/HTND/domain/consensus/utils/constants"
	"github.com/Hoosat-Oy/HTND/infrastructure/config"
	"github.com/Hoosat-Oy/HTND/infrastructure/network/netadapter/router"
	"github.com/pkg/errors"
)

// SendVirtualSelectedParentInvContext is the interface for the context needed for the SendVirtualSelectedParentInv flow.
type SendVirtualSelectedParentInvContext interface {
	Domain() domain.Domain
	Config() *config.Config
}

// SendVirtualSelectedParentInv sends a peer the selected parent hash of the virtual
func SendVirtualSelectedParentInv(context SendVirtualSelectedParentInvContext,
	outgoingRoute *router.Route, peer *peerpkg.Peer) error {

	virtualSelectedParent, err := context.Domain().Consensus().GetVirtualSelectedParent()
	if err != nil {
		return err
	}
	if virtualSelectedParent.Equal(context.Config().NetParams().GenesisHash) {
		log.Debugf("Skipping sending the virtual selected parent hash to peer %s because it's the genesis.", peer)
		return nil
	}
	block, found, err := context.Domain().Consensus().GetBlock(virtualSelectedParent)
	if block.Header.Version() >= constants.PoWIntegrityMinVersion {
		if err != nil {
			return errors.Wrapf(err, "unable to fetch requested virtual selected parent block %s", virtualSelectedParent)
		}
		if !found {
			return protocolerrors.Errorf(false, "virtual selected parent block %s not found", virtualSelectedParent)
		}
		if block.PoWHash == "" {
			log.Debugf("Skipping sending the virtual selected parent hash to peer %s because it's PoW hash is empty.", peer)
			return nil
		}
	}

	log.Debugf("Sending virtual selected parent hash %s to peer %s", virtualSelectedParent, peer)
	virtualSelectedParentInv := appmessage.NewMsgInvBlock(virtualSelectedParent)
	return outgoingRoute.Enqueue(virtualSelectedParentInv)
}
