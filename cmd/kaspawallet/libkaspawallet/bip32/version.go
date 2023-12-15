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

// nexepaMainnetPrivate is the version that is used for
// nexepa mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var nexepaMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// nexepaMainnetPublic is the version that is used for
// nexepa mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var nexepaMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// nexepaTestnetPrivate is the version that is used for
// nexepa testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var nexepaTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// nexepaTestnetPublic is the version that is used for
// nexepa testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var nexepaTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// nexepaDevnetPrivate is the version that is used for
// nexepa devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var nexepaDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// nexepaDevnetPublic is the version that is used for
// nexepa devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var nexepaDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// nexepaSimnetPrivate is the version that is used for
// nexepa simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var nexepaSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// nexepaSimnetPublic is the version that is used for
// nexepa simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var nexepaSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case nexepaMainnetPrivate:
		return nexepaMainnetPublic, nil
	case nexepaTestnetPrivate:
		return nexepaTestnetPublic, nil
	case nexepaDevnetPrivate:
		return nexepaDevnetPublic, nil
	case nexepaSimnetPrivate:
		return nexepaSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case nexepaMainnetPrivate:
		return true
	case nexepaTestnetPrivate:
		return true
	case nexepaDevnetPrivate:
		return true
	case nexepaSimnetPrivate:
		return true
	}

	return false
}
