package db

//Materials is a Blueprint Material
type Materials struct {
	Quantity int `json:"quantity,omitempty" yaml:"quantity"`
	TypeID   int `json:"typeID,omitempty" yaml:"typeID"`
	Probability float64 `json:"probability,omitempty" yaml:"probability"`
}

//Skills is a Blueprint Process Requirement
type Skills struct {
	Level  int `json:"level,omitempty" yaml:"level"`
	TypeID int `json:"typeID,omitempty" yaml:"typeID"`
}

//Process is a Blueprint Activities Process
type Process struct {
	Materials []Materials `json:"materials,omitempty" yaml:"materials"`
	Products  []Materials `json:"products,omitempty" yaml:"products"`
	Skills    []Skills    `json:"skills,omitempty" yaml:"skills"`
	Time      int         `json:"time,omitempty" yaml:"time"`
}

//Activities is a Blueprint Activity
type Activities struct {
	Reaction         Process `json:"reaction,omitempty" yaml:"reaction"`
	Copying          Process `json:"copying,omitempty" yaml:"copying"`
	Invention        Process `json:"invention,omitempty" yaml:"invention"`
	Manufacturing    Process `json:"manufacturing,omitempty" yaml:"manufacturing"`
	ResearchMaterial Process `json:"research_material,omitempty" yaml:"research_material"`
	ResearchTime     Process `json:"research_time,omitempty" yaml:"research_time"`
}

//Blueprint holds blueprint information
type Blueprint struct {
	Activities         `json:"activities,omitempty" yaml:"activities"`
	BlueprintTypeID    int `json:"blueprintTypeID,omitempty" yaml:"blueprintTypeID"`
	MaxProductionLimit int `json:"maxProductionLimit,omitempty" yaml:"maxProductionLimit"`
}