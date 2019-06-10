package db

//CertificateLevel holds the ranks needed
//to fulfill a certificates respective level
type CertificateLevel struct {
	Advanced int `json:"advanced" yaml:"advanced"`
	Basic    int `json:"basic" yaml:"basic"`
	Elite    int `json:"elite" yaml:"elite"`
	Improved int `json:"improved" yaml:"improved"`
	Standard int `json:"standard" yaml:"standard"`
}
//Certificate holds all information
//on a given eve certificate
type Certificate struct {
	Description    string                    `json:"description" yaml:"description"`
	GroupID        int                       `json:"groupID" yaml:"groupID"`
	Name           string                    `json:"name" yaml:"name"`
	RecommendedFor []int                     `json:"recommendedFor" yaml:"recommendedFor"`
	SkillTypes     map[int]*CertificateLevel `json:"skillTypes" yaml:"skillTypes"`
}