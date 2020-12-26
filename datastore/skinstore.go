package datastore

import "github.com/josephbmanley/OpenSkins-Common/datatypes"

type Skinstore interface {
	Initialize() error
	GetSkin(string) (datatypes.Skin, error)
	AddSkin(string, []byte) error
}
