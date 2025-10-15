package shadowsocks_2022

import (
	"github.com/luoluodaduan/xray-core/common/protocol"
	"google.golang.org/protobuf/proto"
)

// MemoryAccount is an account type converted from Account.
type MemoryAccount struct {
	Key string
}

// AsAccount implements protocol.AsAccount.
func (u *Account) AsAccount() (protocol.Account, error) {
	return &MemoryAccount{
		Key: u.GetKey(),
	}, nil
}

// Equals implements protocol.Account.Equals().
func (a *MemoryAccount) Equals(another protocol.Account) bool {
	if account, ok := another.(*MemoryAccount); ok {
		return a.Key == account.Key
	}
	return false
}

func (a *MemoryAccount) ToProto() proto.Message {
	return &Account{
		Key: a.Key,
	}
}
