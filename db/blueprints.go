package db

//Materials is a Blueprint Material
type Materials struct {
	Quantity int `json:"quantity" yaml:"quantity"`
	TypeID   int `json:"typeID" yaml:"typeID"`
	Probability float64 `json:"probability" yaml:"probability"`
}

//Skills is a Blueprint Process Requirement
type Skills struct {
	Level  int `json:"level" yaml:"level"`
	TypeID int `json:"typeID" yaml:"typeID"`
}

//Process is a Blueprint Activities Process
type Process struct {
	Materials []Materials `json:"materials" yaml:"materials"`
	Products  []Materials `json:"products" yaml:"products"`
	Skills    []Skills    `json:"skills" yaml:"skills"`
	Time      int         `json:"time" yaml:"time"`
}

//Activities is a Blueprint Activity
type Activities struct {
	Reaction         Process `json:"reaction" yaml:"reaction"`
	Copying          Process `json:"copying" yaml:"copying"`
	Invention        Process `json:"invention" yaml:"invention"`
	Manufacturing    Process `json:"manufacturing" yaml:"manufacturing"`
	ResearchMaterial Process `json:"research_material" yaml:"research_material"`
	ResearchTime     Process `json:"research_time" yaml:"research_time"`
}

//Blueprint holds blueprint information
type Blueprint struct {
	Activities         `json:"activities" yaml:"activities"`
	BlueprintTypeID    int `json:"blueprintTypeID" yaml:"blueprintTypeID"`
	MaxProductionLimit int `json:"maxProductionLimit" yaml:"maxProductionLimit"`
}