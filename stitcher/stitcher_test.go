package stitcher

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/vikramyadav1/weaver/parsers"
	"github.com/vikramyadav1/weaver/renderer/mocks"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestStitch(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(t)
	mockRenderings := new(mocks.Renderings)
	rootDir := "/tmp/stitcher"
	rd := parsers.ResourceDefinition{
		Name: "product",
	}
	fs := new(afero.MemMapFs)

	stitcher := NewStitcher(rootDir, fs, mockRenderings, rd)

	mockRenderings.On("Model").Return([]byte("model rendering"))
	mockRenderings.On("Repository").Return([]byte("repository rendering"))
	mockRenderings.On("Server").Return([]byte("server rendering"))
	mockRenderings.On("Main").Return([]byte("main rendering"))
	mockRenderings.On("UpMigration").Return([]byte("up migration rendering"))
	mockRenderings.On("DownMigration").Return([]byte("down migration rendering"))
	err := stitcher.Stitch()

	assert.Nil(err)

	modelRendering, modelErr := afero.ReadFile(fs, filepath.Join(rootDir, "models", rd.Name, "model.go"))
	repositoryRendering, repositoryErr := afero.ReadFile(fs, filepath.Join(rootDir, "models", rd.Name, "repository.go"))
	serverRendering, serverErr := afero.ReadFile(fs, filepath.Join(rootDir, "api", rd.Name+"Server.go"))
	mainRendering, mainErr := afero.ReadFile(fs, filepath.Join(rootDir, "main.go"))

	upFiles, _ := afero.Glob(fs, filepath.Join(rootDir, "migrations", "*.up.sql"))
	downFiles, _ := afero.Glob(fs, filepath.Join(rootDir, "migrations", "*.down.sql"))

	upMigrationRendering, upErr := afero.ReadFile(fs, upFiles[0])
	downMigrationRendering, downErr := afero.ReadFile(fs, downFiles[0])

	mockRenderings.AssertExpectations(t)
	assert.Nil(modelErr)
	assert.Nil(repositoryErr)
	assert.Nil(serverErr)
	assert.Nil(upErr)
	assert.Nil(downErr)
	assert.Nil(mainErr)

	assert.Equal("model rendering", string(modelRendering))
	assert.Equal("repository rendering", string(repositoryRendering))
	assert.Equal("server rendering", string(serverRendering))
	assert.Equal("main rendering", string(mainRendering))
	assert.Equal("up migration rendering", string(upMigrationRendering))
	assert.Equal("down migration rendering", string(downMigrationRendering))
}

func TestStitch_WithExistingResources(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(t)
	mockRenderings := new(mocks.Renderings)
	rootDir := "/tmp/stitcher"
	rd := parsers.ResourceDefinition{
		Name: "product",
	}
	fs := new(afero.MemMapFs)

	stitcher := NewStitcher(rootDir, fs, mockRenderings, rd)

	modelFilepath := filepath.Join(rootDir, "models", rd.Name, "model.go")
	repositoryFilepath := filepath.Join(rootDir, "models", rd.Name, "repository.go")
	serverFilepath := filepath.Join(rootDir, "api", rd.Name+"Server.go")
	upMigrationFilepath := filepath.Join(rootDir, "migrations", "123"+".up.sql")
	downMigrationFilepath := filepath.Join(rootDir, "migrations", "123"+".down.sql")

	afero.WriteFile(fs, modelFilepath, []byte("model exists"), 0644)
	afero.WriteFile(fs, repositoryFilepath, []byte("repository exists"), 0644)
	afero.WriteFile(fs, serverFilepath, []byte("server exists"), 0644)
	afero.WriteFile(fs, upMigrationFilepath, []byte("up migration exists"), 0644)
	afero.WriteFile(fs, downMigrationFilepath, []byte("down migration exists"), 0644)

	mockRenderings.On("Model").Return([]byte("model rendering"))
	mockRenderings.On("Repository").Return([]byte("repository rendering"))
	mockRenderings.On("Server").Return([]byte("server rendering"))
	mockRenderings.On("Main").Return([]byte("main rendering"))
	mockRenderings.On("UpMigration").Return([]byte("up migration rendering"))
	mockRenderings.On("DownMigration").Return([]byte("down migration rendering"))

	err := stitcher.Stitch()

	assert.Nil(err)

	modelContents, _ := afero.ReadFile(fs, modelFilepath)
	repositoryContents, _ := afero.ReadFile(fs, repositoryFilepath)
	serverContents, _ := afero.ReadFile(fs, serverFilepath)

	upFiles, _ := afero.Glob(fs, filepath.Join(rootDir, "migrations", "*.up.sql"))
	downFiles, _ := afero.Glob(fs, filepath.Join(rootDir, "migrations", "*.down.sql"))

	upMigrationContents, _ := afero.ReadFile(fs, upFiles[0])
	downMigratinContents, _ := afero.ReadFile(fs, downFiles[0])

	assert.Equal("model exists", string(modelContents))
	assert.Equal("repository exists", string(repositoryContents))
	assert.Equal("server exists", string(serverContents))
	assert.Equal("up migration exists", string(upMigrationContents))
	assert.Equal("down migration exists", string(downMigratinContents))
}

func TestStitch_PartialMainRendering(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(t)
	mockRenderings := new(mocks.Renderings)
	rootDir := "/tmp/stitcher"
	rd := parsers.ResourceDefinition{
		Name: "product",
	}
	fs := new(afero.MemMapFs)

	stitcher := NewStitcher(rootDir, fs, mockRenderings, rd)

	mockRenderings.On("Model").Return([]byte("model rendering"))
	mockRenderings.On("Repository").Return([]byte("repository rendering"))
	mockRenderings.On("Server").Return([]byte("server rendering"))
	mockRenderings.On("Main").Return([]byte("main rendering"))
	mockRenderings.On("UpMigration").Return([]byte("up migration rendering"))
	mockRenderings.On("DownMigration").Return([]byte("down migration rendering"))
	mockRenderings.On("PartialMain").Return([]byte("partial main rendering"))

	afero.WriteFile(fs, filepath.Join(rootDir, "main.go"), []byte("//weaver:renderEnd"), 0644)

	err := stitcher.Stitch()
	assert.Nil(err)

	mainContents, _ := afero.ReadFile(fs, filepath.Join(rootDir, "main.go"))
	assert.Equal("partial main rendering", string(mainContents))
}

func TestStitch_ResourceRendering(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(t)
	mockRenderings := new(mocks.Renderings)
	rootDir := "/tmp/stitcher"
	rd := parsers.ResourceDefinition{
		Name: "product",
	}
	fs := new(afero.MemMapFs)

	stitcher := NewStitcher(rootDir, fs, mockRenderings, rd)

	mockRenderings.On("Model").Return([]byte("model rendering"))
	mockRenderings.On("Repository").Return([]byte("repository rendering"))
	mockRenderings.On("Server").Return([]byte("server rendering"))
	mockRenderings.On("Main").Return([]byte("main rendering"))
	mockRenderings.On("UpMigration").Return([]byte("up migration rendering"))
	mockRenderings.On("DownMigration").Return([]byte("down migration rendering"))
	mockRenderings.On("PartialMain").Return([]byte("partial main rendering"))

	afero.WriteFile(fs, filepath.Join(rootDir, "main.go"), []byte("product\n//weaver:renderEnd"), 0644)

	err := stitcher.Stitch()
	assert.Nil(err)

	mainContents, _ := afero.ReadFile(fs, filepath.Join(rootDir, "main.go"))
	assert.Equal("product\n//weaver:renderEnd", string(mainContents))
}
