package main

import (
	"Animation/assets"
	"io/fs"
	"path/filepath"
	"strings"
)

type FileBasedAssetManager struct {
	Assets map[string]*FileBasedAssetEntry
}

func (am *FileBasedAssetManager) Get() (*assets.Asset, error) { return nil, nil }

func (am *FileBasedAssetManager) Scan(rootPath string) error {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		assetPath := strings.ReplaceAll(path[0:len(path)-len(filepath.Ext(path))], "\\", "/")

		if filepath.Ext(path) == ".png" {

		}

		if !d.IsDir() {
			am.assets[assetPath] = &assets.Asset2{Image: nil, OsPath: "." + "/" + path}

		} else {

		}

		return nil

	})

	return nil
}
