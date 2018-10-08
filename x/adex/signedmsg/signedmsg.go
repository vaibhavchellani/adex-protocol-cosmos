package signedmsg

import (
	"github.com/tendermint/tendermint/crypto"
)

func IsSigned(pubkey crypto.PubKey, data []byte, sig []byte) bool {
	// This is abstracted away so that we can easily modify the logic here, in case we need custom signature types
	return pubkey.VerifyBytes(data, sig)
}
