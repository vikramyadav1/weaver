package stitcher

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/vikramyadav1/weaver/parsers"
	"github.com/vikramyadav1/weaver/renderer"
	"log"
	"path/filepath"
	"strings"
	"time"
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
	fns := []func(s *stitcher) error{tryCreateModel, tryCreateRepository, tryCreateServer, tryCreateMain, tryCreateMigrations}
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
	fp := filepath.Join(dirPath, "model.go")

	exists, _ := afero.Exists(s.fs, fp)
	if exists {
		log.Printf("Skipping model for %s generation as it already exists.\n", s.rd.Name)
		return nil
	}

	err := s.fs.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	return afero.WriteFile(s.fs, fp, s.renderings.Model(), 0777)
}

func tryCreateRepository(s *stitcher) error {
	dirPath := filepath.Join(s.rootDir, "models", s.rd.Name)
	fp := filepath.Join(dirPath, "repository.go")

	exists, _ := afero.Exists(s.fs, fp)
	if exists {
		log.Printf("Skipping model for %s generation as it already exists.\n", s.rd.Name)
		return nil
	}

	err := s.fs.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	return afero.WriteFile(s.fs, fp, s.renderings.Repository(), 0777)
}

func tryCreateServer(s *stitcher) error {
	dirPath := filepath.Join(s.rootDir, "api")
	fp := filepath.Join(dirPath, s.rd.Name+"Server.go")

	exists, _ := afero.Exists(s.fs, fp)
	if exists {
		log.Printf("Skipping model for %s generation as it already exists.\n", s.rd.Name)
		return nil
	}

	err := s.fs.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	return afero.WriteFile(s.fs, fp, s.renderings.Server(), 0777)
}

func tryCreateMain(s *stitcher) error {
	var mainRendering []byte
	mainFilepath := filepath.Join(s.rootDir, "main.go")

	// Error not caught. Will be handle in future
	mainExists, err := afero.Exists(s.fs, mainFilepath)
	if err != nil {
		log.Fatalf("Error while looking for main file")
		return err
	}

	if !mainExists {
		mainRendering = s.renderings.Main()
		return afero.WriteFile(s.fs, filepath.Join(s.rootDir, "main.go"), mainRendering, 0777)
	}

	mainRendering = s.renderings.PartialMain()
	mainContents, err := afero.ReadFile(s.fs, mainFilepath)
	if err != nil {
		fmt.Printf("Error reading main file.\nError: %v", err)
	}

	if strings.Contains(string(mainContents), s.rd.Name) {
		log.Printf("Main already contains reference to %v resource and therefore, the  resource addition is skipped.", s.rd.Name)
		return nil
	} else {
		newMainContent := strings.Replace(string(mainContents), "//weaver:renderEnd", string(mainRendering), 1)
		return afero.WriteFile(s.fs, filepath.Join(s.rootDir, "main.go"), []byte(newMainContent), 0777)
	}
}

func tryCreateMigrations(s *stitcher) error {
	dirPath := filepath.Join(s.rootDir, "migrations")
	existingUpMigrations, err := afero.Glob(s.fs, filepath.Join(dirPath, "*.up.sql"))
	if err != nil {
		log.Fatalf("Error searching for up migraion. error: %v\n", err)
	}

	existingDownMigrations, err := afero.Glob(s.fs, filepath.Join(dirPath, "*.up.sql"))
	if err != nil {
		log.Fatalf("Error searching for down migraion. error: %v\n", err)
	}

	err = s.fs.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	upFilename := fmt.Sprintf("%s.up.sql", formatted)
	downFilename := fmt.Sprintf("%s.down.sql", formatted)

	if len(existingUpMigrations) == 0 {
		err = afero.WriteFile(s.fs, filepath.Join(dirPath, upFilename), s.renderings.UpMigration(), 0777)
		if err != nil {
			return err
		}
	}

	if len(existingDownMigrations) == 0 {
		return afero.WriteFile(s.fs, filepath.Join(dirPath, downFilename), s.renderings.DownMigration(), 0777)
	}

	return nil
}
