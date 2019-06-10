package db

//GraphicID holds the graphic descriptions from Eve
type GraphicID struct {
	Description    string   `json:"description,omitempty" yaml:"description"`
	GraphicFile    string   `json:"graphicFile,omitempty" yaml:"graphicFile"`
	IconInfo       iconInfo `json:"iconInfo,omitempty" yaml:"iconInfo"`
	SofFactionName string   `json:"sofFactionName,omitempty" yaml:"sofFactionName"`
	SofHullName    string   `json:"sofHullName,omitempty" yaml:"sofHullName"`
	SofRaceName    string   `json:"sofRaceName,omitempty" yaml:"sofRaceName"`
}
type iconInfo struct {
	Backgrounds []string `json:"backgrounds,omitempty" yaml:"backgrounds"`
	Folder      string   `json:"folder,omitempty" yaml:"folder"`
	Foregrounds []string `json:"foregrounds,omitempty" yaml:"foregrounds"`
}
