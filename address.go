package address

import (
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

var ErrUnknownKeyType = errors.New("Unknown key type")

func MakeAddress(masterPub string, change bool, index int) (string, error) {
	var segwit, testnet bool
	if strings.HasPrefix(masterPub, "xpub") {
		segwit = false
		testnet = false
	} else if strings.HasPrefix(masterPub, "zpub") {
		segwit = true
		testnet = false
	} else if strings.HasPrefix(masterPub, "tpub") {
		segwit = false
		testnet = true
	} else if strings.HasPrefix(masterPub, "vpub") {
		segwit = true
		testnet = true
	} else {
		return "", ErrUnknownKeyType
	}

	acc0, err := hdkeychain.NewKeyFromString(masterPub)
	if err != nil {
		return "", fmt.Errorf("hdkeychain.NewKeyFromString: %v", err)
	}

	chainIndex := 0
	if change {
		chainIndex = 1
	}

	chain, err := acc0.Child(uint32(chainIndex))
	if err != nil {
		return "", fmt.Errorf("acc0.Child: %v", err)
	}

	leaf, err := chain.Child(uint32(index))
	if err != nil {
		return "", fmt.Errorf("chain.Child: %v", err)
	}

	params := &chaincfg.MainNetParams
	if testnet {
		params = &chaincfg.TestNet3Params
	}

	if !segwit {
		addr, err := leaf.Address(params)
		if err != nil {
			return "", fmt.Errorf("leaf.Address: %v", err)
		}
		return addr.String(), nil
	}

	ecpub, err := leaf.ECPubKey()
	if err != nil {
		return "", fmt.Errorf("leaf.ECPubKey: %v", err)
	}

	addr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(ecpub.SerializeCompressed()), params)
	if err != nil {
		return "", fmt.Errorf("btcutil.NewAddressWitnessPubKeyHash: %v", err)
	}

	return addr.String(), nil
}
