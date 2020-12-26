package datatypes

// Skin struct holds the
type Skin struct {
	Name     string            `json:"name"`
	Location string            `json:"location"`
	Metadata map[string]string `json:"metadata"`
}

// Skins are an array of Skin
type Skins []Skin
