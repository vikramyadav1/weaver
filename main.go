package main

import (
	"go.uber.org/zap"
	"os"

	"github.com/spf13/afero"
	"github.com/urfave/cli/v2"
	"github.com/vikramyadav1/weaver/parsers"
	"github.com/vikramyadav1/weaver/renderer"
	"github.com/vikramyadav1/weaver/stitcher"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	app := &cli.App{
		Name:  "weaver",
		Usage: "weaver <path-to-resource-json-file>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "jsonfile",
				Aliases:  []string{"f"},
				Value:    "",
				Usage:    "resource json file path",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "projectdir",
				Aliases:  []string{"d"},
				Value:    ".",
				Usage:    "project root directory",
				Required: true,
			},
		},
		Action: weave,
	}

	err := app.Run(os.Args)
	if err != nil {
		sugar.Errorf("An error occured when running weaver.\nError: %v", err)
	}
}

func weave(c *cli.Context) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	jsonFilePath := c.String("jsonfile")
	sugar.Infof("Rendering resources using %s", jsonFilePath)

	projectDir := c.String("projectdir")
	sugar.Infof("Rendering resources inside %s", projectDir)

	fs := afero.NewOsFs()
	jsonResource, err := afero.ReadFile(fs, jsonFilePath)
	if err != nil {
		return err
	}

	rd := parsers.NewJsonParser(string(jsonResource)).Parse()
	renderings := renderer.NewGoRenderer("templates", rd).Render()
	return stitcher.NewStitcher(projectDir, fs, renderings, rd).Stitch()
}
