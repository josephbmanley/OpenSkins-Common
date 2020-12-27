package datatypes

// Character represents something that can be skinned
type Character struct {
	UID  string `json:"uid"`
	Skin Skin   `json:"skin"`
}

// Characters are an array of Character
type Characters []Character
