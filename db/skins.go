package db

//SkinIDs holds all ship skin informations
type Skin struct {
	AllowCCPDevs       bool   `json:"allowCCPDevs,omitempty" yaml:"allowCCPDevs"`
	InternalName       string `json:"internalName,omitempty" yaml:"internalName"`
	SkinID             int    `json:"skinID,omitempty" yaml:"skinID"`
	SkinMaterialID     int    `json:"skinMaterialID,omitempty" yaml:"skinMaterialID"`
	Types              []int  `json:"types,omitempty" yaml:"types"`
	VisibleSerenity    bool   `json:"visibleSerenity,omitempty" yaml:"visibleSerenity"`
	VisibleTranquility bool   `json:"visibleTranquility,omitempty" yaml:"visibleTranquility"`
	SkinDescription string `json:"skinDescription,omitempty" yaml:"skinDescription"`
	IsStructureSkin bool `json:"isStructureSkin,omitempty" yaml:"isStructureSkin"`
}