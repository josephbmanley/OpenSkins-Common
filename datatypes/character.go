package datatypes

// Character represents something that can be skinned
type Character struct {
	Name string `json:"name"`
	Skin Skin   `json:"skin"`
}
