package datastore

import "github.com/josephbmanley/OpenSkins-Common/datatypes"

type Userstore interface {
	Initialize() error
	GetUser(uid string) (datatypes.User, error)
	SetUser(datatypes.User) error
}
