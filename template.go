package main

import (
	"embed"
	"io/fs"
	"text/template"

	"github.com/kissjacky/pdevtool/util"
)

const (
	layoutsDir   = "templates/layouts"
	templatesDir = "templates"
	extension    = "/*.html"
)

var (
	//go:embed templates/*
	files embed.FS
)

func LoadTemplates() error {
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			return err
		}

		util.Templates[tmpl.Name()] = pt
	}
	return nil
}
