package db

//IconID holds all Iconinformation from eve
type IconID struct {
	Backgrounds []string `json:"backgrounds,omitempty" yaml:"backgrounds"`
	Description string   `json:"description,omitempty" yaml:"description"`
	Foregrounds []string `json:"foregrounds,omitempty" yaml:"foregrounds"`
	IconFile    string   `json:"iconFile,omitempty" yaml:"iconFile"`
	Obsolete    bool     `json:"obsolete,omitempty" yaml:"obsolete"`
}
