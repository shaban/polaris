package db

//CategoryID hols a categoryID
type CategoryID struct {
	IconID    int        `json:"iconID,omitempty" yaml:"iconID"`
	Name      translated `json:"name,omitempty" yaml:"name"`
	Published bool       `json:"published,omitempty" yaml:"published"`
}