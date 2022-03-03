package main

import (
	"regexp"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/rubblelabs/ripple/websockets"
)

var rAddressReg = "^r[1-9a-km-zA-HJ-NP-Z]{32,33}$"

// IsValidAddress check address
func IsValidAddress(addr string) bool {
	match, _ := regexp.MatchString(rAddressReg, addr)
	return match
}

// Account interface
type Account interface {
	String() string
	GetBalance(*websockets.Remote) float64
	CheckAddress(*websockets.Remote) (map[string]interface{}, error)
}

// AccountBase struct
type AccountBase struct{}

// GetBalance returns address balance
func (a *AccountBase) GetBalance(r *websockets.Remote) float64 {
	return 0
}

// AddressStatus type
type AddressStatus string

const (
	active   AddressStatus = "active"
	inactive AddressStatus = "inactive"
)

// CheckAddress check address format and status
func (a *AccountBase) CheckAddress(r *websockets.Remote) (AddressStatus, error) {
	return inactive, nil
}

// ActiveAccount struct
type ActiveAccount struct {
	AccountBase
	crypto.Key
	keyseq uint
}

func (a ActiveAccount) String() string {
	prefix := []byte{0}
	sequence := new(uint32)
	*sequence = uint32(a.keyseq)
	k := a.Key
	return crypto.Base58Encode(append(prefix, k.Id(sequence)...), crypto.ALPHABET)
}

// ObservingAccount struct
type ObservingAccount struct {
	AccountBase
	address string
	pubkey  string
}

func (a ObservingAccount) String() string {
	return a.address
}
