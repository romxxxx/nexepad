package mining

import (
	"math"
	"math/rand"

	"github.com/pkg/errors"
	"github.com/romxxxx/nexepad/domain/consensus/model/externalapi"
	"github.com/romxxxx/nexepad/domain/consensus/utils/pow"
)

// SolveBlock increments the given block's nonce until it matches the difficulty requirements in its bits field
func SolveBlock(block *externalapi.DomainBlock, rd *rand.Rand) {
	header := block.Header.ToMutable()
	state := pow.NewState(header)
	for state.Nonce = rd.Uint64(); state.Nonce < math.MaxUint64; state.Nonce++ {
		if state.CheckProofOfWork() {
			header.SetNonce(state.Nonce)
			block.Header = header.ToImmutable()
			return
		}
	}

	panic(errors.New("went over all the nonce space and couldn't find a single one that gives a valid block"))
}
