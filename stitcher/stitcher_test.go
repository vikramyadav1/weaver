package stitcher

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/vikramyadav1/weaver/parsers"
	"github.com/vikramyadav1/weaver/renderer/mocks"
	"path/filepath"
	"testing"
)

func TestStitch(t *testing.T) {
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

	err := stitcher.Stitch()

	assert.Nil(err)

	modelRendering, modelErr := afero.ReadFile(fs, filepath.Join(rootDir, "models", rd.Name, "model.go"))
	repositoryRendering, repositoryErr := afero.ReadFile(fs, filepath.Join(rootDir, "models", rd.Name, "repository.go"))
	serverRendering, serverErr := afero.ReadFile(fs, filepath.Join(rootDir, "api", rd.Name+"Server.go"))

	mockRenderings.AssertExpectations(t)
	assert.Nil(modelErr)
	assert.Nil(repositoryErr)
	assert.Nil(serverErr)

	assert.Equal("model rendering", string(modelRendering))
	assert.Equal("repository rendering", string(repositoryRendering))
	assert.Equal("server rendering", string(serverRendering))
}
