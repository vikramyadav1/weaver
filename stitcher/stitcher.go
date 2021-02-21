package stitcher

import (
	"github.com/spf13/afero"
	"github.com/vikramyadav1/weaver/parsers"
	"github.com/vikramyadav1/weaver/renderer"
	"path/filepath"
)

type stitcher struct {
	rootDir    string
	fs         afero.Fs
	renderings renderer.Renderings
	rd         parsers.ResourceDefinition
}

func NewStitcher(rootDir string, fs afero.Fs, r renderer.Renderings, rd parsers.ResourceDefinition) stitcher {
	return stitcher{
		rootDir:    rootDir,
		fs:         fs,
		renderings: r,
		rd:         rd,
	}
}

func (s stitcher) Stitch() error {
	fns := []func(s *stitcher) error{tryCreateModel, tryCreateRepository, tryCreateServer}
	var err error
	for _, fn := range fns {
		if err = fn(&s); err != nil {
			return err
		}
	}
	return nil
}

func tryCreateModel(s *stitcher) error {
	dirPath := filepath.Join(s.rootDir, "models", s.rd.Name)
	err := s.fs.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	return afero.WriteFile(s.fs, filepath.Join(dirPath, "model.go"), s.renderings.Model(), 0777)
}

func tryCreateRepository(s *stitcher) error {
	dirPath := filepath.Join(s.rootDir, "models", s.rd.Name)
	err := s.fs.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	return afero.WriteFile(s.fs, filepath.Join(dirPath, "repository.go"), s.renderings.Repository(), 0777)
}

func tryCreateServer(s *stitcher) error {
	dirPath := filepath.Join(s.rootDir, "api")
	err := s.fs.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	return afero.WriteFile(s.fs, filepath.Join(dirPath, s.rd.Name+"Server.go"), s.renderings.Server(), 0777)
}
