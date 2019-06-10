package db

//GroupID holds all Item Groups of Eve
type GroupID struct {
	Anchorable           bool       `json:"anchorable,omitempty" yaml:"anchorable"`
	Anchored             bool       `json:"anchored,omitempty" yaml:"anchored"`
	CategoryID           int        `json:"categoryID,omitempty" yaml:"categoryID"`
	FittableNonSingleton bool       `json:"fittableNonSingleton,omitempty" yaml:"fittableNonSingleton"`
	Name                 translated `json:"name,omitempty" yaml:"name"`
	Published            bool       `json:"published,omitempty" yaml:"published"`
	UseBasePrice         bool       `json:"useBasePrice,omitempty" yaml:"useBasePrice"`
	IconID int	`json:"iconID,omitempty" yaml:"iconID"`
}
