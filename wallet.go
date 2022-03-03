package main

import (
	"fmt"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/rubblelabs/ripple/websockets"
)

var (
	myWallet *Wallet
)

// Wallet struct
type Wallet struct {
	accounts          map[string]Account
	observingAccounts map[string]Account
	endpoint          string
}

func (w *Wallet) remote() *websockets.Remote {
	return nil
}

// GetWallet returns myWallet
func GetWallet() *Wallet {
	if myWallet == nil {
		myWallet = new(Wallet)
		myWallet.Init()
	}
	return myWallet
}

// Init initailizes myWallet
func (w *Wallet) Init() {
	w.accounts = make(map[string]Account)
	// TODO remote
}

// SetRemote sets up websocket connection
func (w *Wallet) SetRemote(endpoint string) error {
	return nil
}

// ImportAccountFromSeed imports private key ripple key seed and key sequence
func (w *Wallet) ImportAccountFromSeed(seed string, keyseq uint) (string, error) {
	key, err := ImportKeyFromSeed(seed, "ecdsa")
	if err != nil {
		return "", fmt.Errorf("Cannot import account for given seed: %v", err.Error())
	}
	acct := Account{
		Key:    key,
		keyseq: keyseq,
	}
	myWallet.accounts[acct.String()] = acct
	return acct.String(), nil
}

// ImportAccountFromEthPrivate imports ethereum style ecdsa private key
func (w *Wallet) ImportAccountFromEthPrivate(privkeyHex string) (string, error) {
	key, err := ImportKeyFromSeed(seed, "ecdsa")
	if err != nil {
		return "", fmt.Errorf("Cannot import account for given seed: %v", err.Error())
	}
	acct := Account{
		Key:    key,
		keyseq: keyseq,
	}
	myWallet.accounts[acct.String()] = acct
	return acct.String(), nil
}

// ImportObservingAccountFromAddress imports account from ripple address
func (w *Wallet) ImportObservingAccountFromAddress(address string) error {}

// ImportObservingAccountFromEthPubkey imports observing account from ethereum public key
func (w *Wallet) ImportObservingAccountFromEthPubkey(address string) error {}

// RemoveAccount removes account from wallet
func (w *Wallet) RemoveAccount(address string) error {
	return nil
}

// ListAccounts returns a slice of wallet accounts
func (w *Wallet) ListAccounts() []interface{} {
	res := make([]interface{}, len(w.accounts)+len(w.observingAccounts))
	for a := range w.accounts {
		res = append(res, a)
	}
	for a := range w.observingAccounts {
		res = append(res, a)
	}
	return res
}

// BuildPayment returns a payment transaction data structure
func (w *Wallet) BuildPayment(from, to, string, amount float64, memo string) (map[string]interface{}, error) {
	return nil, nil
}

//CommitPayment signs and pushes payment transaction
func (w *Wallet) CommitPayment(txJSON string) (map[string]interface{}, error) {
	return nil, nil
}

// CheckTx returns transaction status
func (w *Wallet) CheckTx(txhash string) (map[string]interface{}, error) {
	return nil, nil
}

// ImportKeyFromSeed converts seed to ripple key
func ImportKeyFromSeed(seed string, cryptoType string) (crypto.Key, error) {
	shash, err := crypto.NewRippleHashCheck(seed, crypto.RIPPLE_FAMILY_SEED)
	if err != nil {
		return nil, err
	}
	switch cryptoType {
	case "ed25519":
		key, err := crypto.NewEd25519Key(shash.Payload())
		return key, err
	case "ecdsa":
		key, err := crypto.NewECDSAKey(shash.Payload())
		return key, err
	default:
		return nil, err
	}
}
