package renderer

import "github.com/vikramyadav1/weaver/parsers"

//go:generate mockery --name=Renderings

type Renderings interface {
	UpMigration() []byte
	DownMigration() []byte
	Model() []byte
	Server() []byte
	Repository() []byte
	Main() []byte
	PartialMain() []byte
}

type ResourceRenderer interface {
	Render(rd parsers.ResourceDefinition) Renderings
}
