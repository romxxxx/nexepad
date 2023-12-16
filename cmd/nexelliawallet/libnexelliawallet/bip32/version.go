package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// NexelliaMainnetPrivate is the version that is used for
// Nexellia mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var NexelliaMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// NexelliaMainnetPublic is the version that is used for
// Nexellia mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var NexelliaMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// NexelliaTestnetPrivate is the version that is used for
// Nexellia testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var NexelliaTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// NexelliaTestnetPublic is the version that is used for
// Nexellia testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var NexelliaTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// NexelliaDevnetPrivate is the version that is used for
// Nexellia devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var NexelliaDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// NexelliaDevnetPublic is the version that is used for
// Nexellia devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var NexelliaDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// NexelliaSimnetPrivate is the version that is used for
// Nexellia simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var NexelliaSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// NexelliaSimnetPublic is the version that is used for
// Nexellia simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var NexelliaSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case NexelliaMainnetPrivate:
		return NexelliaMainnetPublic, nil
	case NexelliaTestnetPrivate:
		return NexelliaTestnetPublic, nil
	case NexelliaDevnetPrivate:
		return NexelliaDevnetPublic, nil
	case NexelliaSimnetPrivate:
		return NexelliaSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case NexelliaMainnetPrivate:
		return true
	case NexelliaTestnetPrivate:
		return true
	case NexelliaDevnetPrivate:
		return true
	case NexelliaSimnetPrivate:
		return true
	}

	return false
}
