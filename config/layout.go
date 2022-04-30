package config

type Layout int

const (
	LayoutYaml Layout = iota
	LayoutJson
	LayoutToml
	LayoutIni
)
