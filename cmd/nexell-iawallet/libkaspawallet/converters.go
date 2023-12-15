package libnexepawallet

import (
	"encoding/hex"

	"github.com/romxxxx/nexepad/app/appmessage"
	"github.com/romxxxx/nexepad/cmd/nexepawallet/daemon/pb"
	"github.com/romxxxx/nexepad/domain/consensus/model/externalapi"
	"github.com/romxxxx/nexepad/domain/consensus/utils/transactionid"
	"github.com/romxxxx/nexepad/domain/consensus/utils/utxo"
)

// nexepawalletdUTXOsTolibnexepawalletUTXOs converts a  []*pb.UtxosByAddressesEntry to a []*libnexepawallet.UTXO
func nexepawalletdUTXOsTolibnexepawalletUTXOs(nexepawalletdUtxoEntires []*pb.UtxosByAddressesEntry) ([]*UTXO, error) {
	UTXOs := make([]*UTXO, len(nexepawalletdUtxoEntires))
	for i, entry := range nexepawalletdUtxoEntires {
		script, err := hex.DecodeString(entry.UtxoEntry.ScriptPublicKey.ScriptPublicKey)
		if err != nil {
			return nil, err
		}
		transactionID, err := transactionid.FromString(entry.Outpoint.TransactionId)
		if err != nil {
			return nil, err
		}
		UTXOs[i] = &UTXO{
			UTXOEntry: utxo.NewUTXOEntry(
				entry.UtxoEntry.Amount,
				&externalapi.ScriptPublicKey{
					Script:  script,
					Version: uint16(entry.UtxoEntry.ScriptPublicKey.Version),
				},
				entry.UtxoEntry.IsCoinbase,
				entry.UtxoEntry.BlockDaaScore,
			),
			Outpoint: &externalapi.DomainOutpoint{
				TransactionID: *transactionID,
				Index:         entry.Outpoint.Index,
			},
		}
	}
	return UTXOs, nil
}

// AppMessageUTXOTonexepawalletdUTXO converts an appmessage.UTXOsByAddressesEntry to a  pb.UtxosByAddressesEntry
func AppMessageUTXOTonexepawalletdUTXO(appUTXOsByAddressesEntry *appmessage.UTXOsByAddressesEntry) *pb.UtxosByAddressesEntry {
	return &pb.UtxosByAddressesEntry{
		Outpoint: &pb.Outpoint{
			TransactionId: appUTXOsByAddressesEntry.Outpoint.TransactionID,
			Index:         appUTXOsByAddressesEntry.Outpoint.Index,
		},
		UtxoEntry: &pb.UtxoEntry{
			Amount: appUTXOsByAddressesEntry.UTXOEntry.Amount,
			ScriptPublicKey: &pb.ScriptPublicKey{
				Version:         uint32(appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Version),
				ScriptPublicKey: appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Script,
			},
			BlockDaaScore: appUTXOsByAddressesEntry.UTXOEntry.BlockDAAScore,
			IsCoinbase:    appUTXOsByAddressesEntry.UTXOEntry.IsCoinbase,
		},
	}
}
