package db

//CertificateLevel holds the ranks needed
//to fulfill a certificates respective level
type CertificateLevel struct {
	Advanced int `json:"advanced,omitempty" yaml:"advanced"`
	Basic    int `json:"basic,omitempty" yaml:"basic"`
	Elite    int `json:"elite,omitempty" yaml:"elite"`
	Improved int `json:"improved,omitempty" yaml:"improved"`
	Standard int `json:"standard,omitempty" yaml:"standard"`
}

//Certificate holds all information
//on a given eve certificate
type Certificate struct {
	Description    string                    `json:"description,omitempty" yaml:"description"`
	GroupID        int                       `json:"groupID,omitempty" yaml:"groupID"`
	Name           string                    `json:"name,omitempty" yaml:"name"`
	RecommendedFor []int                     `json:"recommendedFor,omitempty" yaml:"recommendedFor"`
	SkillTypes     map[int]*CertificateLevel `json:"skillTypes,omitempty" yaml:"skillTypes"`
}
