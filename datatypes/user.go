package datatypes

// User is an object representing a user
type User struct {
	UID        string `json:"uid"`
	OwnedSkins Skins  `json:"skins"`
}
